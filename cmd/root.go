package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tiger",
	Short: "Tiger: On Premise Deployments",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
