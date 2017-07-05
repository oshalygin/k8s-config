package services

import (
	"strings"

	"errors"

	"io/ioutil"

	"github.com/oshalygin/k8s-config/models"
	"gopkg.in/yaml.v2"
)

// UpdateDeploymentConfiguration takes in a file, image or imageTag to provide an updated configuration object
func UpdateDeploymentConfiguration(file []byte, image string, imageTag string, destination string) (string, models.Deployment, error) {
	deployment, err := parseConfigurationFile(file)
	updatedDeployment := models.Deployment{}
	if err != nil {
		return "", models.Deployment{}, err
	}

	originalImage := deployment.Spec.Template.Spec.Containers[0].Image

	if imageTag != "" {
		updatedDeployment, err = updateImageTag(deployment, imageTag)
	} else {
		updatedDeployment, err = updateImage(deployment, image)
	}

	if err != nil {
		return originalImage, models.Deployment{}, err
	}

	if destination != "" {
		err := saveConfiguration(updatedDeployment, destination)
		if err != nil {
			return originalImage, updatedDeployment, err
		}

	}
	return originalImage, updatedDeployment, nil

}

func parseConfigurationFile(file []byte) (configurationFile models.Deployment, err error) {
	err = yaml.Unmarshal(file, &configurationFile)
	return
}

func saveConfiguration(configurationFile models.Deployment, destination string) error {
	data, err := yaml.Marshal(configurationFile)

	if err != nil {
		return errors.New("Could not save the file to the destination")
	}
	err = ioutil.WriteFile(destination, data, 0644)
	if err != nil {
		return errors.New("Could not save the file to the destination")
	}

	return err
}

func updateImageTag(deployment models.Deployment, imageTag string) (models.Deployment, error) {

	multipleContainers := isMultipleContainerDeployment(deployment)

	if multipleContainers {
		return models.Deployment{}, errors.New("Only a single container is supported at this time")
	}

	currentImage := deployment.Spec.Template.Spec.Containers[0].Image
	image := strings.Split(currentImage, ":")[0]
	updatedImage := image + ":" + imageTag

	deployment.Spec.Template.Spec.Containers[0].Image = updatedImage
	return deployment, nil
}

func updateImage(deployment models.Deployment, image string) (models.Deployment, error) {
	multipleContainers := isMultipleContainerDeployment(deployment)

	if multipleContainers {
		return models.Deployment{}, errors.New("Only a single container is supported at this time")
	}

	//TODO: Perform error checking on whether or not it is a valid image
	deployment.Spec.Template.Spec.Containers[0].Image = image
	return deployment, nil
}

func isMultipleContainerDeployment(deployment models.Deployment) bool {
	numberOfContainers := len(deployment.Spec.Template.Spec.Containers)
	if numberOfContainers > 1 {
		return true
	}
	return false
}
