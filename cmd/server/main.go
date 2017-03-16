package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/arschles/goprox/gen"
	"github.com/arschles/goprox/storage"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
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
