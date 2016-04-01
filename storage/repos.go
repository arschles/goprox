package storage

import (
	"io"
	"os"
	"path/filepath"
)

// WriteStreamToDisk writes every file in packageStream to disk, under localDir
func WriteStreamToDisk(packageStream *PackageStream, localDir string) error {
	defer packageStream.Close()

	dir := filepath.Join(localDir, packageStream.PackageName())
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// iterate all files and write them to disk
	for {
		file, err := packageStream.NextFile()
		if err == ErrNoMoreFiles {
			return nil
		} else if err != nil {
			return err
		}

		fullPath := filepath.Join(dir, file.Key)
		fd, err := os.Create(fullPath)
		if err != nil {
			fd.Close()
			file.Close()
			return err
		}

		if _, err := io.Copy(fd, file); err != nil {
			fd.Close()
			file.Close()
			return err
		}
		fd.Close()
		file.Close()
	}
}
