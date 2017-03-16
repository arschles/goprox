package main

import (
	"context"
	"errors"
	"io"

	"github.com/arschles/goprox/gen"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func newExistsCommand(out io.Writer, conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exists PACKAGE VERSION",
		Short: "Check if a PACKAGE exists with the given VERSION",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("package & version are required")
			}
			name, version := args[0], args[1]
			return exists(out, conn, name, version)
		},
	}

	return cmd
}

func exists(out io.Writer, conn *grpc.ClientConn, name, version string) error {
	cl := gen.NewGoProxDClient(conn)
	resp, err := cl.PackageExists(
		context.Background(),
		&gen.PackageMeta{Name: name, Version: version},
	)
	if err != nil {
		return err
	}
	if resp.Exists {
		printf("Package %s@%s exists!", name, version)
	} else {
		printf("Package %s@%s doesn't exist :(", name, version)
	}
	return nil
}
