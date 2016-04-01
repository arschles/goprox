package storage

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"

	s3 "github.com/minio/minio-go"
)

// UntarObjectToDisk untars the object at obj to repoDir on disk. Assumes that repoDir already exists, and returns any errors along the way
func UntarObjectToDisk(obj *s3.Object, repoDir string) error {
	tr := tar.NewReader(obj)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fullPath := filepath.Join(repoDir, hdr.Name)
		if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
			return err
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
