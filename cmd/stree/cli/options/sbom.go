package options

import "github.com/spf13/cobra"

type SbomOptions struct {
	Out string
}

// var _ Interface = (*JsonOptions)(nil)

// AddFlags implements Interface
func (o *SbomOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.Out, "out", "o", "", "Output file to write the tree structure")
}