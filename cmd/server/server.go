package main

import (
	"log"

	"github.com/arschles/goprox/gen"
	"github.com/arschles/goprox/logs"
	"github.com/arschles/goprox/storage"
	context "golang.org/x/net/context"
)

// server implements the GoProxDServer interface
type server struct {
	debug   bool
	logger  *log.Logger
	fetcher storage.Fetcher
	checker storage.ExistenceChecker
}

// GetPackages is the AdminServer interface implementation
func (s *server) GoGet(ctx context.Context, meta *gen.PackageMeta) (*gen.FullPackage, error) {
	if s.debug {
		ctx = logs.DebugContext(ctx)
	}
	return goGet(ctx, s.fetcher, meta)
}

func (s *server) UpgradePackage(context.Context, *gen.FullPackage) (*gen.Empty, error) {
	return &gen.Empty{}, nil
}

func (s *server) AddPackage(context.Context, *gen.FullPackage) (*gen.Empty, error) {
	return &gen.Empty{}, nil
}

func (s *server) PackageExists(ctx context.Context, meta *gen.PackageMeta) (*gen.PackageExistsResponse, error) {
	if s.debug {
		ctx = logs.DebugContext(ctx)
	}
	exists, err := s.checker.Exists(ctx, meta.Name, meta.Version)
	if err != nil {
		return nil, err
	}
	return &gen.PackageExistsResponse{Exists: exists, Meta: meta}, nil
}
