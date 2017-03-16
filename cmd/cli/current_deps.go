package main

import (
	"github.com/arschles/goprox/cmd/cli/deps"
	"github.com/spf13/cobra"
)

type currentDepsCmd struct {
	depsFileName string
}

func newCurrentDepsCommand() *cobra.Command {
	cDepsCmd := currentDepsCmd{}
	cmd := &cobra.Command{
		Use:   "current-deps",
		Short: "Print out a list of the current dependencies listed in the dependencies file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cDepsCmd.run()
		},
	}
	cmd.Flags().StringVarP(
		&cDepsCmd.depsFileName,
		"file",
		"f",
		deps.DefaultFileName,
		"The dependencies file name to parse",
	)

	return cmd
}

func (c currentDepsCmd) run() error {
	debugf("Parsing file %s", c.depsFileName)
	depsFile, err := deps.ParseFile(c.depsFileName)
	if err != nil {
		return err

	}
	imports := depsFile.Import
	printf("Listing %d dependencies", len(imports))

	for i, imp := range imports {
		ver := imp.Version
		if ver == "" {
			ver = "(unset)"
		}
		printf("%d: %s@%s", i, imp.Package, ver)
	}
	return nil
}
