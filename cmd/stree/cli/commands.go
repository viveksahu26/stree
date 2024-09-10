package cli

import (
	"github.com/spf13/cobra"
	"github.com/viveksahu26/stree/cmd/stree/cli/options"
)

var ro = &options.RootOptions{}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stree",
		Short: "A tool to convert json files into directory structure.",
		// DisableAutoGenTag: true,
		SilenceUsage: true, // Don't show usage on errors
	}

	ro.AddFlags(cmd)

	// Add sub-commands
	cmd.AddCommand(Json())
	cmd.AddCommand(Sbom())

	return cmd
}
