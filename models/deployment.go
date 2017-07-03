package models

// Deployment provides declarative updates for Pods and ReplicaSets (the next-generation ReplicationController).
type Deployment struct {
	APIVersion string
	Kind       string
	MetaData
	Spec
}

// MetaData field
type MetaData struct {
	Name      string
	Namespace string
	Labels
}

// Labels represents arbitrary labels that can be attached to a kubernetes type
type Labels struct {
	Version string
	Date    string
	App     string
}

// Spec represents the Deployment specification
type Spec struct {
	Replicas             int
	RevisionHistoryLimit int
	Strategy
	Template
}

// Strategy represents the way in which updates occur
type Strategy struct {
	RollingUpdate
	Type string
}

// RollingUpdate describes the rolling update configuration
type RollingUpdate struct {
	MaxUnavailable int
}

// Template represents the deployment template
type Template struct {
	MetaData
	TemplateSpec
}

// TemplateSpec provides details for all containers and volumes in a pod
type TemplateSpec struct {
	Containers []Container
	Volumes    []PodVolumes
}

// Container represents the container in the pod
type Container struct {
	Name         string
	Image        string
	VolumeMounts []ContainerVolumeMounts
	Env          []EnvironmentVariable
}

// ContainerVolumeMounts represents the mounted volumes on a container
type ContainerVolumeMounts struct {
	MountPath string
	Name      string
}

// EnvironmentVariable describes the environment variable configuration
type EnvironmentVariable struct {
	Name      string
	ValueFrom []EnvironmentValueFrom
}

// EnvironmentValueFrom represents the scheme in how the environment variable is retrieved
type EnvironmentValueFrom struct {
	SecretKeyRef
}

// SecretKeyRef provides the key value pair of a secret configuration
type SecretKeyRef struct {
	Name string
	Key  string
}

// PodVolumes provides the pod volume configuration
type PodVolumes struct {
	Name     string
	EmptyDir string
}
