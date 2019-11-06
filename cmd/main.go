package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	// Initializing logger
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	// Initializing commands
	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringSliceP("users", "u", make([]string, 0), "Set users for config")
	addCmd.Flags().StringSliceP("clusters", "t", make([]string, 0), "Set clusters used for client")
	addCmd.Flags().StringSliceP("contexts", "x", make([]string, 0), "Set contexts for application deployment for the client")
}
