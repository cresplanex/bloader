/*
Copyright © 2024 cresplanex <open-source-github@cresplanex.com>
*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cresplanex/bloader/internal/config"
)

// storeClearCmd represents the storeClear command
var storeClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the store",
	Long: `This command clears the store.
It removes all the data from the store.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		if err := ctr.Store.Clear(); err != nil {
			color.Red("Failed to clear the store: %w", err)
			return
		}
		color.Green("Store cleared successfully")
	},
}

func init() {
	storeCmd.AddCommand(storeClearCmd)
}
