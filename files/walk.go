package files

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/arschles/goprox/logs"
)

var (
	errNilFileInfo = errors.New("nil file info")
)

// List lists and returns names of files under dir, as paths relative to dir.
// Call filepath.Join(dir, file) on each returned file to get the absolute path
func List(ctx context.Context, dir string, excludes ...string) ([]string, error) {
	logger := logs.FromContext(ctx)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logger.Printf("%s doesn't exist (%s)", dir, err)
		return nil, err
	}

	files := []string{}
	if err := filepath.Walk(dir, getWalkFunc(ctx, dir, &files, excludes...)); err != nil {
		return nil, err
	}
	return files, nil
}

func isExcluded(name string, excludes ...string) bool {
	for _, exclude := range excludes {
		matched, err := regexp.Match(exclude, []byte(name))
		if err == nil && matched {
			log.Printf("%s is excluded on %s", name, exclude)
			return true
		}
	}
	return false
}

func getWalkFunc(
	ctx context.Context,
	baseDir string,
	files *[]string,
	excludes ...string,
) filepath.WalkFunc {
	logger := logs.FromContext(ctx)

	return func(path string, info os.FileInfo, err error) error {
		logger.Printf("walking %s", path)
		// skip if this is an exclude
		if isExcluded(path, excludes...) {
			logger.Printf("%s is excluded", path)
			return nil
		}
		if err != nil {
			logger.Printf("walk error (%s)", err)
			// TODO: handle this case!
			return nil
		}
		if info == nil {
			logger.Printf("file info was nil!")
			// TODO: handle this case!
			return errNilFileInfo
		}
		if info.IsDir() {
			logger.Printf("bail out on directory %s", path)
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
