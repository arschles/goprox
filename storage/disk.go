package storage

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

// UntarToDisk untars the tarball contained in obj to repoDir on disk. Assumes that repoDir already exists, and returns any errors along the way
func UntarToDisk(obj io.Reader, repoDir string) error {
	log.Printf("UntarToDisk to %s", repoDir)
	tr := tar.NewReader(obj)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fullPath := filepath.Join(repoDir, hdr.Name)
		if merr := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); merr != nil {
			return merr
		}
		fd, err := os.Create(filepath.Base(fullPath))
		if err != nil {
			return err
		}
		if _, err := io.Copy(fd, tr); err != nil {
			fd.Close()
			return err
		}
		fd.Close()
	}
	return nil
}
