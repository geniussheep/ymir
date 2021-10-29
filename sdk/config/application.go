package config

type Application struct {
	AppId    int32  `yaml:"appId"`
	AppKey   string `yaml:"appKey"`
	AppName  string `yaml:"appName"`
	HttpPort int64  `yaml:"httpPort"`
}
