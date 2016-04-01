package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arschles/goprox/handlers"
)

const (
	gitScheme = "http"
	appName   = "goprox"
)

func buildServeMux(host string, port, gitPort int) (string, http.Handler) {
	m := http.NewServeMux()
	hostStr := fmt.Sprintf("%s:%d", host, port)
	m.Handle("/", handlers.NewWeb(port, gitScheme, gitPort))
	return hostStr, m
}

func buildGitMux(host string, port int, tmpDir string) (string, http.Handler) {
	hostStr := fmt.Sprintf("%s:%d", host, port)
	gh := handlers.NewGit(hostStr, tmpDir)
	return hostStr, gh
}

func main() {

	conf, err := getConfig("goprox")
	if err != nil {
		log.Fatalf("Error getting config (%s)", err)
	}
	host := conf.BindHost
	servePort := conf.WebPort
	gitPort := conf.GitPort

	tmpDir, err := createTempDir()
	if err != nil {
		log.Fatalf("Error creating temp dir (%s)", err)
	}
	defer os.RemoveAll(tmpDir)

	srvCh := make(chan error)
	gitCh := make(chan error)
	go func() {
		hostStr, mux := buildServeMux(host, servePort, gitPort)
		log.Printf("Serving web on %s", hostStr)
		srvCh <- http.ListenAndServe(hostStr, mux)
	}()
	go func() {
		hostStr, mux := buildGitMux(host, gitPort, tmpDir)
		log.Printf("Serving git on %s", hostStr)
		gitCh <- http.ListenAndServe(hostStr, mux)
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
