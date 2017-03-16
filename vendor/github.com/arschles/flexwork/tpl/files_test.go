package tpl

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/arschles/assert"
)

var (
	files = NewFiles("file1", "file2", "file3")
)

func TestFilesLen(t *testing.T) {
	assert.Equal(t, files.len(), len(files.list)+1, "length")
}

func TestFilesMapKey(t *testing.T) {
	assert.Equal(t, files.mapKey().String(), strings.Join(files.list, ","), "map key")
}

func TestFilesAbsPaths(t *testing.T) {
	const absPath = "path1"
	absPaths := files.absPaths(absPath)
	assert.Equal(t, len(absPaths), files.len(), "slice length")
	for i := 0; i < files.len(); i++ {
		var fn string
		if i == 0 {
			fn = files.first
		} else {
			fn = files.list[i-1]
		}
		assert.Equal(t, absPaths[i], filepath.Join(absPath, fn), fmt.Sprintf("abs path %d", i))
	}

	files := NewFiles("file1", "file2", "file3")
	absPaths = files.absPaths("a")
	assert.Equal(t, len(absPaths), 3, "number of abs paths")
	assert.Equal(t, absPaths[0], "a/file1", "file 1")
	assert.Equal(t, absPaths[1], "a/file2", "file 2")
	assert.Equal(t, absPaths[2], "a/file3", "file 3")
}
