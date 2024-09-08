package main

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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	root, err := parseJSON(filename)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	printTree(root, "", true)
}
