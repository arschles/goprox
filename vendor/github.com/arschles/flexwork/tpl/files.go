package tpl

import (
	"path/filepath"
	"strings"
)

// Files is an ordered list of filenames that comprise a template.
// The first filename in the list should be the primary template that will
// be rendered, and the remaining ones should be supporting templates like
// block definitions
type Files struct {
	first     string
	list      []string
	mapKeyStr string
}

type filesMapKey string

func (f filesMapKey) String() string {
	return string(f)
}

// NewFiles creates a new Files struct from the given filenames
func NewFiles(file string, files ...string) Files {
	return Files{
		first:     file,
		list:      files,
		mapKeyStr: strings.Join(files, ","),
	}
}

// First returns the first file listed in f
func (f Files) First() string {
	return f.first
}

func (f Files) len() int {
	return len(f.list) + 1
}

func (f Files) absPaths(absPath string) []string {
	var allList []string
	allList = append(allList, f.first)
	allList = append(allList, f.list...)
	ret := make([]string, len(allList))
	for i, fileName := range allList {
		ret[i] = filepath.Join(absPath, fileName)
	}
	return ret
}

func (f Files) mapKey() filesMapKey {
	return filesMapKey(f.mapKeyStr)
}
