package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "gdr",
		Short: "A tool to rename Go module names",
		Long:  `gdr is a tool to rename Go module names. It can be used to rename the module name in the go.mod file and all the import paths in the project.`,
	}

	root.AddCommand(
		NewRenameCmd(),
		NewVersionCmd(),
	)
	return root
}
