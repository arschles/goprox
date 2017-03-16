package files

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	errNilFileInfo = errors.New("nil file info")
)

// List lists and returns names of files under dir, as paths relative to dir.
// Call filepath.Join(dir, file) on each returned file to get the absolute path
func List(dir string, excludes ...string) ([]string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}

	files := []string{}
	if err := filepath.Walk(dir, getWalkFunc(dir, &files, excludes...)); err != nil {
		return nil, err
	}
	return files, nil
}

func isExcluded(name string, excludes ...string) bool {
	for _, exclude := range excludes {
		matched, err := regexp.Match(exclude, []byte(name))
		if err == nil && matched {
			log.Printf("excluding file %s", name)
			return true
		}
	}
	return false
}

func getWalkFunc(baseDir string, files *[]string, excludes ...string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		// skip if this is an exclude
		if isExcluded(path, excludes...) {
			return nil
		}
		log.Printf("walking %s", path)
		if err != nil {
			// TODO: handle this case!
			return nil
		}
		if info == nil {
			// TODO: handle this case!
			return errNilFileInfo
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(baseDir, path)
		if err != nil {
			return err
		}

		*files = append(*files, rel)
		return nil
	}
}
