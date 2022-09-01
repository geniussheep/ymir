package cli

import (
	"github.com/spf13/cobra"
)

type Cmd interface {
	Cmd() *cobra.Command
	FlagsCmd() map[string]FlagSet
	Setup()
	Execute()
}
