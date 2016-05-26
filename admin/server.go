package admin

import (
	"errors"
	"fmt"
	"net"

	"github.com/arschles/goprox/storage"
	s3 "github.com/minio/minio-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	errNotYetImplemented = errors.New("not yet implemented")
)

// server implements the AdminServer interface
type server struct {
	s3Client   *s3.Client
	bucketName string
}

// GetPackages is the AdminServer interface implementation
func (s *server) GetPackages(ctx context.Context, in *Empty) (*PackageList, error) {
	doneCh := make(chan struct{})
	defer close(doneCh)
	objCh := s.s3Client.ListObjects(s.bucketName, "", false, doneCh)
	packageList := &PackageList{}
	for objInfo := range objCh {
		packageList.Packages = append(packageList.Packages, &Package{
			Name: storage.ReverseName(objInfo.Key),
		})
	}

	return packageList, nil
}

func (s *server) AddPackage(ctx context.Context, fp *FullPackage) (*FullPackage, error) {
	return nil, errNotYetImplemented
}

func (s *server) UpgradePackage(ctx context.Context, fp *FullPackage) (*FullPackage, error) {
	return nil, errNotYetImplemented
}

// StartServer starts the admin server
func StartServer(s3Client *s3.Client, bucketName string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	RegisterAdminServer(srv, &server{s3Client: s3Client, bucketName: bucketName})
	return srv.Serve(lis)
}
