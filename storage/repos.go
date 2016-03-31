package storage

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	s3 "github.com/minio/minio-go"
)

// File is a single file in a package repsitory
type File struct {
	Key string
	*s3.Object
}

// PackageStream is a collection of files that make up a package that are being streamed from S3
type PackageStream struct {
	doneCh chan<- struct{}
	// Files is the unordered collection of files streamed down from S3. This channel will be closed when there are no more files available.
	Files <-chan *File
}

func (o *PackageStream) Next() io.Reader {
	return nil
}

// Close stops files from streaming from S3. Always call this when you're done, even if no more files have been streamed
func (o *PackageStream) Close() error {
	close(o.doneCh)
	return nil
}

func StreamForPackage(cl *s3.Client, bucketName, pkgName string) (*PackageStream, error) {
	doneCh := make(chan struct{})
	objectInfoCh := cl.ListObjects(bucketName, pkgName, true, doneCh)

}

// DownloadPackage downloads the package at bucketName/pkgName to localDir
func DownloadPackage(cl *s3.Client, bucketName, pkgName, localDir string) error {
	doneCh := make(chan struct{})
	defer close(doneCh)
	objectInfoCh := cl.ListObjects(bucketName, pkgName, true, doneCh)

	var wg sync.WaitGroup
	objAndInfoCh := make(chan objAndInfo)
	for objectInfo := range objectInfoCh {
		wg.Add(1)
		go func(objInfo s3.ObjectInfo) {
			defer wg.Done()
			obj, err := cl.GetObject(bucketName, objInfo.Key)
			if err != nil {
				log.Printf("Error downloading object %s (%s)", objInfo.Key, err)
				return
			}
			objAndInfoCh <- objAndInfo{obj: obj, info: objInfo}
		}(objectInfo)
	}
	go func() {
		wg.Wait()
		close(objAndInfoCh)
	}()

	for objAndInfo := range objAndInfoCh {
		full := filepath.Join(localDir, pkgName, objAndInfo.info.Key)
		dir := filepath.Dir(full)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Printf("Error creating directory %s (%s)", dir, err)
			objAndInfo.obj.Close()
			continue
		}
		fd, err := os.Create(full)
		if err != nil {
			log.Printf("Error creating file %s (%s)", full, err)
			objAndInfo.obj.Close()
			continue
		}
		if _, err := io.Copy(fd, objAndInfo.obj); err != nil {
			log.Printf("Error copying into file %s (%s)", full, err)
			objAndInfo.obj.Close()
			continue
		}
		objAndInfo.obj.Close()
	}
	return nil
}
