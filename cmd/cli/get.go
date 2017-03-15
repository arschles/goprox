package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"path/filepath"

	"github.com/arschles/goprox/gen"
	"github.com/arschles/goprox/storage"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func newGetCommand(out io.Writer, conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get PACKAGE VERSION",
		Short: "Download a package to your vendor directory",
		Long:  "Download PACKAGE to your vendor directory at the given VERSION",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("package & version are required")
			}
			name, version := args[0], args[1]
			return get(out, conn, name, version)
		},
	}

	return cmd
}

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
