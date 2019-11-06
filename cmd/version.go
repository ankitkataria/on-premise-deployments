package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of the current build On-Premise-Deployment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initial Build")
		os.Exit(0)
	},
}
