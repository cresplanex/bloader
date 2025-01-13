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

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh access token",
	Long: `This command refreshes the access token for the application.
It reads the refresh token from the configuration file and sends 
a request to the authorization server to get a new access token.`,
	Run: func(_ *cobra.Command, args []string) {
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
		yellow := color.New(color.FgYellow).SprintFunc()
		fmt.Println(yellow("Token has expired. Refreshing token..."))
		if v, ok := ctr.AuthenticatorContainer.Container[target]; ok {
			if err := (*v).Refresh(ctr.Ctx, ctr.Store); err != nil {
				red := color.New(color.FgRed).SprintFunc()
				fmt.Println(red(fmt.Sprintf("Failed to refresh token: %v", err)))
				fmt.Println("You may need to re-authenticate, if want to access the credential API.")
			} else {
				green := color.New(color.FgGreen).SprintFunc()
				fmt.Println(green("Successfully refreshed token"))
			}
		} else {
			color.Red("Auth setting not found")
			return
		}
	},
}

func init() {
	authCmd.AddCommand(refreshCmd)
}
