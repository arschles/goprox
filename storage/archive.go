package storage

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// archiveFiles produces an archive of all the files in files.
// It reads each file that lives under prefix and stores it in an archive as the path in the
// slice
func archiveFiles(prefix string, files ...string) (io.Reader, error) {
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
