package options

import (
	"time"

	"github.com/spf13/cobra"
)

// RootOptions define flags and options for the root tree cli.
type RootOptions struct {
	OutputFile string
	Verbose    bool
	Timeout    time.Duration
}

// DefaultTimeout specifies the default timeout for commands.
const DefaultTimeout = 3 * time.Minute

func (o *RootOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&o.OutputFile, "output-file", "",
		"log output to a file")
	_ = cmd.Flags().SetAnnotation("output-file", cobra.BashCompFilenameExt, []string{})

	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "d", false,
		"log debug output")

	cmd.PersistentFlags().DurationVarP(&o.Timeout, "timeout", "t", DefaultTimeout,
		"timeout for commands")
}
