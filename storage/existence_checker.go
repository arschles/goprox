package storage

import "context"

// ExistenceChecker is the interface to check whether a given package name exists
type ExistenceChecker interface {
	Exists(ctx context.Context, pkgName, version string) (bool, error)
}
