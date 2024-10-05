package tjson

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Package represents a package with its name and children
type Package struct {
	Name     string    `json:"name"`
	Version  string    `json:"version,omitempty"`
	Language string    `json:"language"`
	ID       string    `json:"id"`
	Paths    []string  `json:"paths"`
	Children []Package `json:"children"`
}

// ParseJSON parses the JSON input into a Package struct
func ParseJSON(filename string) (Package, error) {
	var root Package
	data, err := os.ReadFile(filename)
	if err != nil {
		return root, err
	}

	// convert JSON data into a Go struct.
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

var colors = []tcell.Color{
	tcell.ColorYellow,
	tcell.ColorRed,
	tcell.ColorGreen,
	tcell.ColorPurple,
	tcell.ColorBlue,
	tcell.ColorDarkCyan,
	tcell.ColorDarkOrange,
	tcell.ColorPink,
}

// The function creates a tree node for a given package.
// It sets the text and color of the node based on the package name and depth.
// It recursively creates and adds child nodes for each child package.
// The result is a tree structure that represents the hierarchy of packages, with each node correctly linked to its children and visually distinguished by color based on depth.
func BuildTreeNode(pkg Package, depth int, doesNodenameToBeShort bool) *tview.TreeNode {
	nodeText := pkg.Name
	if doesNodenameToBeShort {
		nodeText = ShortName(nodeText)
	}
	color := colors[depth%len(colors)]

	tv := tview.NewTreeNode(nodeText).SetColor(color)
	node := tv.SetReference(pkg)

	for _, child := range pkg.Children {
		childNode := BuildTreeNode(child, depth+1, doesNodenameToBeShort)
		node.AddChild(childNode)
	}

	return node
}

func ShortName(longName string) string {
	parts := strings.Split(longName, "/")
	var shortname string
	if len(parts) > 0 {
		shortname = parts[0]

		if len(parts) > 1 {
			shortname = parts[0] + "/" + parts[1]
		}

		if len(parts) > 2 && parts[1] != "" && parts[2] != "" {
			shortname = parts[1] + "/" + parts[2]
		}
	}
	return shortname
}
