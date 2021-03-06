package storage

import (
	"context"
	"os"

	"github.com/arschles/goprox/files"
)

// ArchiveToDisk creates an archive of the contents of directory and writes that tarball
// to the file at target. It creates target if it didn't already exist
func ArchiveToDisk(ctx context.Context, directory, target string) error {
	files, err := files.List(ctx, directory) // TODO: excludes
	if err != nil {
		return err
	}

	fd, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	if err := archiveFiles(directory, fd, files...); err != nil {
		return err
	}
	return nil
}
