package storage

import (
	"bytes"
	"context"
	"io"
	"path/filepath"

	"github.com/arschles/goprox/files"
	"github.com/arschles/goprox/logs"
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
func (d DiskFetcher) GetContents(ctx context.Context, pkgName string, version string) (io.Reader, error) {
	logger := logs.FromContext(ctx)
	logger.Printf("Package %s", pkgName)
	goPathSrcDir := filepath.Join(d.Gopath, srcDir)
	repoPrefix := filepath.Join(goPathSrcDir, pkgName)
	files, err := files.List(ctx, repoPrefix, d.Excludes...)
	if err != nil {
		return nil, err
	}
	logger.Printf("Files: %#v", files)
	buf := new(bytes.Buffer)
	if err := archiveFiles(repoPrefix, buf, files...); err != nil {
		return nil, err
	}
	return buf, err
}
