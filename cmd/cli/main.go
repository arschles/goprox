package main

import (
	"bytes"
	"context"
	"log"
	"path/filepath"

	"github.com/arschles/goprox/gen"
	"github.com/arschles/goprox/storage"
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

	log.Printf("got package %s@%s", pkg.Metadata.Name, pkg.Metadata.Version)
	untarTo := filepath.Join("vendor", pkg.Metadata.Name)
	if err := storage.UntarToDisk(bytes.NewBuffer(pkg.Payload), untarTo); err != nil {
		log.Fatal(err)
	}
	log.Printf("package written to %s", untarTo)

	// conn, err := grpc.Dial(conf.AdminHostString, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("error dialing server (%s)", err)
	// }
	// defer conn.Close()
}
