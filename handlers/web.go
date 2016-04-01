package handlers

import (
	"log"
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
		goGet: goGet(gitConfig.Scheme, gitConfig.Host, gitConfig.Port),
		head:  head(),
	}
}

func (m primary) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.URL)
	if r.Method == headMethod {
		log.Printf(r.Method)
		m.head.ServeHTTP(w, r)
		return
	} else if r.Method == getMethod && r.URL.Query().Get(goGetQueryKey) == "1" {
		log.Printf("GET with go-get=1")
		m.goGet.ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}
