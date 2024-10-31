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

# see dependencies in a directory tree or file tree structure for the provided sbom
  stree sbom samples/sbomqs-fossa-cyclonedx.json
		
# see only primary dependencies in a directory tree or file tree structure for the provided SBOM
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
