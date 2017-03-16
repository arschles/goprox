package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/arschles/goprox/storage"
	"github.com/spf13/cobra"
)

const packageFileExtension = ".tar"

type errNoSuchPackage struct{ pkg string }

func (e errNoSuchPackage) Error() string {
	return fmt.Sprintf("%s does not exist", e.pkg)
}

func newPackageDirCommand(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "package PACKAGE_PATH",
		Short: "Archive the contents of PACKAGE_PATH to prepare for upload to the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("package path is required")
			}
			path := args[0]
			return packageDir(path)
		},
	}

	return cmd
}

func packageDir(packageName string) error {
	packageName = strings.Trim(packageName, "/")
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return errors.New("GOPATH is not set")
	}
	fullPath := filepath.Join(gopath, "src", packageName)

	printf("checking if %s exists", fullPath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return errNoSuchPackage{pkg: packageName}
	}
	target := strings.Replace(packageName, string(os.PathSeparator), "-", -1) + packageFileExtension
	target = strings.Trim(target, "-")
	printf("Packaging %s into %s", packageName, target)
	if err := storage.ArchiveToDisk(fullPath, target); err != nil {
		return err
	}
	printf("Packaged %s into %s", packageName, target)
	return nil
}
