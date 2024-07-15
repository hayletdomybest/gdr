package cmd

import (
	"fmt"
	"os"

	"github.com/hayletdomybest/gdr/internal/rename"
	"github.com/spf13/cobra"
)

func NewRenameCmd() *cobra.Command {
	var projectPath string

	cmd := &cobra.Command{
		Use:   "rename [new_name]",
		Short: "A tool to rename Go module names",
		Long:  `gomodrename is a tool to rename Go module names. It can be used to rename the module name in the go.mod file and all the import paths in the project.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			newName := args[0]
			if projectPath == "" {
				var err error
				projectPath, err = os.Getwd()
				if err != nil {
					return fmt.Errorf("failed to get current directory: %v", err)
				}
			}
			return rename.RenameModule(projectPath, newName)
		},
	}

	cmd.Flags().StringVarP(&projectPath, "path", "p", "", "Path to the project (default is current directory)")

	return cmd
}
