package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	c := New("./conf/config.yaml")
	println(fmt.Sprint("app :%s", c.Get("application")))
}
