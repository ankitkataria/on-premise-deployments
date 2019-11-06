package main

import (
	"os"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of the current build On-Premise-Deployment",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Initial Build")
		os.Exit(0)
	},
}
