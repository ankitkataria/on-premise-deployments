package main

import (
	"os"
	"strings"

	"github.com/ankitkataria/on-premise-deployments/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add and configure new clients",
	Long: `This command creates a new client kubernetes config. It expects the user to enter
	the certificates and permissions for users to access the client's kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Error("No client-name specified")
			os.Exit(1)
		}

		clientName := strings.Join(args, "-")
		f, err := os.Create("./contexts/" + clientName + "-prod.yml")
		checkFileError(err)

		configContents := parseForContent(cmd)
		_, err = f.WriteString(configContents)

		defer f.Close()

		log.Info("[+] Created empty client config")
		os.Exit(0)
	},
}

func parseForContent(cmd *cobra.Command) string {
	users, _ := cmd.Flags().GetStringSlice("users")
	contexts, _ := cmd.Flags().GetStringSlice("contexts")
	clusters, _ := cmd.Flags().GetStringSlice("clusters")
	config := utils.Config{}
	config.Create(users, clusters, contexts)
	configContents, _ := yaml.Marshal(&config)

	return string(configContents)
}

func checkFileError(err error) {
	if err != nil {
		log.Error(err)
		log.Error("Could not create new client")
		os.Exit(1)
	}
}

func init() {
	addCmd.Flags().StringSliceP("users", "u", make([]string, 0), "Set users for config")
	addCmd.Flags().StringSliceP("clusters", "t", make([]string, 0), "Set clusters used for client")
	addCmd.Flags().StringSliceP("contexts", "x", make([]string, 0), "Set contexts for application deployment for the client")
}
