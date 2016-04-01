package storage

import (
	"errors"
	"log"
	"sync"

	s3 "github.com/minio/minio-go"
)

var (
	// ErrNoMoreFiles is the error that's returned when another files is requested in a package, but none are left
	ErrNoMoreFiles = errors.New("no more files in package")
)

// PackageStream is a collection of files that make up a package that are being streamed from S3
type PackageStream struct {
	doneCh chan<- struct{}
	// Files is the unordered collection of files streamed down from S3. This channel will be closed when there are no more files available.
	files       <-chan *File
	packageName string
}

// GetStreamForPackage returns a new PackageStream for the given package name
func GetStreamForPackage(cl *s3.Client, bucketName, packageName string) (*PackageStream, error) {
	// TODO: make this completely tarball based

	doneCh := make(chan struct{})
	objInfoCh := cl.ListObjects(bucketName, packageName, true, doneCh)

	var wg sync.WaitGroup
	filesCh := make(chan *File)
	for objInfo := range objInfoCh {
		wg.Add(1)
		go func(objInfo s3.ObjectInfo) {
			defer wg.Done()
			obj, err := cl.GetObject(bucketName, objInfo.Key)
			if err != nil {
				log.Printf("Error downloading object %s (%s)", objInfo.Key, err)
				return
			}
			filesCh <- &File{Key: objInfo.Key, ReadCloser: obj}
		}(objInfo)
	}
	return &PackageStream{doneCh: doneCh, files: filesCh, packageName: packageName}, nil
}

// NextFile returns the next file in the package. Returns ErrNoMoreFiles if there are none left in object storage
func (o *PackageStream) NextFile() (*File, error) {
	file, ok := <-o.files
	if !ok {
		return nil, ErrNoMoreFiles
	}
	return file, nil
}

// Close stops files from streaming from S3. Always call this when you're done, even if no more files have been streamed
func (o *PackageStream) Close() error {
	close(o.doneCh)
	return nil
}

// PackageName returns the name of the package that is being streamed from object storage
func (o *PackageStream) PackageName() string {
	return o.packageName
}
