/*
Copyright © 2024 hayashi kenta <k.hayashi@cresplanex.com>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cresplanex/bloader/internal/config"
)

// storeObjectPutCmd represents the storeObjectPut command
var storeObjectPutCmd = &cobra.Command{
	Use:   "put",
	Short: "Put an object in the specified bucket",
	Long: `This command puts an object in the specified bucket.
It reads the object from the specified file and sends a request to the storage server to store the object.
For example:

bloader store object put --bucket 1234 objectKey objectValue`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		if len(args) < 2 {
			fmt.Println("Please provide the object key and value")
			return
		}
		objKey := args[0]
		objVal := args[1]

		if storeObjectEncrypt != "" {
			encryper, ok := ctr.EncypterContainer[storeObjectEncrypt]
			if !ok {
				color.Red("Encrypt setting not found")
				return
			}
			b, err := encryper.Encrypt([]byte(objVal))
			if err != nil {
				color.Red("Failed to encrypt: %w", err)
				return
			}
			objVal = string(b)
		}

		if err := ctr.Store.PutObject(bucketID, objKey, []byte(objVal)); err != nil {
			color.Red("Failed to put object: %w", err)
			return
		}
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Println(green("Object put"))
	},
}

func init() {
	storeObjectCmd.AddCommand(storeObjectPutCmd)

	storeObjectPutCmd.PersistentFlags().
		StringVarP(
			&storeObjectEncrypt,
			"encrypt",
			"e",
			"",
			`ID of the encryption setting. This is optional.`,
		)
}
