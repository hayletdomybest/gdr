package cmd

import (
	"fmt"

	"github.com/hayletdomybest/gdr/internal"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "version",
		Short: "Print the version of gomodrename",
		Long:  `Print the version of gomodrename.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(internal.Version)
		},
	}
	return root
}
