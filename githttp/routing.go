package githttp

import (
	"log"
	"net/http"
	"os"
	"strings"
)

// Request handling function
func (g *GitHTTP) requestHandler(w http.ResponseWriter, r *http.Request) {
	// Get service for URL
	repo, service := getServiceForPath(g, r.URL.Path)

	// No url match
	if service == nil {
		renderNotFound(w)
		return
	}

	// Bad method
	if service.Method != r.Method {
		renderMethodNotAllowed(w, r)
		return
	}

	// Rpc type
	rpc := service.RPC

	// Get specific file
	file := strings.Replace(r.URL.Path, repo+"/", "", 1)

	// Resolve directory
	dir, err := g.getGitDir(repo)

	// Repo not found on disk
	if err != nil {
		log.Printf("not found (%s)", err)
		renderNotFound(w)
		return
	}

	// Build request info for handler
	hr := HandlerReq{
		w:    w,
		r:    r,
		RPC:  rpc,
		Dir:  dir,
		File: file,
	}

	// Call handler
	if err := service.Handler(hr); err != nil {
		if os.IsNotExist(err) {
			renderNotFound(w)
			return
		}
		switch err.(type) {
		case *ErrorNoAccess:
			renderNoAccess(w)
			return
		}
		http.Error(w, err.Error(), 500)
	}
}
