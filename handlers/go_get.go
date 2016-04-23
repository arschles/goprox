package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

const (
	goGetQueryKey = "go-get"
)

var (
	goGetHTMLTpl = template.Must(template.New("").Parse(`<html>
<head>
<meta name="go-import" content="{{.ImportPrefix}} git {{.RepoRoot}}">
</head>
<body></body>
</html>
`))
)

type goGetTplData struct {
	ImportPrefix string
	RepoRoot     string
}

func getPackage(u *url.URL) string {
	return u.Path[1:]
}

func standardPort(scheme string, port int) bool {
	return (scheme == "http" && port == 80) || (scheme == "https" && port == 443)
}

func getRepoRoot(gitScheme, gitHost, pkg string, outwardPort int) string {
	repoRoot := fmt.Sprintf("%s://%s:%d/%s", gitScheme, gitHost, outwardPort, pkg)
	if standardPort(gitScheme, outwardPort) {
		repoRoot = fmt.Sprintf("%s://%s/%s", gitScheme, gitHost, pkg)
	}
	return repoRoot
}

func getImportPrefix(gitScheme, webHost, pkg string, outwardPort int) string {
	importPrefix := fmt.Sprintf("%s:%d/%s", webHost, outwardPort, pkg)
	if standardPort(gitScheme, outwardPort) {
		importPrefix = fmt.Sprintf("%s/%s", webHost, pkg)
	}
	return importPrefix
}

func goGet(webHost string, outwardPort int, gitScheme, gitHost string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get(goGetQueryKey) != "1" {
			http.Error(w, `?go-get="1" expected`, http.StatusBadRequest)
			return
		}
		pkg := getPackage(r.URL)
		repoRoot := getRepoRoot(gitScheme, gitHost, pkg, outwardPort)
		importPrefix := getImportPrefix(gitScheme, webHost, pkg, outwardPort)
		data := goGetTplData{
			ImportPrefix: importPrefix,
			RepoRoot:     repoRoot,
		}
		if err := goGetHTMLTpl.Execute(w, data); err != nil {
			log.Printf("Error executing template (%s)", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
