package config

type Application struct {
	AppId    int32  `yaml:"appid"`
	AppKey   string `yaml:"appkey"`
	AppName  string `yaml:"appname"`
	HttpPort int64  `yaml:"httpport"`
}
