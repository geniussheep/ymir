package config

type Application struct {
	AppId    int32  `yaml:"appId" mapstructure:"appId"`
	AppKey   string `yaml:"appKey" mapstructure:"appKey"`
	AppName  string `yaml:"appName" mapstructure:"appName"`
	HttpPort int64  `yaml:"httpPort" mapstructure:"httpPort"`
	Mode     string `yaml:"mode" mapstructure:"mode"`
}

var ApplicationConfig = new(Application)
