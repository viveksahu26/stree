package options

import "github.com/spf13/cobra"

type JsonOptions struct {
	Out  string
	Name bool
}

// var _ Interface = (*JsonOptions)(nil)

// AddFlags implements Interface
func (o *JsonOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.Out, "out", "o", "", "Output file to write the tree structure")
	cmd.Flags().BoolVarP(&o.Name, "name", "n", false, "short name of component")
}
