package cli

import "github.com/spf13/cobra"

type FlagSet interface {
	Cmd() *cobra.Command
}
