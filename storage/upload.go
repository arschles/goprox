package storage

import (
	"archive/tar"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	s3 "github.com/minio/minio-go"
)

const (
	tarContentType = "application/x-tar"
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
	// get all files from dir
	files := []string{}
	if err := filepath.Walk(dir, getWalkFunc(dir, &files)); err != nil {
		return err
	}
	// write all files to the tar writer
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	for _, file := range files {
		fileBytes, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		hdr := &tar.Header{Name: file, Mode: int64(os.ModePerm), Size: int64(len(fileBytes))}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := tw.Write(fileBytes); err != nil {
			return err
		}
	}
	if err := tw.Close(); err != nil {
		return err
	}

	// upload tarball to S3
	if _, err := cl.PutObject(bucketName, packageName, buf, tarContentType); err != nil {
		return err
	}
	return nil
}
