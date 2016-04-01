package handlers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
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

func goGet(webPort int, gitScheme string, gitPort int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		importPrefix := r.URL.Path[1:]
		log.Printf("go-get import prefix = %s", importPrefix)
		host := r.URL.Host
		data := goGetTplData{
			// ImportPrefix: importPrefix,
			ImportPrefix: fmt.Sprintf("%s:%d/%s", host, webPort, importPrefix),
			RepoRoot:     fmt.Sprintf("%s://%s:%d/%s", gitScheme, host, gitPort, importPrefix),
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
