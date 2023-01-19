package config

import (
	"fmt"
	"testing"
)

type ExtendTest struct {
	Jenkins      string `yaml:"jenkins" defaultValue:"https://jenkins.benlai.cloud"`
	JenkinsAgent string `yaml:"jenkinsAgent" defaultValue:"jenkins-agent"`
}

var ExtendConfigTest ExtendTest

func TestDefault(t *testing.T) {

	ExtendConfig = &ExtendConfigTest
	Default()

	println(fmt.Sprintf("appId: %d, appName:%s", ApplicationConfig.AppId, ApplicationConfig.AppName))
	println(fmt.Sprintf("jenkins: %s, jenkinsAgent:%s", ExtendConfigTest.Jenkins, ExtendConfigTest.JenkinsAgent))

	for k, v := range DatabaseConfig {
		println(fmt.Sprintf("%s: %v", k, v))
	}

	for k, v := range RedisConfig {
		println(fmt.Sprintf("%s: %v", k, v))
	}
}
