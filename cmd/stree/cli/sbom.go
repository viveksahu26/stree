package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viveksahu26/stree/cmd/stree/cli/options"
	"github.com/viveksahu26/stree/cmd/stree/cli/sbom"
)

func Sbom() *cobra.Command {
	o := &options.SbomOptions{}

	sbomCmd := &cobra.Command{
		Use:          "sbom",
		Short:        "represent SBOM Dependencies as a directory tree or file tree",
		SilenceUsage: true,
		Example: `stree sbom [--dependencies direct|indirect] <sbom file>

# create a directory tree or file tree structure for the provided sbom
  stree sbom sbom.spdx.json
		
# create a directory tree or file tree structure for the provided sbom and output into provided file
  stree sbom --dependencies indirect example-sbom.json
		
# create a tree directory like structure for the provided SBOM with only primary dependencies
  stree sbom --primary example-sbom.json

		
		`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			if err := sbom.SbomCmd(*o, args); err != nil {
				return fmt.Errorf("coverting sbom into directory tree or file tree structure for file %v: %w", args, err)
			}
			return nil
		},
	}
	o.AddFlags(sbomCmd)
	return sbomCmd
}
