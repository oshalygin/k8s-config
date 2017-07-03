package models

type Deployment struct {
	ApiVersion string
	Kind       string
	MetaData
	Spec
}

type MetaData struct {
	Name      string
	Namespace string
	Labels
}

type Labels struct {
	Version string
	Date    string
	App     string
}

type Spec struct {
	Replicas             int
	RevisionHistoryLimit int
	Strategy
	Template
}

type Strategy struct {
	RollingUpdate
	Type string
}

type RollingUpdate struct {
	MaxUnavailable int
}

type Template struct {
	MetaData
	TemplateSpec
}

type TemplateSpec struct {
	Containers []Container
	Volumes    []PodVolumes
}

type Container struct {
	Name         string
	Image        string
	VolumeMounts []ContainerVolumeMounts
	Env          []EnvironmentVariable
}

type ContainerVolumeMounts struct {
	MountPath string
	Name      string
}

type EnvironmentVariable struct {
	Name      string
	ValueFrom []EnvironmentValueFrom
}

type EnvironmentValueFrom struct {
	SecretKeyRef
}

type SecretKeyRef struct {
	Name string
	Key  string
}

type PodVolumes struct {
	Name     string
	EmptyDir string
}
