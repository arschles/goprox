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

func goGet(webHost string, outwardPort int, gitScheme, gitHost string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get(goGetQueryKey) != "1" {
			http.Error(w, `?go-get="1" expected`, http.StatusBadRequest)
			return
		}
		pkg := getPackage(r.URL)
		repoRoot := fmt.Sprintf("%s://%s:%d/%s", gitScheme, gitHost, outwardPort, pkg)
		importPrefix := fmt.Sprintf("%s:%d/%s", webHost, outwardPort, pkg)
		if (gitScheme == "http" && outwardPort == 80) || (gitScheme == "https" && outwardPort == 443) {
			repoRoot = fmt.Sprintf("%s://%s/%s", gitScheme, gitHost, pkg)
			importPrefix = fmt.Sprintf("%s/%s", webHost, pkg)
		}
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
