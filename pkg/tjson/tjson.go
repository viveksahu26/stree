package tjson

import (
	"encoding/json"
	"fmt"
	"os"
)

// Package represents a package with its name and children
type Package struct {
	Name     string    `json:"name"`
	Children []Package `json:"children"`
}

// ParseJSON parses the JSON input into a Package struct
func ParseJSON(filename string) (Package, error) {
	var root Package
	data, err := os.ReadFile(filename)
	if err != nil {
		return root, err
	}
	err = json.Unmarshal(data, &root)
	return root, err
}

// PrintTree recursively prints the tree structure with pointers
func PrintTree(pkg Package, prefix string, isLast bool) string {
	var result string
	if isLast {
		result += fmt.Sprintf("%s└── %s\n", prefix, pkg.Name)
		prefix += "    "
	} else {
		result += fmt.Sprintf("%s├── %s\n", prefix, pkg.Name)
		prefix += "│   "
	}

	for i, child := range pkg.Children {
		result += PrintTree(child, prefix, i == len(pkg.Children)-1)
	}
	return result
}
