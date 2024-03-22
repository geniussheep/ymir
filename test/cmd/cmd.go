package cmd

import (
	"github.com/geniussheep/ymir/cli"
	"github.com/geniussheep/ymir/sdk/cmd"
	"github.com/geniussheep/ymir/sdk/common"
	"github.com/geniussheep/ymir/sdk/middleware"
	"github.com/geniussheep/ymir/test/api/router"
)

func Execute() {
	p := &cli.Program{
		Program:        "TestCmd",
		Desc:           "test cmd",
		AppRoutersScan: []func(){router.InitRouter},
		ConfigFilePath: common.DEFAULT_CONFIG_FILE_PATH,
		ExtendConfig:   nil,
		InitFuncArray:  []func(){middleware.AppendDefault},
	}
	rootCmd := cmd.New(p)
	rootCmd.Execute()
}
