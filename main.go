package main

import (
	"html/template"
	"log"
	"os"

	"github.com/arschles/flexwork/tpl"
	"github.com/arschles/goprox/config"
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

	adminErrCh := startAdminServer(s3Client, s3Conf.Bucket, srvConf.AdminPort)
	webErrCh := startWebServer(s3Client, tplCtx, tmpDir, srvConf, gitConf, s3Conf)
	select {
	case adminErr := <-adminErrCh:
		log.Fatalf("Error running admin server (%s)", adminErr)
	case webErr := <-webErrCh:
		log.Fatalf("Error running web server (%s)", webErr)
	}
}
