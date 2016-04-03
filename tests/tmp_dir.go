package tests

import (
	"io/ioutil"
)

const (
	tmpDirPrefix = "goproxtest"
)

// CreateTempDir creates a temp dir for use with a single test and returns it
func CreateTempDir() (string, error) {
	if _, err := ioutil.TempDir("", tmpDirPrefix); err != nil {
		return "", err
	}
	td, err := ioutil.TempDir("", tmpDirPrefix+"/")
	if err != nil {
		return "", nil
	}
	return td, nil
}
