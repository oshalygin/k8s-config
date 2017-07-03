package main

import (
	"errors"
	"strings"
)

func checkRequiredFlags(configurationType string, image string, imageTag string, filePath string) error {

	if filePath == "" {
		err := `The "file-path" flag must be set`
		return errors.New(err)
	}

	if image == "" && imageTag == "" {
		err := `Either the "image" or "tag" flag MUST be passed in.  Both cannot be empty`
		return errors.New(err)
	}

	// Currently only deployments are supported
	if configurationType != "" && configurationType != "deployment" {
		err := `Only the "deployment" type is supported at this time`
		return errors.New(err)
	}

	if image != "" {
		tag := strings.Split(image, ":")
		if len(tag) <= 1 {
			err := `No tag was supplied in the image flag`
			return errors.New(err)
		}
	}

	if image != "" && imageTag != "" {
		err := `Set either the "image" or "imageTag" flag, NOT both`
		return errors.New(err)
	}

	return nil

}
