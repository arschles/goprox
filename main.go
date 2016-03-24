package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arschles/goprox/handlers"
)

const (
	resp                 = "d049f6c27a2244e12041955e262a404c7faba355	refs/heads/master"
	contentTypeHeaderVal = "Content-Type"
	contentType          = "text/plain; charset=utf-8"
	port                 = 8080
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	mux := http.NewServeMux()
	hostStr := fmt.Sprintf("%s:%d", host, port)
	mux.Handle("/", handlers.New(hostStr))
	// router.Handle()
	// router.HandleFunc("/{repo}/HEAD", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Printf("/abc.git/HEAD")
	// 	w.Write([]byte("ref: refs/heads/master"))
	// })
	// router.HandleFunc("/abc.git/info/refs", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Printf("/abc.git/info/refs")
	// 	w.Header().Set(contentTypeHeaderVal, contentType)
	// 	w.Write([]byte(resp))
	// })
	// router.HandleFunc("/objects/d0/49f6c27a2244e12041955e262a404c7faba355", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Printf("/objects/d0/49f6c27a2244e12041955e262a404c7faba355")
	// 	obj := git.NewObject(git.ObjectBlob, "this is stuff!")
	// 	w.Write(obj.Bytes())
	// })
	log.Printf("hosting on %s", hostStr)
	http.ListenAndServe(hostStr, mux)
}
