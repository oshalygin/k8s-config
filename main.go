package main

import (
	"flag"
	"fmt"
)

func main() {
	configurationType := flag.String("type", "deployment", "Kubernetes configuration type, eg: deployment, rc, secret")
	image := flag.String("image", "", "Docker image name")
	imageTag := flag.String("tag", "", "Docker image tag")
	filePath := flag.String("file-path", "", "Configuration file location")

	flag.Parse()
	CheckRequiredFlags(*configurationType, *image, *imageTag, *filePath)

	fmt.Printf("type: %s\n", *configurationType)
	fmt.Printf("image: %s\n", *image)
	fmt.Printf("tag: %s\n", *imageTag)
	fmt.Printf("file-path: %s\n", *filePath)

}
