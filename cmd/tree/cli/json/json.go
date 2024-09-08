package json

import (
	"fmt"
	"os"

	"github.com/viveksahu26/tree/cmd/tree/cli/options"
	"github.com/viveksahu26/tree/pkg/tjson"
)

func JsonCmd(jo options.JsonOptions, args []string) error {
	fileName := args[0]
	fmt.Println("fileName: ", fileName)
	root, err := tjson.ParseJSON(fileName)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}
	treeOutput := tjson.PrintTree(root, "", true)
	if jo.Out != "" {
		err = os.WriteFile(jo.Out, []byte(treeOutput), 0o644)
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	} else {
		fmt.Print(treeOutput)
	}
	return nil
}
