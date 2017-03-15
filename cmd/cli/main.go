package main

import (
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
		Use:          "goprox",
		Short:        "The Goprox dependency manager",
		SilenceUsage: true,
	}
	cmd.AddCommand(newGetCommand(out, conn))
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
