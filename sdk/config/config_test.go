package config

import (
	"fmt"
	"testing"
)

func TestDefault(t *testing.T) {
	c := Default()
	println(fmt.Sprintf("appId: %d, appName:%s", c.Application.AppId, c.Application.AppName))
}
