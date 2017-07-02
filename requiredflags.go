package main

import (
	"errors"
)

func CheckRequiredFlags(configurationType string, image string, imageTag string, filePath string) error {
	if image == "" && imageTag == "" {
		err := `Either the "image" or "tag" flag MUST be passed in.  Both cannot be empty`
		return errors.New(err)
	}

	return nil
}
