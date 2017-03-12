package storage

import (
	"io"
)

// Fetcher is the interface to fetch packages from backing storage
type Fetcher interface {
	// GetContents returns tarballed contents of the package at pkgName, at the given version.
	// Some implementations may ignore version
	GetContents(pkgName string, version string) (io.Reader, error)
}
