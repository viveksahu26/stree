package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viveksahu26/tree/cmd/tree/cli/json"
	"github.com/viveksahu26/tree/cmd/tree/cli/options"
)

func Json() *cobra.Command {
	o := &options.JsonOptions{}

	jsonCmd := &cobra.Command{
		Use:          "json",
		Short:        "Cconvert json into tree structure",
		SilenceUsage: true,
		Example:      ` tree json --file=<json file>`,
		Args:         cobra.MinimumNArgs(1),
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
