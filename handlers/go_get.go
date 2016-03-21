package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GoGet is Handler implementation to handle the endpoint that "go get" makes requests to.
// for example, it is able to handle "https://goproxserver.com/github.com/my/package?go-get=1"
type GoGet struct {
}

func (g *GoGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// Register is the Handler interface implementation
func (g *GoGet) Register(r *mux.Router) {
	r.Handle("*", g).Methods("GET").Queries("go-get", "1")
}
