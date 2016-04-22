package handlers

import (
	"fmt"
	"net/http"
)

// MatchHost creates a handler that invokes the appropriate handler for the
// incoming request's Host header, according to handlers. It looks for exact
// host values (no regexes) in handlers
func MatchHost(handlers map[string]http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler, found := handlers[r.Header.Get("Host")]
		if !found {
			http.Error(w, fmt.Sprintf("%s not found", r.URL.Path), http.StatusNotFound)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
