package version

import (
	"fmt"
	"github.com/geniussheep/ymir/cli"
	"github.com/geniussheep/ymir/sdk/common"
	"github.com/spf13/cobra"
)

type versionFlag struct {
	cmd *cobra.Command
}

func New(p *cli.Program) cli.FlagSet {
	f := versionFlag{
		cmd: &cobra.Command{
			Use:     "version",
			Short:   "Get version info",
			Example: fmt.Sprintf("%s version", p.Program),
			PreRun: func(cmd *cobra.Command, args []string) {

			},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("the app:%s version is %s", p.Program, p.Version)
			},
		},
	}
	f.cmd.PersistentFlags().StringVarP(&p.ConfigFilePath, "config", "c", common.DEFAULT_CONFIG_FILE_PATH, "start server with provided configuration file")
	return &f
}

func (f *versionFlag) Cmd() *cobra.Command {
	return f.cmd
}
