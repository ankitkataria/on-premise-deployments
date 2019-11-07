package deploy

import (
	"os"
	"strings"

	"github.com/ankitkataria/on-premise-deployments/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create deployments to client cluster",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("[*] Starting deployment")

		configPath, image, replicas, deploymentName := GetClusterParams(cmd, args)

		utils.CreateDeployment(configPath, deploymentName, image, replicas)

		os.Exit(0)
	},
}

func init() {
	CreateCmd.Flags().StringP("image", "i", "", "Set image to be deployed over the clusters")
	CreateCmd.Flags().Int32P("replicas", "r", 0, "Set number of replicas required over the cluster")
	CreateCmd.MarkFlagRequired("image")
	CreateCmd.MarkFlagRequired("replicas")
}

func GetClusterParams(cmd *cobra.Command, args []string) (string, string, int32, string) {
	configPath, _ := cmd.Flags().GetString("cluster-config")
	image, _ := cmd.Flags().GetString("image")
	replicas, _ := cmd.Flags().GetInt32("replicas")

	if len(args) == 0 {
		log.Error("No deployment name specified")
		os.Exit(1)
	}

	deploymentName := strings.Join(args, "-")

	return configPath, image, replicas, deploymentName
}
