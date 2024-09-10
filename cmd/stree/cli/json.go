package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viveksahu26/stree/cmd/stree/cli/json"
	"github.com/viveksahu26/stree/cmd/stree/cli/options"
)

func Json() *cobra.Command {
	o := &options.JsonOptions{}

	jsonCmd := &cobra.Command{
		Use:          "json",
		Short:        "Cconvert json into tree structure",
		SilenceUsage: true,
		Example: ` tree json [--out <path>] <json file>

	# create a tree directory like structure for the provided json
	tree json example-sca.json

	# create a tree directory like structure for the provided json and output into provided file
	tree json -o <output-new-file.json>  example-sca.json

		
		`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			if err := json.JsonCmd(*o, args); err != nil {
				return fmt.Errorf("coverting json into tree structure for file %v: %w", args, err)
			}
			return nil
		},
	}
	o.AddFlags(jsonCmd)
	return jsonCmd
}
