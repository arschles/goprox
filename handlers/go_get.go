package handlers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
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

func getPackage(u *url.URL) string {
	return u.Path[1:]
}

func goGet(webHost string, webPort int, gitScheme, gitHost string, gitPort int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pkg := getPackage(r.URL)
		data := goGetTplData{
			ImportPrefix: fmt.Sprintf("%s:%d/%s", webHost, webPort, pkg),
			RepoRoot:     fmt.Sprintf("%s://%s:%d/%s", gitScheme, gitHost, gitPort, pkg),
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
