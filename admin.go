package main

import (
	"log"

	"github.com/arschles/goprox/admin"
	s3 "github.com/minio/minio-go"
)

func startAdminServer(s3Client *s3.Client, bucketName string, port int) <-chan error {
	log.Printf("starting admin server on port %d", port)
	ch := make(chan error)
	go func() {
		ch <- admin.StartServer(s3Client, bucketName, port)
	}()
	return ch
}
