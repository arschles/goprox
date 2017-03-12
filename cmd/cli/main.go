package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	// "google.golang.org/grpc"
)

const (
	appName = "goprox"
)

func main() {
	log.Printf("hello world!")
	app := cli.NewApp()
	app.Run(os.Args)

	// conn, err := grpc.Dial(conf.AdminHostString, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("error dialing server (%s)", err)
	// }
	// defer conn.Close()
}
