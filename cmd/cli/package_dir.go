package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/arschles/goprox/storage"
	"github.com/spf13/cobra"
)

const packageFileExtension = ".tar"

func newPackageDirCommand(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "package PACKAGE_PATH",
		Short: "Archive the contents of PACKAGE_PATH to prepare for upload to the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("package path is required")
			}
			path := args[0]
			if _, err := os.Stat(path); os.IsNotExist(err) {
				return fmt.Errorf("%s does not exist on disk", path)
			}
			return packageDir(path)
		},
	}

	return cmd
}

func packageDir(path string) error {
	target := strings.Replace(path, string(os.PathSeparator), "-", -1) + packageFileExtension
	log.Printf("Packaging %s into %s", path, target)
	if err := storage.ArchiveToDisk(path, target); err != nil {
		return err
	}
	log.Printf("Packaged %s into %s", path, target)
	return nil
}
