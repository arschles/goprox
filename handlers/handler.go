package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is a http.Handler that can also register itself with a gorilla mux router
type Handler interface {
	http.Handler
	Register(r *mux.Router)
}
