package cli

import (
	"github.com/spf13/cobra"
	"github.com/viveksahu26/tree/cmd/tree/cli/options"
)

var ro = &options.RootOptions{}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "treee",
		Short: "A tool to convert json files into directory structure.",
		// DisableAutoGenTag: true,
		SilenceUsage: true, // Don't show usage on errors
	}

	ro.AddFlags(cmd)

	// Add sub-commands
	cmd.AddCommand(Json())

	return cmd
}
