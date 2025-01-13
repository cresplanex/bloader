/*
Copyright © 2024 cresplanex <open-source-github@cresplanex.com>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cresplanex/bloader/internal/config"
)

// storeListCmd represents the storeList command
var storeListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List a list of buckets in the store.",
	Long:    `You can view a list of buckets currently held in the store.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		buckets, err := ctr.Store.ListBuckets()
		if err != nil {
			color.Red("Failed to list the store: %w", err)
			return
		}
		if len(buckets) == 0 {
			color.Yellow("No buckets found in the store")
			return
		}
		color.Green("Buckets in the store:")
		for _, bucket := range buckets {
			fmt.Println(bucket)
		}
	},
}

func init() {
	storeCmd.AddCommand(storeListCmd)
}
