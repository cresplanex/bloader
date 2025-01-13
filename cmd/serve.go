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

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"sv"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		fmt.Println("serve called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
