package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tiger",
	Short: "On premise deployments",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
