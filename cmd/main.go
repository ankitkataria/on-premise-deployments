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
}
