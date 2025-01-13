/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cresplanex/bloader/internal/config"
)

// storeObjectListCmd represents the storeObjectList command
var storeObjectListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all object keys in the specified bucket",
	Long: `This command lists all objects in the specified bucket.
It sends a request to the storage server to retrieve the list of objects and prints them to the console.
For example:

bloader store object list --bucket 1234`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		b, err := ctr.Store.ListObjects(bucketID)
		if err != nil {
			color.Red(fmt.Sprintf("Error: %s", err))
			return
		}
		if len(b) == 0 {
			color.Yellow("No objects found")
			return
		}
		color.Green(fmt.Sprintf("Found %d objects:", len(b)))
		for _, o := range b {
			fmt.Println(o)
		}
	},
}

func init() {
	storeObjectCmd.AddCommand(storeObjectListCmd)
}
