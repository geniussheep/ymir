package config

type K8S struct {
	OutOfCluster   bool   `yaml:"outOfCluster"  mapstructure:"outOfCluster"`
	KubeConfigPath string `yaml:"kubeConfigPath"  mapstructure:"kubeConfigPath"`
	WebBaseURL     string `yaml:"webBaseURL"  mapstructure:"webBaseURL"`
}

var K8SConfig = make(map[string]*K8S)
