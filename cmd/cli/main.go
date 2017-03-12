package main

import (
	"context"
	"log"

	"github.com/arschles/goprox/gen"
	"google.golang.org/grpc"
)

const (
	appName = "goprox"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cl := gen.NewGoProxDClient(conn)
	pkg, err := cl.GoGet(context.Background(), &gen.PackageMeta{
		Name:    "github.com/arschles/kubehttpbin",
		Version: "HEAD",
	})
	if err != nil {
		log.Fatal(err)
	}

	// conn, err := grpc.Dial(conf.AdminHostString, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("error dialing server (%s)", err)
	// }
	// defer conn.Close()
}
