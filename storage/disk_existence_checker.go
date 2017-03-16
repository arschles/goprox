package storage

import (
	"context"
	"os"
	"path/filepath"
)

// DiskExistenceChecker is an ExistenceChecker that reads from the Gopath given.
// It will not change the gopath in any way
type DiskExistenceChecker struct {
	Gopath string
}

// Exists is the ExistenceChecker implementation
func (d DiskExistenceChecker) Exists(ctx context.Context, pkgName, version string) (bool, error) {
	fullPath := filepath.Join("vendor", "src", pkgName)
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
