package config

import (
	"fmt"
	"testing"
)

type Extend struct {
	Jenkins map[string]string `yaml:"jenkins"`
}

var ExtendConfigTest Extend

func TestDefault(t *testing.T) {

	ExtendConfig = &ExtendConfigTest
	Default()

	println(fmt.Sprintf("appId: %d, appName:%s", ApplicationConfig.AppId, ApplicationConfig.AppName))
	for k, v := range ExtendConfigTest.Jenkins {
		println(fmt.Sprintf("%s: %s", k, v))
	}

	for k, v := range DatabaseConfig {
		println(fmt.Sprintf("%s: %v", k, v))
	}

	for k, v := range RedisConfig {
		println(fmt.Sprintf("%s: %v", k, v))
	}
}
