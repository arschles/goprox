package storage

import (
	s3 "github.com/minio/minio-go"
)

const (
	tarContentType = "application/x-tar"
)

// UploadPackage uploads all files in dir to the correct location under bucketName for packageName
func UploadPackage(cl *s3.Client, bucketName, packageName, dir string) error {
	files, err := getFiles(dir)
	if err != nil {
		return err
	}

	buf, err := archiveFiles(dir, files...)
	if err != nil {
		return err
	}

	// upload tarball to S3
	if _, err := cl.PutObject(bucketName, Name(packageName), buf, tarContentType); err != nil {
		return err
	}
	return nil
}
