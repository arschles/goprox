package handlers

import (
	"net/http"

	"github.com/gorilla"
)

type Handler interface {
	http.Handler
	PathInfo() (string, string)
	Handle(router *mux.Router)
}
