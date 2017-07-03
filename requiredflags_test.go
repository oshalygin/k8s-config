package main

import (
	"testing"
)

func Test_checkRequiredFlags_ShouldThrowAnErrorIfImageAndTagAreNotPassed(t *testing.T) {
	imageTag := ""
	configurationType := ""
	image := ""
	filePath := "deployment.yaml"

	err := checkRequiredFlags(configurationType, image, filePath, imageTag)
	if err == nil {
		t.Errorf("should throw an error if both image and filepath are empty")
	}
}

func Test_checkRequiredFlags_ShouldThrowAnErrorIfTheFilePathWasNotSet(t *testing.T) {
	imageTag := "1.0.0"
	configurationType := ""
	image := ""
	filePath := ""

	err := checkRequiredFlags(configurationType, image, imageTag, filePath)
	if err == nil {
		t.Errorf("should throw an error if the file path was not set")
	}

}

func Test_checkRequiredFlags_ShouldThrowAnErrorIfAnythingButDeploymentIsSetToType(t *testing.T) {
	imageTag := "1.0.0"
	configurationType := "secret"
	image := ""
	filePath := "secret.yaml"

	err := checkRequiredFlags(configurationType, image, imageTag, filePath)
	if err == nil {
		t.Errorf("should throw an error if anything but deployment is set as the type")
	}
}

func Test_checkRequiredFlags_ShouldNotThrowAnErrorIfTheMinimalConfigurationIsPassedIn(t *testing.T) {
	imageTag := "1.0.0"
	configurationType := "deployment"
	image := ""
	filePath := "deployment.yaml"

	err := checkRequiredFlags(configurationType, image, imageTag, filePath)
	if err != nil {
		t.Errorf("should succeed on a valid arguments being passed in")
	}
}

func Test_checkRequiredFlags_ShouldNotThrowAnErrorIfTheFullConfigurationIsPassedIn(t *testing.T) {
	imageTag := ""
	configurationType := "deployment"
	image := "oshalygin/my-web-site:1.0.0"
	filePath := "deployment.yaml"

	err := checkRequiredFlags(configurationType, image, imageTag, filePath)
	if err != nil {
		t.Errorf("should succeed on a valid arguments being passed in")
	}
}

func Test_checkRequiredFlags_ShouldThrowAnErrorIfTheImageDoesNotContainATag(t *testing.T) {
	imageTag := ""
	configurationType := "deployment"
	image := "oshalygin/my-web-site"
	filePath := "deployment.yaml"

	err := checkRequiredFlags(configurationType, image, imageTag, filePath)
	if err == nil {
		t.Errorf("should throw an error if the tag is missing from the image")
	}
}
