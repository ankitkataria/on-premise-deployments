package utils

import (
	log "github.com/sirupsen/logrus"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typedv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateDeployment(configPath string, deploymentName string, image string, replicas int32) {
	deploymentsClient := getDeploymentClient(configPath)
	deployment := getDeployment(deploymentName, image, replicas)
	result, err := deploymentsClient.Create(&deployment)

	if err != nil {
		panic(err)
	}

	log.Info("Created deployment: ", result.GetObjectMeta().GetName())
}

func UpdateImage(configPath string, deploymentName string, image string) {
	deploymentsClient := getDeploymentClient(configPath)
	result, err := deploymentsClient.Get(deploymentName, metav1.GetOptions{})

	if err != nil {
		panic(err)
	}

	result.Spec.Template.Spec.Containers[0].Image = image

	_, updateErr := deploymentsClient.Update(result)

	if updateErr != nil {
		panic(err)
	}
}

func UpdateReplicas(configPath string, deploymentName string, replicas int32) {
	deploymentsClient := getDeploymentClient(configPath)
	result, err := deploymentsClient.Get(deploymentName, metav1.GetOptions{})

	if err != nil {
		panic(err)
	}

	result.Spec.Replicas = int32Ptr(replicas)

	_, updateErr := deploymentsClient.Update(result)

	if updateErr != nil {
		panic(err)
	}
}

func getDeploymentClient(configPath string) typedv1.DeploymentInterface {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	return deploymentsClient
}

func getDeployment(name string, image string, replicas int32) appsv1.Deployment {
	deployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment
}

func int32Ptr(i int32) *int32 { return &i }
