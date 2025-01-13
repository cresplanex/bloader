/*
Copyright © 2024 cresplanex <open-source-github@cresplanex.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var outputIDs []string

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:     "output",
	Aliases: []string{"out"},
	Short:   "Perform output management in client cli",
	Long:    `It operates the output in the client cli and can clear output file, etc.`,
}

func init() {
	rootCmd.AddCommand(outputCmd)
}
