package options

import "github.com/spf13/cobra"

type SbomOptions struct {
	Out          string
	Dependencies string
	Primary      bool
}

// var _ Interface = (*JsonOptions)(nil)

// AddFlags implements Interface
func (o *SbomOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.Out, "out", "o", "", "Output file to write the tree structure")
	cmd.Flags().StringVar(&o.Dependencies, "dependencies", "direct", "Specify 'direct' for direct dependencies or 'indirect' for all dependencies")
	cmd.Flags().BoolVar(&o.Primary, "primary", false, "Show dependencies only for the primary component")
}
