package deploy

import (
	"os"

	"github.com/ankitkataria/on-premise-deployments/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update deployments to client cluster",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("[*] Updating deployment")

		configPath, image, replicas, deploymentName := GetClusterParams(cmd, args)

		if image != "" {
			utils.UpdateImage(configPath, deploymentName, image)
		}

		if replicas != 0 {
			utils.UpdateReplicas(configPath, deploymentName, replicas)
		}

		os.Exit(0)
	},
}

func init() {
	UpdateCmd.Flags().StringP("image", "i", "", "Set image to be deployed over the clusters")
	UpdateCmd.Flags().Int32P("replicas", "r", 0, "Set number of replicas required over the cluster")
}
