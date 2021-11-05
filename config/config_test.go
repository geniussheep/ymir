package config

import (
	"testing"
)

func TestNew(t *testing.T) {

	c := New("./conf/config.yaml")

	c.Get("databases.monitor.")
}
