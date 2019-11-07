package deploy

import (
	"os"

	"github.com/ankitkataria/on-premise-deployments/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete application deployment from client clusters",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("[*] Deleting deployment")

		configPath, _, _, deploymentName := GetClusterParams(cmd, args)

		utils.DeleteDeployment(configPath, deploymentName)

		os.Exit(0)
	},
}
