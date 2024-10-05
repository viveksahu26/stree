package tsbom

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Component struct {
	Ref       string   `json:"ref"`
	DependsOn []string `json:"dependsOn,omitempty"`
}

type SBOM struct {
	Components   []Component  `json:"components"`
	Dependencies []Dependency `json:"dependencies"`
}

type Dependency struct {
	Ref       string   `json:"ref"`
	DependsOn []string `json:"dependsOn,omitempty"`
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

func ParseSBOM(filename string) (SBOM, error) {
	var sbom SBOM
	data, err := os.ReadFile(filename)
	if err != nil {
		return sbom, err
	}

	err = json.Unmarshal(data, &sbom)
	return sbom, err
}

func BuildSBOMTreeNode(ref string, sbom SBOM, depth int, dependencies string, primary bool) *tview.TreeNode {
	nodeText := ref
	parts := strings.Split(nodeText, "@")
	nodeText = parts[0]
	color := colors[depth%len(colors)]

	tv := tview.NewTreeNode(nodeText).SetColor(color)
	node := tv.SetReference(ref)

	for _, dep := range sbom.Dependencies {
		if dep.Ref == ref {
			for _, childRef := range dep.DependsOn {
				childNode := BuildSBOMTreeNode(childRef, sbom, depth+1, dependencies, primary)
				node.AddChild(childNode)
			}
			break
		}
	}

	return node
}
