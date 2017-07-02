package main

import (
	"flag"
	"fmt"
	"os"
)

func checkRequiredFlags(configurationType *string, image *string, imageTag *string, filePath *string) {
	if *image == "" && *imageTag == "" {
		fmt.Printf(`Either the "image" or "tag" flag MUST be passed in.  Both cannot be empty`)
		flag.PrintDefaults()
		os.Exit(1)
	}
}
