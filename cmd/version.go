package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of the current build of On-Premise-Deployment",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Initial Build")
		os.Exit(0)
	},
}
