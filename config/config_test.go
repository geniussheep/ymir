package config

import (
	"fmt"
	"gitlab.benlai.work/go/ymir/sdk/common"
	"testing"
)

func TestNew(t *testing.T) {

	c, err := New(common.DEFAULT_CONFIG_FILE_PATH)
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Println(c.Get("databases.monitor"))
}
