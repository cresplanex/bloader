/*
Copyright © 2024 hayashi kenta <k.hayashi@cresplanex.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the version of the application
	Version = "dev"
	// Commit is the commit hash when the application was built
	Commit = "Unknown"
	// BuildTime is the time when the application was built
	BuildTime = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version",
	Long: `This command displays the version of the application.
It displays the version of the application and the version of the configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Bloader Version: %s\nCommit: %s\nBuild Time: %s\n", Version, Commit, BuildTime)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
