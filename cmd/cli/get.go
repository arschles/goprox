package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"path/filepath"

	"github.com/arschles/goprox/gen"
	"github.com/arschles/goprox/storage"
	"google.golang.org/grpc"
)

func get(out io.Writer, conn *grpc.ClientConn, name, version string) error {
	cl := gen.NewGoProxDClient(conn)
	pkg, err := cl.GoGet(context.Background(), &gen.PackageMeta{Name: name, Version: version})
	if err != nil {
		return err
	}

	log.Printf("got package %s@%s", pkg.Metadata.Name, pkg.Metadata.Version)
	untarTo := filepath.Join("vendor", pkg.Metadata.Name)
	if err := storage.UntarToDisk(bytes.NewBuffer(pkg.Payload), untarTo); err != nil {
		return err
	}
	log.Printf("package written to %s", untarTo)
	return nil
}
