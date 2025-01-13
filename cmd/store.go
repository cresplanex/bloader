/*
Copyright © 2024 hayashi kenta <k.hayashi@cresplanex.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:     "store",
	Aliases: []string{"st"},
	Short:   "Perform store management in client cli",
	Long: `It operates the store in the client cli and can create/delete buckets, 
backup, put, etc.`,
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
