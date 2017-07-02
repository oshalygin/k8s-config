package main

import (
	"testing"
)

func Test_CheckRequiredFlags_ShouldExit1IfImageAndTagAreNotPassed(t *testing.T) {
	imageTag := ""
	configurationType := ""
	image := ""
	filePath := ""

	err := CheckRequiredFlags(configurationType, image, filePath, imageTag)
	if err == nil {
		t.Errorf("should throw an error if both image and filepath are empty")
	}

}
