package json

import (
	"fmt"

	"github.com/viveksahu26/tree/cmd/tree/cli/options"
	"github.com/viveksahu26/tree/pkg/tjson"
)

func JsonCmd(jsonOption options.JsonOptions, args []string) error {
	fileName := args[0]
	fmt.Println("fileName: ", fileName)
	root, err := tjson.ParseJSON(fileName)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}
	tjson.PrintTree(root, "", true)
	return nil
}
