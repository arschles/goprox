package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type goGetTplData struct {
	ImportPrefix string
	RepoRoot     string
}

var (
	goGetHTMLTpl = template.Must(template.New("").Parse(`<html>
<head><meta name="go-import" content="{{.ImportPrefix}} git {{.RepoRoot}}"/></head>
<body></body>
</html>
`))
)

func goGet(gitSrvScheme string, gitSrvPort int, srvRoot string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		importPrefix := r.URL.Path[1:]
		data := goGetTplData{
			ImportPrefix: importPrefix,
			RepoRoot:     fmt.Sprintf("%s://%s:%d/%s", gitSrvScheme, srvRoot, gitSrvPort, importPrefix),
		}
		if err := goGetHTMLTpl.Execute(w, data); err != nil {
			log.Printf("Error executing template (%s)", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
