/*
Copyright © 2024 cresplanex <open-source-github@cresplanex.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/container"
	"github.com/cresplanex/bloader/internal/utils"
)

var ctr = container.NewContainer()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bloader",
	Short: "The tool for load testing",
	Long: `This tool is used to perform load testing.
It sends requests to the specified server and measures the response time.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctr.Ctx = context.Background()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	defer ctr.Close()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().
		StringP(
			"config",
			"c",
			"",
			`config file (default is ./bloader.yaml, $HOME/bloader.yaml, 
			or /etc/bloader/bloader.yaml)`,
		)
	if err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")); err != nil {
		fmt.Printf("Error binding flag: %v\n", err)
		os.Exit(1)
	}
}

func initConfig() {

	if shouldSkipConfig() {
		return
	}

	configFile := viper.GetString("config")
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Failed to get home directory: %v\n", err)
			os.Exit(1)
		}
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Failed to get current directory: %v\n", err)
			os.Exit(1)
		}
		viper.AddConfigPath(currentDir)
		viper.AddConfigPath(homeDir)
		viper.AddConfigPath("/etc/bloader")
		viper.SetConfigName("bloader")
		// viper.SetConfigType("yaml")
	}

	// Load environment variables
	viper.AutomaticEnv()
	// Prefix for environment variables
	viper.SetEnvPrefix("BLOADER")
	// ex. "BLOADER_SERVER_PORT" -> "server.port"
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Load config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	var cfgForOverride config.ForOverride
	if err := viper.Unmarshal(&cfgForOverride, func(m *mapstructure.DecoderConfig) {
		m.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToSliceHookFunc(","),
		)
	}); err != nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		os.Exit(1)
	}
	validForOverride, err := cfgForOverride.Validate()
	if err != nil {
		fmt.Printf("Error validating config: %v\n", err)
		os.Exit(1)
	}
	for _, override := range validForOverride.Override {
		if !override.EnabledEnv.All && !utils.Contains(override.EnabledEnv.Values, validForOverride.Env) {
			continue
		}
		switch override.Type {
		case config.OverrideTypeStatic:
			if err := config.SetNestedValue(viper.GetViper(), override.Key, override.Value); err != nil {
				fmt.Printf("failed to set nested value: %v\n", err)
				os.Exit(1)
			}
		case config.OverrideTypeFile:
			if override.Partial {
				f, err := os.Open(override.Path)
				if err != nil {
					fmt.Printf("failed to load file: %v\n", err)
					os.Exit(1)
				}
				defer f.Close()
				overrideMap := make(map[string]any)
				switch override.FileType {
				case config.OverrideFileTypesYAML:
					decoder := yaml.NewDecoder(f)
					if err := decoder.Decode(&overrideMap); err != nil {
						fmt.Printf("failed to decode file: %v\n", err)
						os.Exit(1)
					}
				}
				for _, v := range override.Vars {
					value := config.GetNestedValueFromMap(overrideMap, v.Value)
					if err := config.SetNestedValue(viper.GetViper(), v.Key, value); err != nil {
						fmt.Printf("failed to set nested value: %v\n", err)
						os.Exit(1)
					}
				}
			} else {
				f, err := os.Open(override.Path)
				if err != nil {
					fmt.Printf("failed to load file: %v\n", err)
					os.Exit(1)
				}
				defer f.Close()
				overrideMap := make(map[string]any)
				switch override.FileType {
				case config.OverrideFileTypesYAML:
					decoder := yaml.NewDecoder(f)
					if err := decoder.Decode(&overrideMap); err != nil {
						fmt.Printf("failed to decode file: %v\n", err)
						os.Exit(1)
					}
				}
				if err := viper.MergeConfigMap(overrideMap); err != nil {
					fmt.Printf("failed to merge config: %v\n", err)
					os.Exit(1)
				}
			}
		}
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		os.Exit(1)
	}
	validCfg, err := cfg.Validate()
	if err != nil {
		fmt.Printf("Error validating config: %v\n", err)
		os.Exit(1)
	}

	if err := ctr.Init(validCfg); err != nil {
		fmt.Printf("Error initializing container: %v\n", err)
		os.Exit(1)
	}

	// for k, v := range ctr.AuthenticatorContainer.Container {
	// 	if expired := (*v).IsExpired(ctr.Ctx, ctr.Store); expired {
	// 		yellow := color.New(color.FgYellow).SprintFunc()
	// 		fmt.Printf(yellow("Token for %s has expired. Refreshing token...\n"), k)
	// 		if err := (*v).Refresh(ctr.Ctx, ctr.Store); err != nil {
	// 			red := color.New(color.FgRed).SprintFunc()
	// 			fmt.Printf(red("Failed to refresh token: %v\n"), err)
	// 			fmt.Printf("You may need to re-authenticate, if want to access the credential API.\n")
	// 		} else {
	// 			green := color.New(color.FgGreen).SprintFunc()
	// 			fmt.Printf(green("Successfully refreshed token for %s\n"), k)
	// 		}
	// 	}
	// }
}

func shouldSkipConfig() bool {
	// Retrieve the subcommand
	// cmd, _, err := rootCmd.Find(os.Args[1:])
	// if err != nil {
	// 	return false
	// }

	// Ignore the following subcommands
	commandsToSkip := map[string]struct{}{
		"__complete": {}, // This is a special command for shell completion
		"completion": {},
		"version":    {},
		"help":       {},
	}

	_, ok := commandsToSkip[os.Args[1]]
	return ok
}
