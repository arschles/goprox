package main

import (
	"context"
	"io/ioutil"

	"github.com/arschles/goprox/gen"
	"github.com/arschles/goprox/storage"
)

func goGet(ctx context.Context, fetcher storage.Fetcher, meta *gen.PackageMeta) (*gen.FullPackage, error) {
	tarball, err := fetcher.GetContents(ctx, meta.Name, meta.Version)
	if err != nil {
		return nil, err
	}
	// TODO: stream this down to the client
	bytes, err := ioutil.ReadAll(tarball)
	if err != nil {
		return nil, err
	}
	return &gen.FullPackage{
		Metadata: meta,
		Payload:  bytes,
	}, nil
}
