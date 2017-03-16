package storage

import (
	"context"
	"os"
	"path/filepath"

	"github.com/arschles/goprox/logs"
)

// DiskExistenceChecker is an ExistenceChecker that reads from the Gopath given.
// It will not change the gopath in any way
type DiskExistenceChecker struct {
	Gopath string
}

// Exists is the ExistenceChecker implementation
func (d DiskExistenceChecker) Exists(ctx context.Context, pkgName, version string) (bool, error) {
	logger := logs.FromContext(ctx)
	fullPath := filepath.Join(d.Gopath, "src", pkgName)
	logger.Printf("checking if %s exists", fullPath)
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		logger.Printf("%s doesn't exist", fullPath)
		return false, nil
	}
	return true, err
}
