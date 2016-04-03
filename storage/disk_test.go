package storage

import (
	"archive/tar"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/arschles/assert"
	"github.com/arschles/goprox/tests"
)

func TestUntarToDisk(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "goproxtest")
	assert.NoErr(t, err)
	defer func() {
		if rerr := os.RemoveAll(tmpDir); rerr != nil {
			log.Printf("Error removing temp dir %s (%s)", tmpDir, rerr)
		}
	}()
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	testDataDir, err := tests.DataDir()
	assert.NoErr(t, err)
	files := tests.ExpectedDataSlice()
	for _, file := range files {
		absFile := filepath.Join(testDataDir, file)
		fBytes, err := ioutil.ReadFile(absFile)
		if err != nil {
			t.Errorf("reading file %s (%s)", file, err)
			continue
		}
		hdr := &tar.Header{
			Name: file,
			Mode: int64(os.ModePerm),
			Size: int64(len(fBytes)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			t.Errorf("writing tar header for %s", file)
			continue
		}
		if _, err := tw.Write(fBytes); err != nil {
			t.Errorf("writing bytes for %s (%s)", file, err)
			continue
		}
	}
	assert.NoErr(t, UntarToDisk(buf, tmpDir))
	pathSet := tests.ExpectedDataSet()
	numFound := 0
	fwErr := filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if _, ok := pathSet[path]; !ok {
			t.Errorf("unexpected path found: %s", path)
			return nil
		}
		numFound++
		return nil
	})
	assert.NoErr(t, fwErr)
	if numFound != len(pathSet) {
		t.Errorf("found %d paths, expected %d", numFound, len(pathSet))
		return
	}

}
