package config

import "github.com/geniussheep/ymir/sdk/pkg"

type Http struct {
	Cors            Cors              `yaml:"cors" mapstructure:"cors"`
	ResponseHeaders map[string]string `yaml:"responseHeaders" mapstructure:"responseHeaders"`
}

type Cors struct {
	AllowCredentials string `yaml:"allowCredentials"  mapstructure:"allowCredentials"`
	AllowOrigin      string `yaml:"allowOrigin"  mapstructure:"allowOrigin"`
	AllowHeaders     string `yaml:"allowHeaders"  mapstructure:"allowHeaders"`
	AllowMethods     string `yaml:"allowMethods"  mapstructure:"allowMethods"`
	MaxAge           string `yaml:"maxAge"  mapstructure:"maxAge"`
}

var HttpConfig = new(Http)

const (
	CorsAllowOrigin      = "*"
	CorsAllowCredentials = "true"
	CorsAllowHeaders     = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With"
	CorsAllowMethods     = "GET, HEAD, POST, PUT, DELETE, TRACE, OPTIONS, PATCH"
	CorsMaxAge           = "3600"
)

func (http *Http) SetDefault() {
	if pkg.StringIsEmpty(http.Cors.AllowCredentials) {
		http.Cors.AllowCredentials = CorsAllowCredentials
	}
	if pkg.StringIsEmpty(http.Cors.AllowMethods) {
		http.Cors.AllowMethods = CorsAllowMethods
	}
	if pkg.StringIsEmpty(http.Cors.AllowHeaders) {
		http.Cors.AllowHeaders = CorsAllowHeaders
	}
	if pkg.StringIsEmpty(http.Cors.AllowOrigin) {
		http.Cors.AllowOrigin = CorsAllowOrigin
	}
	if pkg.StringIsEmpty(http.Cors.MaxAge) {
		http.Cors.MaxAge = CorsMaxAge
	}
}
