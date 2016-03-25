package githttp

import (
	"fmt"
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

func (hr HandlerReq) String() string {
	return fmt.Sprintf("RPC = %s, Dir = %s, File = %s", hr.RPC, hr.Dir, hr.File)
}
