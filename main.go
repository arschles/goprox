package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arschles/goprox/config"
	"github.com/arschles/goprox/handlers"
)

const (
	servePort = 8080
	gitPort   = 8081
	gitScheme = "http"
	appName   = "goprox"
)

func main() {
	gitConf, err := config.GetGitServer(appName)
	if err != nil {
		log.Fatalf("Error getting git config (%s)", err)
	}
	webConf, err := config.GetWebServer(appName)
	if err != nil {
		log.Fatalf("Error getting web config (%s)", err)
	}

	tmpDir, err := createTempDir()
	if err != nil {
		log.Fatalf("Error creating temp dir (%s)", err)
	}
	defer os.RemoveAll(tmpDir)

	srvCh := make(chan error)
	gitCh := make(chan error)
	go func() {
		hostStr := fmt.Sprintf("%s:%d", webConf.BindHost, webConf.BindPort)
		handler := handlers.NewWeb(webConf, gitConf)
		log.Printf("Serving web on %s", hostStr)
		srvCh <- http.ListenAndServe(hostStr, handler)
	}()
	go func() {
		hostStr := fmt.Sprintf("%s:%d", gitConf.BindHost, gitConf.BindPort)
		log.Printf("Serving git on %s", hostStr)
		handler := handlers.NewGit(tmpDir)
		gitCh <- http.ListenAndServe(hostStr, handler)
	}()
	select {
	case err := <-srvCh:
		log.Printf("Error serving web (%s)", err)
		os.Exit(1)
	case err := <-gitCh:
		log.Printf("Error serving git (%s)", err)
		os.Exit(1)
	}
}
