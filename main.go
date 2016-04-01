package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arschles/goprox/handlers"
)

const (
	servePort = 8080
	gitPort   = 8081
	gitScheme = "http"
)

func buildServeMux(host string, port int) (string, http.Handler) {
	m := http.NewServeMux()
	hostStr := fmt.Sprintf("%s:%d", host, port)
	m.Handle("/", handlers.NewWeb(servePort, gitScheme, gitPort))
	return hostStr, m
}

func buildGitMux(host string, port int, tmpDir string) (string, http.Handler) {
	hostStr := fmt.Sprintf("%s:%d", host, port)
	gh := handlers.NewGit(hostStr, tmpDir)
	return hostStr, gh
}

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	tmpDir, err := createTempDir()
	if err != nil {
		log.Fatalf("Error creating temp dir (%s)", err)
	}
	defer os.RemoveAll(tmpDir)

	srvCh := make(chan error)
	gitCh := make(chan error)
	go func() {
		hostStr, mux := buildServeMux(host, servePort)
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
