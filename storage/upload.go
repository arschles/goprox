package storage

import (
	"archive/tar"
	"bytes"
	"io"
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

// get a list of relative paths of all files under dir
func getFiles(dir string) ([]string, error) {
	files := []string{}
	if err := filepath.Walk(dir, getWalkFunc(dir, &files)); err != nil {
		return nil, err
	}
	return files, nil
}

// tarFiles produces a tarball of all files. It reads each file that lives under prefix and stores it in the tarball as the path in the slice
func tarFiles(prefix string, files []string) (io.Reader, error) {
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

// UploadPackage uploads all files in dir to the correct location under bucketName for packageName
func UploadPackage(cl *s3.Client, bucketName, packageName, dir string) error {
	files, err := getFiles(dir)
	if err != nil {
		return err
	}

	buf, err := tarFiles(dir, files)
	if err != nil {
		return err
	}

	// upload tarball to S3
	if _, err := cl.PutObject(bucketName, Name(packageName), buf, tarContentType); err != nil {
		return err
	}
	return nil
}
