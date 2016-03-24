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

func goGet(srvRoot string) http.Handler {
	tpl := template.Must(template.New("").Parse(`<html>
<head><meta name="go-import" content="{{.ImportPrefix}} git {{.RepoRoot}}"/></head>
<body></body>
</html>
`))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		importPrefix := r.URL.Path[1:]
		data := goGetTplData{
			ImportPrefix: importPrefix,
			RepoRoot:     fmt.Sprintf("https://%s/%s", srvRoot, importPrefix),
		}
		if err := tpl.Execute(w, data); err != nil {
			log.Printf("Error executing template (%s)", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
