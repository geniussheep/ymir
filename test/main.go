package main

import (
	"gitlab.benlai.work/go/ymir/cli"
	"gitlab.benlai.work/go/ymir/sdk/cmd"
	"gitlab.benlai.work/go/ymir/sdk/common"
	"gitlab.benlai.work/go/ymir/sdk/middleware"
	"gitlab.benlai.work/go/ymir/test/api/router"
)

func main() {
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
