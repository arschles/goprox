package tests

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	expectedSet = map[string]struct{}{
		"dir1/file1":                struct{}{},
		"dir2/dir1/dir1/file1":      struct{}{},
		"dir2/dir1/dir1/dir1/file1": struct{}{},
		"file1":                     struct{}{},
	}

	goPath string
)

func init() {
	goPath = os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatalf("GOPATH env var not found, cannot continue")
	}
}

// ExpectedDataSlice returns a slice of the files in the testdata directory
func ExpectedDataSlice() []string {
	ret := make([]string, len(expectedSet))
	i := 0
	for fname := range expectedSet {
		ret[i] = fname
		i++
	}
	return ret
}

// ExpectedDataSet returns a set of the files in the testdata directory
func ExpectedDataSet() map[string]struct{} {
	return expectedSet
}

// DataDir gets the absolute location on disk to the testdata directory
func DataDir() (string, error) {
	dir, err := filepath.Abs(fmt.Sprintf("%s/src/github.com/arschles/goprox/testdata", goPath))
	if err != nil {
		return "", err
	}
	return dir, nil
}
