package storage

import (
	"io"
)

// File is a single file in a package repsitory
type File struct {
	Key string
	io.ReadCloser
}
