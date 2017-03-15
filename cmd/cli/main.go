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

func newRootCmd(out io.Writer) (*cobra.Command, *grpc.ClientConn) {
	cmd := &cobra.Command{
		Use:          "goprox",
		Short:        "The Goprox dependency manager",
		SilenceUsage: true,
	}
	p := cmd.PersistentFlags()
	hostStr := p.String("host", "localhost:8080", "The host of the goprox server")
	conn, err := grpc.Dial(*hostStr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	cmd.AddCommand(newGetCommand(out, conn))
	cmd.AddCommand(newExistsCommand(out, conn))
	return cmd, conn
}

func main() {
	cmd, conn := newRootCmd(os.Stdout)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	defer conn.Close()
}
