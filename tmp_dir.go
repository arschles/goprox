package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	tempDirPrefix = "goprox"
)

func createTempDir() (string, error) {
	if err := os.MkdirAll(fmt.Sprintf("%s/%s/", os.TempDir(), tempDirPrefix), os.ModePerm); err != nil {
		return "", err
	}
	tmpDir, err := ioutil.TempDir(os.TempDir(), "/"+tempDirPrefix+"/")
	if err != nil {
		return "", err
	}
	return tmpDir, nil
}
