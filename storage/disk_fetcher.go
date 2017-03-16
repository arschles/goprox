package storage

import (
	"bytes"
	"io"
	"log"
	"path/filepath"

	"github.com/arschles/goprox/files"
)

const (
	srcDir = "src"
)

// DiskFetcher is a Fetcher that reads from the gopath given. it ignores all versions and
// will not change the gopath in any way
type DiskFetcher struct {
	Gopath   string
	Excludes []string
}

// GetContents is the fetcher interface implementation
func (d DiskFetcher) GetContents(pkgName string, version string) (io.Reader, error) {
	log.Printf("package %s", pkgName)
	goPathSrcDir := filepath.Join(d.Gopath, srcDir)
	repoPrefix := filepath.Join(goPathSrcDir, pkgName)
	files, err := files.List(repoPrefix, d.Excludes...)
	if err != nil {
		return nil, err
	}
	log.Printf("files %#v", files)
	buf := new(bytes.Buffer)
	if err := archiveFiles(repoPrefix, buf, files...); err != nil {
		return nil, err
	}
	return buf, err
}
