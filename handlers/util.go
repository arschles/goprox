package handlers

import (
	"github.com/gorilla/mux"
)

// Register calls Register(r) on each handler in handlers
func Register(r *mux.Router, handlers ...Handler) {
	for _, handler := range handlers {
		handler.Register(r)
	}
}
