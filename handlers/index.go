package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/arschles/flexwork/tpl"
	s3 "github.com/minio/minio-go"
)

var (
	funcMap = map[string]interface{}{
		"pluralize": func(num int, singular, plural string) string {
			if num == 0 {
				return singular
			}
			return plural
		},
	}
)

// Index is the handler for the front page of the server
func index(s3Client *s3.Client, bucketName, tplCtx tpl.Context) (http.Handler, error) {
	tpl, err := template.ParseFiles(
		filepath.Join(tplDir, "index.html"),
	)
	if err != nil {
		return nil, err
	}
	tpl = tpl.Funcs(funcMap)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doneCh := make(chan struct{})
		defer close(doneCh)
		objCh := s3Client.ListObjects(bucketName, "", false, doneCh)
		i := 0
		for range objCh {
			i++
		}
		data := map[string]interface{}{
			"NumObjects": i,
		}
		if err := tpl.Execute(w, data); err != nil {
			http.Error(w, "error executing index template", http.StatusInternalServerError)
			return
		}
	}), nil
}
