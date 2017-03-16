package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/arschles/goprox/gen"
	"github.com/arschles/goprox/logs"
	"github.com/arschles/goprox/storage"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

var (
	errNotYetImplemented = errors.New("not yet implemented")
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

func startServer(
	debug bool,
	fetcher storage.Fetcher,
	checker storage.ExistenceChecker,
	port int,
) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	gen.RegisterGoProxDServer(srv, &server{
		debug:   true,
		fetcher: fetcher,
		checker: checker,
	})
	return srv.Serve(lis)
}

func main() {
	fetcher := storage.DiskFetcher{
		Gopath:   os.Getenv("GOPATH"),
		Excludes: []string{"\\.git/*", "vendor/*"},
	}
	checker := storage.DiskExistenceChecker{
		Gopath: os.Getenv("GOPATH"),
	}
	log.Printf("Serving goproxd on port 8080")
	log.Fatal(startServer(true, fetcher, checker, 8080))
}
