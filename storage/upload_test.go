package storage

import (
	"archive/tar"
	"io"
	"testing"

	"github.com/arschles/assert"
	"github.com/arschles/goprox/tests"
)

func TestGetFiles(t *testing.T) {
	dir, err := tests.DataDir()
	assert.NoErr(t, err)
	files, err := getFiles(dir)
	assert.NoErr(t, err)
	expected := tests.ExpectedDataSlice()
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
	dir, err := tests.DataDir()
	assert.NoErr(t, err)
	rdr, err := tarFiles(dir, tests.ExpectedDataSlice())
	assert.NoErr(t, err)
	tr := tar.NewReader(rdr)
	set := tests.ExpectedDataSet()
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
	if numFound != len(set) {
		t.Fatalf("found only %d of %d files in the testdata dir", len(tests.ExpectedDataSlice()), numFound)
	}

}
