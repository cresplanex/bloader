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

// storeObjectGetCmd represents the storeObjectGet command
var storeObjectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an object from the specified bucket",
	Long: `This command retrieves an object from the specified bucket.
It sends a request to the storage server to retrieve the object and writes it to the specified file.
For example:

bloader store object get --bucket 1234 objectKey`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		if len(args) == 0 {
			fmt.Println("Please provide the object key")
			return
		}
		objKey := args[0]
		objVal, err := ctr.Store.GetObject(bucketID, objKey)
		if err != nil {
			color.Red("Failed to get object: %w", err)
			return
		}
		if len(objVal) == 0 {
			color.Yellow("Object not found")
			return
		}
		if storeObjectEncrypt != "" {
			encryper, ok := ctr.EncypterContainer[storeObjectEncrypt]
			if !ok {
				color.Red("Encrypt setting not found")
				return
			}
			b, err := encryper.Decrypt(string(objVal))
			if err != nil {
				color.Red("Failed to encrypt: %w", err)
				return
			}
			objVal = b
		}
		// green := color.New(color.FgGreen).SprintFunc()
		fmt.Println(string(objVal))
	},
}

func init() {
	storeObjectCmd.AddCommand(storeObjectGetCmd)

	storeObjectGetCmd.PersistentFlags().
		StringVarP(
			&storeObjectEncrypt,
			"encrypt",
			"e",
			"",
			`ID of the encryption setting. This is optional.`,
		)
}
