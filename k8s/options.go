package k8s

type Option func(*options)

type options struct {
	OutOfCluster   bool   `yaml:"outOfCluster"`
	KubeConfigPath string `yaml:"kubeConfigPath"`
	WebBaseURL     string `yaml:"webBaseURL"`
}

func setDefault() options {
	return options{
		OutOfCluster:   false,
		KubeConfigPath: "",
		WebBaseURL:     "",
	}
}

func SetOutOfCluster(b bool) Option {
	return func(o *options) {
		o.OutOfCluster = b
	}
}

func SetKubeConfigPath(path string) Option {
	return func(o *options) {
		o.KubeConfigPath = path
	}
}

func SetWebBaseURL(url string) Option {
	return func(o *options) {
		o.WebBaseURL = url
	}
}
