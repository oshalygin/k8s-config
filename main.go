package main

import (
	"flag"
	"fmt"
	"os"

	"io/ioutil"

	"github.com/fatih/color"
	"github.com/oshalygin/k8s-config/models"
	"gopkg.in/yaml.v2"
)

func main() {
	configurationFile := models.Deployment{}

	configurationType := flag.String("type", "deployment", "Kubernetes configuration type, eg: deployment, rc, secret")
	image := flag.String("image", "", "Docker image name")
	imageTag := flag.String("tag", "", "Docker image tag")
	filePath := flag.String("file-path", "", "Configuration file location")

	flag.Parse()

	file, err := ioutil.ReadFile("./test-files/deployment.yaml")
	if err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, &configurationFile)
	if err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	err = checkRequiredFlags(*configurationType, *image, *imageTag, *filePath)
	if err != nil {
		color.Red("Error: %v", err)
		color.Black("--------------------")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("type: %s\n", *configurationType)
	fmt.Printf("image: %s\n", *image)
	fmt.Printf("tag: %s\n", *imageTag)
	fmt.Printf("file-path: %s\n", *filePath)

}
