package main

import (
	"flag"
	"fmt"
	"os"

	"io/ioutil"

	"strings"

	"github.com/fatih/color"
	"github.com/oshalygin/k8s-config/services"
)

var configurationType string
var image string
var imageTag string
var filePath string

func init() {
	flag.StringVar(&configurationType, "type", "deployment", "Kubernetes configuration type, eg: deployment, rc, secret")
	flag.StringVar(&image, "image", "", "Docker image name")
	flag.StringVar(&imageTag, "tag", "", "Docker image tag")
	flag.StringVar(&filePath, "file-path", "", "Configuration file location")
	flag.Parse()
}

func main() {
	var err error

	err = checkRequiredFlags(configurationType, image, imageTag, filePath)

	if err != nil {
		errorOutput := fmt.Sprintf("Error: %v", err)
		color.Red(errorOutput)
		fmt.Println(strings.Repeat("-", len(errorOutput)))

		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	deployment, updatedDeployment, err := services.UpdateDeploymentConfiguration(file, image, imageTag)

	if err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	currentImage := deployment.Spec.Template.Spec.Containers[0].Image
	newImage := updatedDeployment.Spec.Template.Spec.Containers[0].Image

	color.Blue("current image: %s", currentImage)
	color.Green("new image: %s\n", newImage)
}
