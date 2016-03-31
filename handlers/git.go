package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/arschles/goprox/githttp"
	s3 "github.com/minio/minio-go"
)

// NewGit returns the handler to be used for the Git server
func NewGit(hostStr, tmpDir string, s3Client *s3.Client) http.Handler {
	hdl := githttp.New(tmpDir)
	// hdl.UploadPack = false
	hdl.FillRepo = func(repoDir string) error {
		if !strings.HasPrefix(repoDir, tmpDir) {
			return fmt.Errorf("invalid repoDir in FillRepo (%s)", repoDir)
		}
		repoName := repoDir[len(tmpDir)+1:]
		if err := os.MkdirAll(repoDir, os.ModePerm); err != nil {
			log.Printf("error creating %s (%s)", repoDir, err)
			return err
		}
		cmd := exec.Command("git", "clone", fmt.Sprintf("https://%s", repoName), repoDir)
		cmd.Dir = repoDir
		log.Printf("executing %s in %s", strings.Join(cmd.Args, " "), cmd.Dir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Printf("error cloning (%s)", err)
			return err
		}
		log.Printf("done filling %s", repoName)
		// TODO: do caching here. see https://github.com/arschles/goprox/issues/3
		return nil
	}
	return hdl
}
