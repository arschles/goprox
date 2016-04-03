package storage

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/arschles/assert"
	"github.com/arschles/goprox/tests"
)

func TestUntarToDisk(t *testing.T) {
	tmpDir, err := tests.CreateTempDir()
	assert.NoErr(t, err)
	defer func() {
		if rerr := os.RemoveAll(tmpDir); rerr != nil {
			log.Printf("Error removing temp dir %s (%s)", tmpDir, rerr)
		}
	}()
	testDataDir, err := tests.DataDir()
	assert.NoErr(t, err)

	testDataFiles, err := getFiles(testDataDir)
	assert.NoErr(t, err)
	rdr, err := tarFiles(testDataDir, testDataFiles)
	assert.NoErr(t, err)
	assert.NoErr(t, UntarToDisk(rdr, tmpDir))
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
