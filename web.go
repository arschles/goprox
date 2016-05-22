package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arschles/flexwork"
	"github.com/arschles/flexwork/tpl"
	"github.com/arschles/goprox/config"
	"github.com/arschles/goprox/handlers"
	s3 "github.com/minio/minio-go"
)

func startWebServer(
	s3Client *s3.Client,
	tplCtx tpl.Context,
	tmpDir string,
	srvConf *config.Server,
	gitConf *config.Git,
	s3Conf *config.S3,
) <-chan error {

	ch := make(chan error)
	go func() {
		webHandler, err := handlers.NewWeb(s3Client, s3Conf.Bucket, srvConf, gitConf, tplCtx)
		if err != nil {
			log.Fatalf("Error creating web handler (%s)", err)
		}
		gitHandler := handlers.NewGit(s3Client, s3Conf.Bucket, tmpDir)
		hostStr := fmt.Sprintf("0.0.0.0:%d", srvConf.BindPort)
		log.Printf("Server config: %s", *srvConf)
		log.Printf("Git config: %s", *gitConf)
		log.Printf("Serving %s and %s on %s", srvConf.Host, gitConf.Host, hostStr)
		ch <- http.ListenAndServe(hostStr, flexwork.MatchHost(map[string]http.Handler{
			srvConf.Host: webHandler,
			gitConf.Host: gitHandler,
		}))
	}()
	return ch
}
