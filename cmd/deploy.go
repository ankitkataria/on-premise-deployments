package main

import (
	"os"

	"github.com/ankitkataria/on-premise-deployments/cmd/deploy"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy <operation>",
	Short: "Manage applications on required client clusters",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	// Add deployment relation commands
	deployCmd.AddCommand(deploy.CreateCmd)
	deployCmd.AddCommand(deploy.UpdateCmd)
	deployCmd.AddCommand(deploy.CheckCmd)

	deployCmd.PersistentFlags().StringP("cluster-config", "y", "./contexts/minikube-prod.yml", "Set k8s config file, default is contexts/minikube-prod.yml")
}
