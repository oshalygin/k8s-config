package services

import (
	"strings"

	"github.com/oshalygin/k8s-config/models"
	"gopkg.in/yaml.v2"
)

// UpdateDeploymentConfiguration takes in a file, image or imageTag to provide an updated configuration object
func UpdateDeploymentConfiguration(file []byte, image string, imageTag string) (models.Deployment, error) {
	original, err := parseConfigurationFile(file)
	if err != nil {
		return models.Deployment{}, err
	}

	return original, nil

}

func parseConfigurationFile(file []byte) (configurationFile models.Deployment, err error) {
	err = yaml.Unmarshal(file, &configurationFile)
	return
}

func updateImageTag(deployment models.Deployment, imageTag string) models.Deployment {
	currentImage := deployment.Spec.Template.Spec.Containers[0].Image
	image := strings.Split(currentImage, ":")[0]
	updatedImage := image + ":" + imageTag

	deployment.Spec.Template.Spec.Containers[0].Image = updatedImage
	return deployment
}
