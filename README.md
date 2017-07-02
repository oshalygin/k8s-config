# k8s-config

This is a simple and straightforward CLI that allows you to update your Kubernetes configuration files dynamically during a build.

### Installation

TBD

### Usage

```bash
# Call the utility at the specified path with the appropriate command line arguments.
./k8s-config --type=deployment --file-path=../test-files/deployment.yaml  --image='oshalygin/foobar:1.0.0'
```

### Command Line Arguments

**type**: (required) The Kubernetes type that you want to update.  Depending on the type, the utility will parse it different, this value is required.

**image**: The new image name
 
**version**: If the image is excluded and just the version is passed in, you can easily just bump the version of the container image(assuming there's a single image per pod)

### License

[MIT](LICENSE)
