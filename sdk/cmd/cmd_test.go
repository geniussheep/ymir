package cmd

import (
	"gitlab.benlai.work/go/ymir/cli"
	"gitlab.benlai.work/go/ymir/sdk/common"
	"testing"
)

func TestRootCmd(t *testing.T) {
	p := &cli.Program{
		Program:        "TestCmd",
		Desc:           "test cmd",
		AppRoutersScan: nil,
		ConfigFilePath: common.DEFAULT_CONFIG_FILE_PATH,
		ExtendConfig:   nil,
	}
	rootCmd := New[any](p)
	rootCmd.Execute()
}
