package handlers

import (
	"net/http"
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
func NewWeb(srvScheme string, srvPort int, srvRoot string) http.Handler {
	return primary{goGet: goGet(srvScheme, srvPort, srvRoot), head: head()}
}

func (m primary) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == headMethod {
		m.head.ServeHTTP(w, r)
		return
	}
	if r.URL.Query().Get(goGetQueryKey) == "1" && r.Method == getMethod {
		m.goGet.ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}
