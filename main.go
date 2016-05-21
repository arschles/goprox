package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/arschles/flexwork"
	"github.com/arschles/flexwork/tpl"
	"github.com/arschles/goprox/config"
	"github.com/arschles/goprox/handlers"
	s3 "github.com/minio/minio-go"
)

const (
	appName = "goprox"
)

var (
	funcs = template.FuncMap(map[string]interface{}{
		"pluralize": func(i int, singular, plural string) string {
			if i == 1 {
				return singular
			}
			return plural
		},
	})
)

func createTplCtx(baseDir string, funcs template.FuncMap) tpl.Context {
	return tpl.NewCachingContext(baseDir, funcs)
}

func main() {
	gitConf, err := config.GetGit(appName)
	if err != nil {
		log.Fatalf("Error getting git config (%s)", err)
	}
	srvConf, err := config.GetServer(appName)
	if err != nil {
		log.Fatalf("Error getting web config (%s)", err)
	}
	s3Conf, err := config.GetS3(appName)
	if err != nil {
		log.Fatalf("Error getting S3 config (%s)", err)
	}

	log.Printf("Creating caching context for directory %s", srvConf.TemplateBaseDir)
	tplCtx := createTplCtx(srvConf.TemplateBaseDir, funcs)

	tmpDir, err := createTempDir()
	if err != nil {
		log.Fatalf("Error creating temp dir (%s)", err)
	}
	defer os.RemoveAll(tmpDir)

	s3Client, err := s3.New(s3Conf.Endpoint, s3Conf.Key, s3Conf.Secret, false)
	if err != nil {
		log.Fatalf("Error creating new S3 client (%s)", err)
	}

	webHandler, err := handlers.NewWeb(s3Client, s3Conf.Bucket, srvConf, gitConf, tplCtx)
	if err != nil {
		log.Fatalf("Error creating web handler (%s)", err)
	}
	gitHandler := handlers.NewGit(s3Client, s3Conf.Bucket, tmpDir)

	hostStr := fmt.Sprintf("0.0.0.0:%d", srvConf.BindPort)
	log.Printf("Server config: %s", *srvConf)
	log.Printf("Git config: %s", *gitConf)
	log.Printf("Serving %s and %s on %s", srvConf.Host, gitConf.Host, hostStr)
	if err := http.ListenAndServe(hostStr, flexwork.MatchHost(map[string]http.Handler{
		srvConf.Host: webHandler,
		gitConf.Host: gitHandler,
	})); err != nil {
		log.Fatalf("Error running web server (%s)", err)
	}
}
