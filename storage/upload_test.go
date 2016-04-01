package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/arschles/assert"
)

var (
	goPath string
)

func init() {
	goPath = os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatalf("GOPATH env var not found, cannot continue")
	}
}

func TestGetFiles(t *testing.T) {
	dir, err := filepath.Abs(fmt.Sprintf("%s/src/github.com/arschles/goprox/testdata", goPath))
	assert.NoErr(t, err)
	files, err := getFiles(dir)
	assert.NoErr(t, err)
	log.Printf("found files %s", files)
	expected := []string{
		"dir1/file1",
		"dir2/dir1/dir1/file1",
		"dir2/dir1/dir1/dir1/file1",
		"file1",
	}
	set := map[string]struct{}{}
	for _, file := range files {
		set[file] = struct{}{}
	}
	for _, ex := range expected {
		_, ok := set[ex]
		assert.True(t, ok, "file %s was not found", ex)
	}
}
