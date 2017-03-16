package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

func newUpgradeCommand(conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade PACKAGE VERSION",
		Short: "Change the version of PACKAGE to VERSION in the local vendor directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("PACKAGE and VERSION are required")
			}
			pkg, version := args[0], args[1]
			return upgrade(conn, pkg, version)
		},
	}

	return cmd
}

func upgrade(conn *grpc.ClientConn, pkg, version string) error {
	fullPath := filepath.Join("vendor", "src", pkg)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("%s isn't installed", pkg)
	}
	printf("Upgrading %s to version %s", pkg, version)
	return get(conn, pkg, version)
}
