package githttp

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	uploadPack  = "upload-pack"
	receivePack = "receive-pack"
)

// GitHTTP is an http.Handler that knows how to serve the git protocol over HTTP
type GitHTTP struct {
	// Root directory to serve repos from
	ProjectRoot string

	// Path to git binary
	GitBinPath string

	// Access rules
	UploadPack  bool
	ReceivePack bool

	// Event handling functions
	EventHandler func(ev Event)

	// FillRepo is called when a repository is not found. This function has the chance to get it from somewhere else.
	// If this returns an error or is nil, then the git operation will fail
	FillRepo func(string) error
}

// Implement the http.Handler interface
func (g *GitHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestHandler(g, w, r)
	return
}

// New is the shorthand constructor for most common scenario
func New(root string) *GitHTTP {
	return &GitHTTP{
		ProjectRoot: root,
		GitBinPath:  "/usr/bin/git",
		UploadPack:  true,
		ReceivePack: true,
	}
}

// Init builds the root directory if doesn't exist
func (g *GitHTTP) Init() (*GitHTTP, error) {
	if err := os.MkdirAll(g.ProjectRoot, os.ModePerm); err != nil {
		return nil, err
	}
	return g, nil
}

// Publish event if EventHandler is set
func (g *GitHTTP) event(e Event) {
	if g.EventHandler != nil {
		g.EventHandler(e)
	}
}

// Actual command handling functions

func (g *GitHTTP) serviceRPC(hr HandlerReq) error {
	access, err := g.hasAccess(hr.r, hr.Dir, hr.RPC, true)
	if err != nil {
		return err
	}

	if access == false {
		return &ErrorNoAccess{hr.Dir}
	}

	// Reader that decompresses if necessary
	reader, err := requestReader(hr.r)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Reader that scans for events
	rpcReader := &RpcReader{
		Reader: reader,
		Rpc:    hr.RPC,
	}

	// Set content type
	hr.w.Header().Set("Content-Type", fmt.Sprintf("application/x-git-%s-result", hr.RPC))

	args := []string{hr.RPC, "--stateless-rpc", "."}
	cmd := exec.Command(g.GitBinPath, args...)
	cmd.Dir = hr.Dir
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()

	err = cmd.Start()
	if err != nil {
		return err
	}

	// Scan's git command's output for errors
	gitReader := &GitReader{
		Reader: stdout,
	}

	// Copy input to git binary
	io.Copy(stdin, rpcReader)

	// Write git binary's output to http response
	io.Copy(hr.w, gitReader)

	// Wait till command has completed
	mainError := cmd.Wait()

	if mainError == nil {
		mainError = gitReader.GitError
	}

	// Fire events
	for _, e := range rpcReader.Events {
		// Set directory to current repo
		e.Dir = hr.Dir
		e.Request = hr.r
		e.Error = mainError

		// Fire event
		g.event(e)
	}

	// May be nil if all is good
	return mainError
}

func (g *GitHTTP) getInfoRefs(hr HandlerReq) error {
	serviceName := getServiceType(hr.r)
	access, err := g.hasAccess(hr.r, hr.Dir, serviceName, false)
	if err != nil {
		return err
	}

	if !access {
		g.updateServerInfo(hr.Dir)
		hdrNocache(hr.w)
		return sendFile("text/plain; charset=utf-8", hr)
	}

	refs, err := g.gitCommand(hr.Dir, serviceName, "--stateless-rpc", "--advertise-refs", hr.Dir)
	if err != nil {
		return err
	}

	hdrNocache(hr.w)
	hr.w.Header().Set("Content-Type", fmt.Sprintf("application/x-git-%s-advertisement", serviceName))
	hr.w.WriteHeader(http.StatusOK)
	hr.w.Write(packetWrite("# service=git-" + serviceName + "\n"))
	hr.w.Write(packetFlush())
	hr.w.Write(refs)

	return nil
}

func (g *GitHTTP) getInfoPacks(hr HandlerReq) error {
	hdrCacheForever(hr.w)
	return sendFile("text/plain; charset=utf-8", hr)
}

func (g *GitHTTP) getLooseObject(hr HandlerReq) error {
	hdrCacheForever(hr.w)
	return sendFile("application/x-git-loose-object", hr)
}

func (g *GitHTTP) getPackFile(hr HandlerReq) error {
	hdrCacheForever(hr.w)
	return sendFile("application/x-git-packed-objects", hr)
}

func (g *GitHTTP) getIdxFile(hr HandlerReq) error {
	hdrCacheForever(hr.w)
	return sendFile("application/x-git-packed-objects-toc", hr)
}

func (g *GitHTTP) getTextFile(hr HandlerReq) error {
	hdrNocache(hr.w)
	return sendFile("text/plain", hr)
}

// Logic helping functions

func sendFile(contentType string, hr HandlerReq) error {
	w, r := hr.w, hr.r
	reqFile := path.Join(hr.Dir, hr.File)

	f, err := os.Stat(reqFile)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", f.Size()))
	w.Header().Set("Last-Modified", f.ModTime().Format(http.TimeFormat))
	http.ServeFile(w, r, reqFile)

	return nil
}

func (g *GitHTTP) getGitDir(filePath string) (string, error) {
	root := g.ProjectRoot

	if root == "" {
		cwd, err := os.Getwd()

		if err != nil {
			return "", err
		}

		root = cwd
	}

	f := path.Join(root, filePath)
	_, statErr := os.Stat(f)
	if statErr != nil && !os.IsNotExist(statErr) {
		return "", statErr
	}
	if statErr != nil && g.FillRepo == nil {
		return "", statErr
	}
	if statErr != nil {
		if err := g.FillRepo(f); err != nil {
			return "", err
		}
	}
	return f, nil
}

func (g *GitHTTP) hasAccess(r *http.Request, dir string, rpc string, checkContentType bool) (bool, error) {
	if checkContentType {
		if r.Header.Get("Content-Type") != fmt.Sprintf("application/x-git-%s-request", rpc) {
			return false, nil
		}
	}

	if !(rpc == uploadPack || rpc == receivePack) {
		return false, nil
	}
	if rpc == receivePack {
		return g.ReceivePack, nil
	}
	if rpc == uploadPack {
		return g.UploadPack, nil
	}

	return g.getConfigSetting(rpc, dir)
}

func (g *GitHTTP) getConfigSetting(service_name string, dir string) (bool, error) {
	service_name = strings.Replace(service_name, "-", "", -1)
	setting, err := g.getGitConfig("http."+service_name, dir)
	if err != nil {
		return false, nil
	}

	if service_name == "uploadpack" {
		return setting != "false", nil
	}

	return setting == "true", nil
}

func (g *GitHTTP) getGitConfig(configName string, dir string) (string, error) {
	args := []string{"config", configName}
	out, err := g.gitCommand(dir, args...)
	if err != nil {
		return "", err
	}
	return string(out)[0 : len(out)-1], nil
}

func (g *GitHTTP) updateServerInfo(dir string) ([]byte, error) {
	args := []string{"update-server-info"}
	return g.gitCommand(dir, args...)
}

func (g *GitHTTP) gitCommand(dir string, args ...string) ([]byte, error) {
	command := exec.Command(g.GitBinPath, args...)
	command.Dir = dir
	out, err := command.CombinedOutput()
	if err != nil {
		return nil, newCmdErr(command, out, err)
	}
	return out, nil
}
