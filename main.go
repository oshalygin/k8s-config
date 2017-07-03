package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	configurationType := flag.String("type", "deployment", "Kubernetes configuration type, eg: deployment, rc, secret")
	image := flag.String("image", "", "Docker image name")
	imageTag := flag.String("tag", "", "Docker image tag")
	filePath := flag.String("file-path", "", "Configuration file location")

	flag.Parse()

	err := checkRequiredFlags(*configurationType, *image, *imageTag, *filePath)
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
