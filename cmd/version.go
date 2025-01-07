package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of beauty-github-activity",
	Long:  `Prints the version number of beauty-github-activity`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("beauty-github-activity version:", version)
	},
}
