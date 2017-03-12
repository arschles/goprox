package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/arschles/goprox/cmd/server/lib"
	s3 "github.com/minio/minio-go"
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
	s3Client   *s3.Client
	bucketName string
}

// GetPackages is the AdminServer interface implementation
func (s *server) GoGet(context.Context, *lib.PackageMeta) (*lib.FullPackage, error) {
	return &lib.FullPackage{
		Metadata: &lib.PackageMeta{},
		TarGZ:    nil,
	}, nil
}

func (s *server) UpgradePackage(context.Context, *lib.FullPackage) (*lib.Empty, error) {
	return &lib.Empty{}, nil
}

func (s *server) AddPackage(context.Context, *lib.FullPackage) (*lib.Empty, error) {
	return &lib.Empty{}, nil
}

// StartServer starts the admin server
func StartServer(s3Client *s3.Client, bucketName string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	lib.RegisterGoProxDServer(srv, &server{s3Client: s3Client, bucketName: bucketName})
	return srv.Serve(lis)
}

func main() {
	log.Printf("not yet implemented!")
	os.Exit(1)
}
