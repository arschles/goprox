package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Head is a Handler that allows a user to determine whether a package is currently in the cache.
// it is a Handler implementation
type Head struct {
}

// ServeHTTP is the http.Handler interface implementation
func (h *Head) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Register is the Handler interface implementation
func (h *Head) Register(r *mux.Router) {
	r.Handle("/*", h).Methods("HEAD")
}
