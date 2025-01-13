/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// slaveCmd represents the slave command
var slaveCmd = &cobra.Command{
	Use:     "slave",
	Aliases: []string{"sl"},
	Short:   "For worker nodes",
	Long: `This command is used to start a worker node.
A worker node is a node that is responsible for running the tasks assigned by the master node.`,
}

func init() {
	rootCmd.AddCommand(slaveCmd)
}
