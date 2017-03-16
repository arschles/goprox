package storage

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/arschles/assert"
	"github.com/arschles/goprox/files"
	"github.com/arschles/goprox/tests"
)

func TestUntarToDisk(t *testing.T) {
	tmpDir, err := tests.CreateTempDir()
	t.Logf("using temp dir %s", tmpDir)
	assert.NoErr(t, err)
	defer os.RemoveAll(tmpDir)

	testDataDir, err := tests.DataDir()
	assert.NoErr(t, err)

	testDataFiles, err := files.List(context.Background(), testDataDir)
	assert.NoErr(t, err)
	buf := new(bytes.Buffer)
	assert.NoErr(t, archiveFiles(testDataDir, buf, testDataFiles...))
	assert.NoErr(t, UntarToDisk(buf, tmpDir))
	pathSet := tests.ExpectedDataSet()
	numFound := 0
	fwErr := filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(tmpDir, path)
		assert.NoErr(t, err)
		if _, ok := pathSet[rel]; !ok {
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
