package handlers

import (
	"log"
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
func NewWeb(webPort int, gitScheme string, gitPort int) http.Handler {
	return primary{goGet: goGet(webPort, gitScheme, gitPort), head: head()}
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
