package sbom

import (
	"fmt"
	"os"

	"github.com/viveksahu26/stree/cmd/stree/cli/options"
	"github.com/viveksahu26/stree/pkg/sbom"
)

func validatePath(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return nil
}

func SbomCmd(so options.SbomOptions, args []string) error {
	fileName := args[0]

	err := validatePath(fileName)
	if err != nil {
		return err
	}

	sbom.SBOM(fileName)

	fmt.Println("fileName: ", fileName)

	return nil
}
