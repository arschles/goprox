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
		host := r.Host
		handler, found := handlers[host]
		if !found {
			http.Error(w, fmt.Sprintf("%s not found for host %s", r.URL.Path, host), http.StatusNotFound)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
