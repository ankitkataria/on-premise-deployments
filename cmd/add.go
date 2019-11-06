package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add and configure new clients",
	Long: `This command creates a new client kubernetes config. It expects the user to enter
	the certificates and permissions for users to access the client's kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		clientName := strings.Join(args, "-")
		f, err := os.OpenFile("./contexts/" + clientName + ".yml", os.O_CREATE, 0755)

		if err != nil {
			log.Error(err)
			log.Error("Could not create new client")
			os.Exit(0)
		}

		defer f.Close()

		log.Info("[+] Created empty client config")
		os.Exit(0)
	},
}
