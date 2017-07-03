package services

import "testing"

const configurationFile string = `
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: fe-web
  namespace: default
  labels:
    version: v1
    date: 20170311T222958Z
spec:
  replicas: 2
  revisionHistoryLimit: 2
  strategy:
    rollingUpdate:
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: fe-web
    spec:
      containers:
        - name: fe-web
          image: 'us.gcr.io/mdjs-io/merchant-dashboard:2.47.4'
          volumeMounts:
            - mountPath: /cache
              name: temp-volume
          env:
            - name: MONGO_IP
              valueFrom:
                secretKeyRef:
                  name: env-variables-secret
                  key: mongo_ip
            - name: MONGO_PORT
              valueFrom:
                secretKeyRef:
                  name: env-variables-secret
                  key: mongo_port
            - name: GCLOUD_PROJECT
              valueFrom:
                secretKeyRef:
                  name: env-variables-secret
                  key: gcloud_project
            - name: IMAGE_STORAGE_BUCKET
              valueFrom:
                secretKeyRef:
                  name: env-variables-secret
                  key: image_storage_bucket
            - name: DOMAIN_ENDPOINT
              valueFrom:
                secretKeyRef:
                  name: env-variables-secret
                  key: domain_endpoint
          imagePullPolicy: Always
          ports:
            - containerPort: 8080
          readinessProbe:
            httpGet:
              path: /healthz
              port: 8080
            initialDelaySeconds: 30
            timeoutSeconds: 5
            successThreshold: 1
            failureThreshold: 10
      volumes:
        - name: temp-volume
          emptyDir: {}
`

var file = []byte(configurationFile)

func Test_UpdateDeploymentConfiguration_ShouldReturnAnErrorIfTheFileCouldNotBeParsed(t *testing.T) {
	file := []byte{}
	image := ""
	imageTag := "1.4.4"
	_, err := UpdateDeploymentConfiguration(file, image, imageTag)

	if err != nil {
		t.Errorf("should throw an error if the file could not be parsed")
	}
}

func Test_parseConfigurationFile_ShouldParseTheConfigurationFileToAYAMLFile(t *testing.T) {

	_, err := parseConfigurationFile(file)

	if err != nil {
		t.Errorf("should successfully parse the configuration file")
	}
}

func Test_parseConfigurationFile_ShouldSetTheConfigurationTypeToDeployment(t *testing.T) {

	expected := "Deployment"
	deployment, _ := parseConfigurationFile(file)

	actual := deployment.Kind

	if expected != actual {
		t.Errorf("expected %s\nactual %s", expected, actual)
	}
}

func Test_parseConfigurationFile_ShouldSetTheContainerImageNameAccordingly(t *testing.T) {

	expected := "us.gcr.io/mdjs-io/merchant-dashboard:2.47.4"
	deployment, _ := parseConfigurationFile(file)

	actual := deployment.Spec.Template.Spec.Containers[0].Image

	if expected != actual {
		t.Errorf("\nexpected: %s\nactual: %s", expected, actual)
	}
}

func Test_updateImageTag_ShouldUpdateTheImageTagToThePassedInTagValue(t *testing.T) {

	imageTag := "1.1.1"

	expected := "us.gcr.io/mdjs-io/merchant-dashboard:1.1.1"
	deployment, _ := parseConfigurationFile(file)
	updatedDeployment := updateImageTag(deployment, imageTag)

	actual := updatedDeployment.Spec.Template.Spec.Containers[0].Image

	if expected != actual {
		t.Errorf("\nexpected: %s\nactual: %s", expected, actual)
	}
}
