package handlers

import (
	"log"
	"net/http"

	githttp "github.com/arschles/go-git-http"
)

func gitClone(repoName string) error {
	log.Printf("cloning repo %s", repoName)
	return nil
}

// NewGit returns the handler to be used for the Git server
func NewGit(hostStr, tmpDir string) http.Handler {
	hdl := githttp.New(tmpDir)
	// hdl.UploadPack = false
	hdl.EventHandler = func(ev githttp.Event) {
		if ev.Type == githttp.FETCH {
			gitClone(ev.Dir)
			return
		}
	}
	return hdl
}
