package handlers

import (
	"net/http"

	"github.com/arschles/flexwork"
	"github.com/arschles/goprox/config"
)

// NewWeb returns the main handler responsible for serving web traffic, including 'go get' traffic
func NewWeb(webConfig *config.Server, gitConfig *config.Git) http.Handler {
	return flexwork.MethodMux(map[flexwork.Method]http.Handler{
		flexwork.Get:  goGet(webConfig.Host, webConfig.OutwardPort, gitConfig.Scheme, gitConfig.Host),
		flexwork.Head: head(),
	})
}
