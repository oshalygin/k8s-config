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

func main() {
	var err error

	configurationType := flag.String("type", "deployment", "Kubernetes configuration type, eg: deployment, rc, secret")
	image := flag.String("image", "", "Docker image name")
	imageTag := flag.String("tag", "", "Docker image tag")
	filePath := flag.String("file-path", "", "Configuration file location")

	flag.Parse()
	err = checkRequiredFlags(*configurationType, *image, *imageTag, *filePath)

	if err != nil {
		errorOutput := fmt.Sprintf("Error: %v", err)
		color.Red(errorOutput)
		fmt.Println(strings.Repeat("-", len(errorOutput)))

		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}

	services.UpdateDeploymentConfiguration(file, *image, *imageTag)

	fmt.Printf("type: %s\n", *configurationType)
	fmt.Printf("image: %s\n", *image)
	fmt.Printf("tag: %s\n", *imageTag)
	fmt.Printf("file-path: %s\n", *filePath)

}
