package storage

import (
	"io"
	"os"
)

// ArchiveToDisk creates an archive of the contents of directory and writes that tarball
// to the file at target. It creates target if it didn't already exist
func ArchiveToDisk(directory, target string) error {
	files, err := getFiles(directory) // TODO: excludes
	if err != nil {
		return err
	}
	buf, err := archiveFiles(directory, files...)
	if err != nil {
		return err
	}
	fd, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	if _, err := io.Copy(fd, buf); err != nil {
		return err
	}

	return nil
}
