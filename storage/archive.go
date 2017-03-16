package storage

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// archiveFiles produces an archive of all the files in files.
// It reads each file that lives under prefix and stores it in an archive as the path in the
// slice
func archiveFiles(prefix string, out io.Writer, files ...string) error {
	tw := tar.NewWriter(out)
	for _, file := range files {
		fileBytes, err := ioutil.ReadFile(filepath.Join(prefix, file))
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

	return nil
}
