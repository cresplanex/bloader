/*
Copyright © 2024 hayashi kenta <k.hayashi@cresplanex.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"conf"},
	Short:   "Print the current configuration",
	Long: `This command prints the current configuration.
It reads the configuration from the configuration file and prints it in YAML format.`,
	Run: func(cmd *cobra.Command, args []string) {
		settings := viper.AllSettings()
		exConfigSettings := map[string]any{}
		for key, value := range settings {
			if key != "config" {
				exConfigSettings[key] = value
			}
		}
		yamlData, err := yaml.Marshal(exConfigSettings)
		if err != nil {
			log.Fatalf("Error converting configuration to YAML: %v", err)
		}
		fmt.Println("Current configuration:")
		fmt.Println(string(yamlData))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
