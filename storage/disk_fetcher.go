package storage

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const (
	srcDir = "src"
)

var (
	errNilFileInfo = errors.New("nil file info")
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
	files, err := getFiles(repoPrefix, d.Excludes...)
	if err != nil {
		return nil, err
	}
	log.Printf("files %#v", files)
	return archiveFiles(repoPrefix, files...)
}

// get a list of relative paths of all files under dir. call filepath.Join(dir, file) for each
// returned file to get the absolute path
func getFiles(dir string, excludes ...string) ([]string, error) {
	files := []string{}
	if err := filepath.Walk(dir, getWalkFunc(dir, &files, excludes...)); err != nil {
		return nil, err
	}
	return files, nil
}

func getWalkFunc(baseDir string, files *[]string, excludes ...string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// TODO: handle this case!
			return nil
		}
		if info == nil {
			// TODO: handle this case!
			return errNilFileInfo
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(baseDir, path)
		if err != nil {
			return err
		}
		for _, exclude := range excludes {
			matched, err := regexp.Match(exclude, []byte(rel))
			if err == nil && matched {
				log.Printf("excluding file %s", rel)
				return nil
			}
		}
		*files = append(*files, rel)
		return nil
	}
}
