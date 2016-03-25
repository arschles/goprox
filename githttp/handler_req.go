package githttp

import (
	"net/http"
)

// HandlerReq is git request handler
type HandlerReq struct {
	w    http.ResponseWriter
	r    *http.Request
	RPC  string
	Dir  string
	File string
}
