package main

import (
	"testing"
)

func Test_main_ShouldUpdateTheDeploymentConfigurationWithTheNewTag(t *testing.T) {

	//TODO: This will require reading the output location accordingly.
	//TODO: At the moment this serves as a stub for an "integration" test

	filePath = "./test-files/deployment.yaml"
	imageTag = "1.3.3"
	main()
}
