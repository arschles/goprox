package handlers

import (
	"net/http"

	"github.com/arschles/goprox/config"
)

const (
	headMethod    = "HEAD"
	getMethod     = "GET"
	goGetQueryKey = "go-get"
)

type primary struct {
	goGet http.Handler
	head  http.Handler
}

// NewWeb returns the main handler responsible for serving web traffic, including 'go get' traffic
func NewWeb(webConfig *config.WebServer, gitConfig *config.GitServer) http.Handler {
	return primary{
		goGet: goGet(webConfig.Host, webConfig.Port, gitConfig.Scheme, gitConfig.Host, gitConfig.Port),
		head:  head(),
	}
}

func (m primary) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == headMethod {
		m.head.ServeHTTP(w, r)
		return
	} else if r.Method == getMethod && r.URL.Query().Get(goGetQueryKey) == "1" {
		m.goGet.ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}
