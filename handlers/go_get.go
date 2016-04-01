package handlers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/arschles/goprox/repo"
)

type goGetTplData struct {
	ImportPrefix string
	RepoRoot     string
}

var (
	goGetHTMLTpl = template.Must(template.New("").Parse(`<html>
<head>
<meta name="go-import" content="{{.ImportPrefix}} git {{.RepoRoot}}">
</head>
<body></body>
</html>
`))
)

func goGet(webHost string, webPort int, gitScheme, gitHost string, gitPort int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host = %s", r.URL.Host)
		repoInfo, err := repo.InfoFromURL(r.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("go-get import prefix = %s", repoInfo.ImportPrefix())
		data := goGetTplData{
			ImportPrefix: repoInfo.ImportPrefix(),
			RepoRoot:     fmt.Sprintf("%s://%s/%s", gitScheme, r.URL.Host, repoInfo.Package()),
		}
		log.Printf("template data = %+v", data)
		out := io.MultiWriter(w, os.Stdout)
		if err := goGetHTMLTpl.Execute(out, data); err != nil {
			log.Printf("Error executing template (%s)", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
