package sbom

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/viveksahu26/stree/pkg/tjson"
)

func SBOM(path string) error {
	var err error

	// Run the opensca-cli command
	cmd := exec.Command("opensca-cli", "-path", path, "-out", "xyz-sca.json")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	if err != nil {
		fmt.Printf("failed to run opensca-cli: %s\n", err)
		return err
	}

	root, err := tjson.ParseJSON("xyz-sca.json")
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}
	// (root, "", true)
	treeOutput := tjson.PrintTree(root, "", true)

	fmt.Print(treeOutput)

	// Print the output for debugging purposes
	fmt.Println("Command output:", out.String())

	// Delete the file "xyz-sca.json"
	err = os.Remove("xyz-sca.json")
	if err != nil {
		fmt.Printf("failed to delete file: %s\n", err)
		return err
	}
	return nil
}

func processFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open %s\n", path)
		return nil, err
	}

	return f, nil
}
