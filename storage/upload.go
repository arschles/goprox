package storage

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	s3 "github.com/minio/minio-go"
)

func getWalkFunc(baseDir string, files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(baseDir, path)
		if err != nil {
			return err
		}
		*files = append(*files, rel)
		return nil
	}
}

// UploadPackage uploads all files in dir to the correct location under bucketName for packageName
func UploadPackage(cl *s3.Client, bucketName, packageName, dir string) error {
	// TODO: make this completely tarball-based
	files := []string{}
	if err := filepath.Walk(dir, getWalkFunc(dir, &files)); err != nil {
		return err
	}
	for _, file := range files {
		objName := strings.Join([]string{packageName, file}, "/")
		absFile := filepath.Join(dir, file)
		go func() {
			if _, err := cl.FPutObject(bucketName, objName, absFile, ""); err != nil {
				log.Printf("Error uploading file %s to %s (%s)", absFile, objName, err)
				return
			}
		}()
	}
	return nil
}
