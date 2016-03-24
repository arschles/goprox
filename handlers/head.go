package handlers

import (
	"net/http"
)

// Head is a Handler that allows a user to determine whether a package is currently in the cache.
// it is a Handler implementation
func head() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not yet implemented", http.StatusNotImplemented)
	})
}
