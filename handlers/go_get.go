package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	goGetRespTplStr = `<html>
<head>
	<meta name="go-import" content="{{.Prefix}} {{.VCS}} {{.RepoRoot}}></meta>
</head>
<body>
	{{.GoGetCmd}}
</body>
</html>`
)

var (
	goGetRespTpl = template.Must(template.New("GoGetTpl").Parse(goGetRespTplStr))
)

type goGetTplData struct {
	Prefix   string
	VCS      string
	RepoRoot string
}

// GoGet is Handler implementation to handle the endpoint that "go get" makes requests to.
// for example, it is able to handle "https://goproxserver.com/github.com/my/package?go-get=1"
type GoGet struct {
	VCS     string
	GitHost string
}

func (g *GoGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tplData := goGetTplData{Prefix: prefix, VCS: g.VCS, RepoRoot: g.RepoRoot}
}

// Register is the Handler interface implementation
func (g *GoGet) Register(r *mux.Router) {
	r.Handle("*", g).Methods("GET").Queries("go-get", "1")
}
