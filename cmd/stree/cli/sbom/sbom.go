package sbom

import (
	"fmt"

	"github.com/viveksahu26/stree/cmd/stree/cli/options"
)

func SbomCmd(_ options.SbomOptions, args []string) error {
	fileName := args[0]
	fmt.Println("fileName: ", fileName)

	return nil
}
