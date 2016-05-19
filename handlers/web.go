package handlers

import (
	"net/http"

	"github.com/arschles/flexwork/tpl"
	"github.com/arschles/goprox/config"
	s3 "github.com/minio/minio-go"
)

const (
	headMethod    = "HEAD"
	getMethod     = "GET"
	goGetQueryKey = "go-get"
)

type primary struct {
	index http.Handler
	goGet http.Handler
	head  http.Handler
}

// NewWeb returns the main handler responsible for serving web traffic, including 'go get' traffic
func NewWeb(
	s3Client *s3.Client,
	bucketName string,
	webConfig *config.Server,
	gitConfig *config.Git,
	tplCtx tpl.Context,
) (http.Handler, error) {

	idx, err := index(s3Client, bucketName, tplCtx)
	if err != nil {
		return nil, err
	}
	return primary{
		index: idx,
		goGet: goGet(webConfig.Host, webConfig.OutwardPort, gitConfig.Scheme, gitConfig.Host),
		head:  head(),
	}, nil
}

func (m primary) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == headMethod {
		m.head.ServeHTTP(w, r)
		return
	} else if r.Method == getMethod && r.URL.Query().Get(goGetQueryKey) == "1" {
		m.goGet.ServeHTTP(w, r)
		return
	} else if r.Method == getMethod {
		m.index.ServeHTTP(w, r)
	}

	http.NotFound(w, r)
}
