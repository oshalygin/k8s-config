[![Build Status](https://travis-ci.org/oshalygin/k8s-config.svg?branch=master)](https://travis-ci.org/oshalygin/k8s-config)
[![Coverage Status](https://coveralls.io/repos/github/oshalygin/k8s-config/badge.svg?branch=master)](https://coveralls.io/github/oshalygin/k8s-config?branch=master)
[![Code Climate](https://codeclimate.com/repos/59598bef371afb02870005c1/badges/d8e88772201d137ea8b7/gpa.svg)](https://codeclimate.com/repos/59598bef371afb02870005c1/feed)
[![Issue Count](https://codeclimate.com/repos/59598bef371afb02870005c1/badges/d8e88772201d137ea8b7/issue_count.svg)](https://codeclimate.com/repos/593e287a150338028600480b/feed)
# Introduction

This is a simple and straightforward CLI that allows you to update your Kubernetes configuration files dynamically during a build.

# Installation

```bash
go get -u github.com/oshalygin/k8s-config
```

# Usage

```
Usage of k8s-config:
  -type string
        The Kubernetes configuration type(default: "deployment")
  -file-path string
        The location of the configuration file
  -image string
        The full image name, 'oshalygin/foobar:1.0.0'
  -tag string
        If just the tag is provided without the image,
        then only the tag will be updated on the existing
        image name.
```

### Example

```bash
# Call the utility at the specified path with the 
# appropriate command line arguments.

k8s-config --type=deployment \
             --file-path=../test-files/deployment.yaml \  
             --image='oshalygin/my-website-image:1.0.3'
             
# Although it is recommended to be explicit, here is
# the minimum amount of flags you can pass and still
# use the utility effectively

k8s-config --tag=1.03
           --file-path=../test-files/deployment.yaml 
```

### Command Line Arguments

**type**: The Kubernetes type that you want to update.  Depending on the type, the utility will parse it different, this value is required.

**image**: The new image name

**file-path**: The location of the configuration file
 
**tag**: If the image is excluded and just the version is passed in, you can easily just bump the version of the container image(assuming there's a single image per pod)

# Limitations
_Note_: This list will change accordingly over the development of this utility.
* Only `yaml` configuration files(Not `JSON`)
* Only `deployment` types.
* Only single container deployment configurations.

# License

[MIT](LICENSE)
