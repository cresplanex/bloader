/*
Copyright © 2024 hayashi kenta <k.hayashi@cresplanex.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var authID string

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Involves authentication",
	Long: `This command involves authentication.
It is used to authenticate the user, refresh the token, etc.`,
}

func init() {
	rootCmd.AddCommand(authCmd)

	loginCmd.PersistentFlags().
		StringVarP(
			&authID,
			"id",
			"i",
			"",
			`ID of the auth setting. If not provided, a default auth setting will be used.`,
		)
}
