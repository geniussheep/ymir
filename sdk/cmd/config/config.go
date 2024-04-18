package config

import (
	"encoding/json"
	"fmt"
	"github.com/geniussheep/ymir/cli"
	"github.com/geniussheep/ymir/sdk/common"
	"github.com/geniussheep/ymir/sdk/config"
	"github.com/spf13/cobra"
)

type configFlag struct {
	cmd *cobra.Command
}

func New(p *cli.Program) cli.FlagSet {
	f := configFlag{
		cmd: &cobra.Command{
			Use:     "config",
			Short:   "get app config info",
			Example: fmt.Sprintf("%s config -c %s", p.Program, common.DefaultConfigFilePath),
			Run: func(cmd *cobra.Command, args []string) {
				run(p)
			},
		},
	}
	f.cmd.PersistentFlags().StringVarP(&p.ConfigFilePath, "config", "c", common.DefaultConfigFilePath, "start server with provided configuration file")
	return &f
}

func (f *configFlag) Cmd() *cobra.Command {
	return f.cmd
}

func run(p *cli.Program) {
	config.ExtendConfig = &p.ExtendConfig
	config.Setup(p.ConfigFilePath)
	application, errs := json.MarshalIndent(config.Instance().Application, "", "") //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println("application:", string(application))

	// todo 需要兼容
	database, errs := json.MarshalIndent(config.Instance().Databases, "", "") //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println("database:", string(database))
}
