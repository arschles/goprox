package storage

import (
	s3 "github.com/minio/minio-go"
)

// PackageExists returns the ObjectInfo for the given package if it exists in bucket and cl can access that bucket. Otherwise, returns nil and a descriptive error of what happened
func PackageExists(cl *s3.Client, bucket, pkg string) (*s3.ObjectInfo, error) {
	st, err := cl.StatObject(bucket, pkg)
	if err != nil {
		return nil, err
	}
	return &st, nil
}
