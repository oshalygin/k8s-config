package models

// Deployment provides declarative updates for Pods and ReplicaSets (the next-generation ReplicationController).
type Deployment struct {
	APIVersion string
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
		RevisionHistoryLimit int
		Strategy             struct {
			RollingUpdate struct {
				MaxUnavailable int
			}
			Type string
		}
		Template struct {
			Metadata struct {
				Labels struct {
					App string
				}
			}
			Spec struct {
				Containers []struct {
					Name         string
					Image        string
					VolumeMounts []struct {
						MountPath string
						Name      string
					}
					Env []struct {
						Name      string
						ValueFrom struct {
							SecretKeyRef struct {
								Name string
								Key  string
							}
						}
					}
					ImagePullPolicy string
					Ports           []struct {
						ContainerPort int
					}
					ReadinessProbe struct {
						HTTPGet struct {
							Path string
							Port int
						}
						InitialDelaySeconds int
						TimeoutSeconds      int
						SuccessThreshold    int
						FailureThreshold    int
					}
				}
				Volumes []struct {
					Name     string
					EmptyDir struct {
					}
				}
			}
		}
	}
}
