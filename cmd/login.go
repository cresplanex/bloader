/*
Copyright © 2024 cresplanex <open-source-github@cresplanex.com>
*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cresplanex/bloader/internal/config"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the application",
	Long: `This command logs in to the application.
It sends a request to the authorization server to get an access token.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		target := authID
		if target == "" {
			target = ctr.AuthenticatorContainer.DefaultAuthenticator
		}
		if target == "" {
			color.Red("No auth setting found")
			return
		}
		color.Green("Logging in with %s", target)
		if v, ok := ctr.AuthenticatorContainer.Container[target]; ok {
			if err := (*v).Authenticate(ctr.Ctx, ctr.Store); err != nil {
				color.Red("Failed to login: %w", err)
				return
			}
		} else {
			color.Red("Auth setting not found")
			return
		}
		color.Green("Successfully logged in")
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
}
