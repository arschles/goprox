package storage

import (
	"context"
	"io"
)

// Fetcher is the interface to fetch packages from backing storage
type Fetcher interface {
	// GetContents returns tarballed contents of the package at pkgName, at the given version.
	// Some implementations may ignore version
	GetContents(ctx context.Context, pkgName string, version string) (io.Reader, error)
}
