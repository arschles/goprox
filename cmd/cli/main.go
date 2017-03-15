package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	appName = "goprox"
)

func newRootCmd(out io.Writer, conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goprox",
		Short: "The Goprox dependency manager",
		// Long:         globalUsage,
		SilenceUsage: true,
		// PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// 	teardown()
		// },
	}
	// p := cmd.PersistentFlags()
	cmd.AddCommand(newGetCommand(out, conn))
	return cmd
}

func newGetCommand(out io.Writer, conn *grpc.ClientConn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get PACKAGE VERSION",
		Short: "download a package to your vendor directory",
		// Long:  createDesc,
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

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cmd := newRootCmd(os.Stdout, conn)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
