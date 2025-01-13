/*
Copyright © 2024 cresplanex <open-source-github@cresplanex.com>
*/
package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/utils"
)

var outputClearAll bool

// outputClearCmd represents the outputClear command
var outputClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the output file",
	Long: `This command clears the output.
It removes all the output file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		if !outputClearAll && len(outputIDs) == 0 {
			color.Yellow("Please specify the output ID to clear or use the --all flag to clear all the outputs")
			return
		}

		for _, o := range ctr.Config.Outputs {
			if !outputClearAll && !utils.Contains(outputIDs, o.ID) {
				continue
			}
			for _, v := range o.Values {
				if v.Env == ctr.Config.Env {
					if err := os.RemoveAll(v.BasePath); err != nil {
						color.Red("Failed to clear the output: %w", err)
						return
					}
				}
			}
		}
		color.Green("Output files cleared successfully")
	},
}

func init() {
	outputCmd.AddCommand(outputClearCmd)

	outputClearCmd.Flags().StringSliceVarP(&outputIDs, "id", "i", []string{}, "ID of the output to clear")
	outputClearCmd.Flags().BoolVarP(&outputClearAll, "all", "A", false, "Clear all the outputs")
}
