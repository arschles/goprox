package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"io/ioutil"

	"github.com/arschles/goprox/gen"
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
	fetcher storage.Fetcher
}

// GetPackages is the AdminServer interface implementation
func (s *server) GoGet(ctx context.Context, meta *gen.PackageMeta) (*gen.FullPackage, error) {
	tarball, err := s.fetcher.GetContents(meta.Name, meta.Version)
	if err != nil {
		return nil, err
	}
	// TODO: stream this down to the client
	bytes, err := ioutil.ReadAll(tarball)
	if err != nil {
		return nil, err
	}
	return &gen.FullPackage{
		Metadata: meta,
		Payload:  bytes,
	}, nil
}

func (s *server) UpgradePackage(context.Context, *gen.FullPackage) (*gen.Empty, error) {
	return &gen.Empty{}, nil
}

func (s *server) AddPackage(context.Context, *gen.FullPackage) (*gen.Empty, error) {
	return &gen.Empty{}, nil
}

func startServer(fetcher storage.Fetcher, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	gen.RegisterGoProxDServer(srv, &server{fetcher: fetcher})
	return srv.Serve(lis)
}

func main() {
	fetcher := storage.DiskFetcher{
		Gopath:   os.Getenv("GOPATH"),
		Excludes: []string{".git/*"},
	}
	log.Printf("Serving goproxd on port 8080")
	log.Fatal(startServer(fetcher, 8080))
}
