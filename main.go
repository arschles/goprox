package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/arschles/goprox/handlers"
)

const (
	resp                 = "d049f6c27a2244e12041955e262a404c7faba355	refs/heads/master"
	contentTypeHeaderVal = "Content-Type"
	contentType          = "text/plain; charset=utf-8"
	servePort            = 8080
	gitPort              = 8081
)

func buildServeMux(host string, port int) (string, http.Handler) {
	m := http.NewServeMux()
	hostStr := fmt.Sprintf("%s:%d", host, port)
	m.Handle("/", handlers.NewWeb(hostStr))
	return hostStr, m
}

func buildGitMux(host string, port int, tmpDir string) (string, http.Handler) {
	m := http.NewServeMux()
	hostStr := fmt.Sprintf("%s:%d", host, port)
	gh := handlers.NewGit(hostStr, tmpDir)
	m.Handle("/", gh)
	return hostStr, m
}

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	tmpDir, err := ioutil.TempDir("", "/goprox/")
	if err != nil {
		log.Printf("Error creating temp dir (%s)", err)
		os.Exit(1)
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
