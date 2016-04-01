package repo

import (
	"fmt"
	"net/url"
)

// Info contains information about a repository
type Info struct {
	u            *url.URL
	importPrefix string
	pkg          string
}

// InfoFromURL creates a new Info struct from a given incoming URL
func InfoFromURL(u *url.URL) (*Info, error) {
	pkg := u.Path[1:]
	prefix := fmt.Sprintf("%s/%s", u.Host, pkg)
	return &Info{u: u, importPrefix: prefix, pkg: pkg}, nil
}

// ImportPrefix returns the import prefix
func (i *Info) ImportPrefix() string {
	return i.importPrefix
}

// Package returns the go package represented by this repository
func (i *Info) Package() string {
	return i.pkg
}
