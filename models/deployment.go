package models

// Deployment provides declarative updates for Pods and ReplicaSets (the next-generation ReplicationController).
type Deployment struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    struct {
			Version string
			Date    string
		}
	}
	Spec struct {
		Replicas             int
		RevisionHistoryLimit int `yaml:"revisionHistoryLimit"`
		Strategy             struct {
			RollingUpdate `yaml:"rollingUpdate"`
			Type          string
		}
		Template struct {
			Metadata struct {
				Labels struct {
					App string
				}
			}
			Spec struct {
				Containers []Container
				Volumes    []struct {
					Name     string
					EmptyDir `yaml:"emptyDir"`
				}
			}
		}
	}
}

// Container represents the running container in a pod
type Container struct {
	Name         string
	Image        string
	VolumeMounts `yaml:"volumeMounts"`
	Env          []struct {
		Name      string
		ValueFrom `yaml:"valueFrom,omitempty"`
		Value     string `yaml:"value,omitempty"`
	}
	ImagePullPolicy string `yaml:"imagePullPolicy"`
	Ports           []struct {
		ContainerPort int `yaml:"containerPort"`
	}
	ReadinessProbe `yaml:"readinessProbe"`
}

// RollingUpdate represents the rolling update strategy
type RollingUpdate struct {
	MaxUnavailable int `yaml:"maxUnavailable"`
}

// HTTPGet configures the readiness probes
type HTTPGet struct {
	Path string
	Port int
}

// ReadinessProbe describes the configuration around readiness checks
type ReadinessProbe struct {
	HTTPGet             `yaml:"httpGet"`
	InitialDelaySeconds int `yaml:"initialDelaySeconds"`
	TimeoutSeconds      int `yaml:"timeoutSeconds"`
	SuccessThreshold    int `yaml:"successThreshold"`
	FailureThreshold    int `yaml:"failureThreshold"`
}

// SecretKeyRef provides a description to how secrets are pulled
type SecretKeyRef struct {
	Name string
	Key  string
}

// FieldRef is the configuration object of how values are pulled from the cluster
type FieldRef struct {
	FieldPath string `yaml:"fieldPath,omitempty"`
}

// ValueFrom is the configuration object of how values are pulled
type ValueFrom struct {
	SecretKeyRef `yaml:"secretKeyRef,omitempty"`
	FieldRef     `yaml:"fieldRef,omitempty"`
}

// VolumeMounts provides a configuration for additional pod volumes
type VolumeMounts []struct {
	MountPath string `yaml:"mountPath"`
	Name      string
}

// EmptyDir provides a descriptor for volume configuration
type EmptyDir struct {
}
