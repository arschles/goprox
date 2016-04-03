package storage

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/arschles/assert"
)

var (
	goPath           string
	testDataDirFiles = []string{
		"dir1/file1",
		"dir2/dir1/dir1/file1",
		"dir2/dir1/dir1/dir1/file1",
		"file1",
	}
)

func init() {
	goPath = os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatalf("GOPATH env var not found, cannot continue")
	}
}

func testDataDirFilesSet() map[string]struct{} {
	ret := make(map[string]struct{})
	for _, file := range testDataDirFiles {
		ret[file] = struct{}{}
	}
	return ret
}

func getTestDataDir() (string, error) {
	dir, err := filepath.Abs(fmt.Sprintf("%s/src/github.com/arschles/goprox/testdata", goPath))
	if err != nil {
		return "", err
	}
	return dir, nil
}

func TestGetFiles(t *testing.T) {
	dir, err := getTestDataDir()
	assert.NoErr(t, err)
	files, err := getFiles(dir)
	assert.NoErr(t, err)
	expected := testDataDirFiles
	set := map[string]struct{}{}
	for _, file := range files {
		set[file] = struct{}{}
	}
	for _, ex := range expected {
		_, ok := set[ex]
		assert.True(t, ok, "file %s was not found", ex)
	}
}

func TestTarFiles(t *testing.T) {
	dir, err := getTestDataDir()
	assert.NoErr(t, err)
	rdr, err := tarFiles(dir, testDataDirFiles)
	assert.NoErr(t, err)
	tr := tar.NewReader(rdr)
	set := testDataDirFilesSet()
	numFound := 0
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Errorf("reading next record in archive (%s)", err)
			continue
		}
		_, found := set[hdr.Name]
		if !found {
			t.Errorf("unknown file %s", hdr.Name)
			continue
		}
		numFound++
	}
	if numFound != len(testDataDirFiles) {
		t.Fatalf("found only %d of %d files in the testdata dir", len(testDataDirFiles), numFound)
	}

}
