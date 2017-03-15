package storage

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
	files, err := getFiles(repoPrefix, d.Excludes...)
	if err != nil {
		return nil, err
	}
	log.Printf("files %#v", files)
	return tarFiles(repoPrefix, files...)
}

// get a list of relative paths of all files under dir
func getFiles(dir string, excludes ...string) ([]string, error) {
	files := []string{}
	if err := filepath.Walk(dir, getWalkFunc(dir, &files, excludes...)); err != nil {
		return nil, err
	}
	return files, nil
}

// tarFiles produces a tarball of all files. It reads each file that lives under prefix and stores it in the tarball as the path in the slice
func tarFiles(prefix string, files ...string) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	for _, file := range files {
		fileBytes, err := ioutil.ReadFile(filepath.Join(prefix, file))
		if err != nil {
			return nil, err
		}
		hdr := &tar.Header{Name: file, Mode: int64(os.ModePerm), Size: int64(len(fileBytes))}
		if err := tw.WriteHeader(hdr); err != nil {
			return nil, err
		}
		if _, err := tw.Write(fileBytes); err != nil {
			return nil, err
		}
	}
	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf, nil
}

func getWalkFunc(baseDir string, files *[]string, excludes ...string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
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
