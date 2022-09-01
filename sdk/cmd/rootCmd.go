package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.benlai.work/go/ymir/cli"
	"gitlab.benlai.work/go/ymir/sdk/cmd/apiServer"
	"gitlab.benlai.work/go/ymir/sdk/cmd/config"
	"gitlab.benlai.work/go/ymir/sdk/cmd/version"
	"gitlab.benlai.work/go/ymir/sdk/pkg"
	"os"
)

type rootCmd struct {
	*cli.Program
	flagSets map[string]cli.FlagSet
	cmd      *cobra.Command
}

// New 新建默认Cli程序主体
func New[T any](p *cli.Program) cli.Cmd {
	c := rootCmd{
		Program:  p,
		flagSets: make(map[string]cli.FlagSet, 0),
		cmd: &cobra.Command{
			Use:          p.Program,
			SilenceUsage: true,
			Long:         p.Desc,
			Args: func(cmd *cobra.Command, args []string) error {
				return nil
			},
			PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
			Run: func(cmd *cobra.Command, args []string) {
				usageStr := `欢迎使用 ` + pkg.Green(fmt.Sprintf("%s %s", p.Program, p.Version)+` 可以使用 `+pkg.Red(`-h`)+` 查看命令`)
				fmt.Printf("%s\n", usageStr)
			},
		},
	}
	c.Setup()
	return &c
}

func (c *rootCmd) Cmd() *cobra.Command {
	return c.cmd
}

func (c *rootCmd) FlagsCmd() map[string]cli.FlagSet {
	return c.flagSets
}

func (c *rootCmd) Setup() {
	c.flagSets["server"] = apiServer.New(c.Program)
	c.flagSets["version"] = version.New(c.Program)
	c.flagSets["config"] = config.New(c.Program)
	for _, f := range c.flagSets {
		c.cmd.AddCommand(f.Cmd())
	}
}

func (c *rootCmd) Execute() {
	if err := c.cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
