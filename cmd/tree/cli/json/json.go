package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/viveksahu26/tree/cmd/tree/cli/options"
)

// Package represents a package with its name and children
type Package struct {
	Name     string    `json:"name"`
	Children []Package `json:"children"`
}

func JsonCmd(jsonOption options.JsonOptions, args []string) error {
	fileName := args[0]
	fmt.Println("fileName: ", fileName)
	root, err := parseJSON(fileName)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}
	printTree(root, "", true)
	return nil
}

// parseJSON parses the JSON input into a Package struct
func parseJSON(filename string) (Package, error) {
	var root Package
	data, err := os.ReadFile(filename)
	if err != nil {
		return root, err
	}
	err = json.Unmarshal(data, &root)
	return root, err
}

// printTree recursively prints the tree structure with pointers
func printTree(pkg Package, prefix string, isLast bool) {
	if isLast {
		fmt.Println(prefix + "└── " + pkg.Name)
		prefix += "    "
	} else {
		fmt.Println(prefix + "├── " + pkg.Name)
		prefix += "│   "
	}

	for i, child := range pkg.Children {
		printTree(child, prefix, i == len(pkg.Children)-1)
	}
}
