package sbom

import (
	"fmt"
	"log"
	"os"

	"github.com/rivo/tview"
	"github.com/viveksahu26/stree/cmd/stree/cli/options"
	"github.com/viveksahu26/stree/pkg/tsbom"
)

func validatePath(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return nil
}

func SbomCmd(so options.SbomOptions, args []string) error {
	fileName := args[0]
	fmt.Println("fileName: ", fileName)
	dependencies := so.Dependencies
	primary := so.Primary
	doesNodenameToBeShort := so.Name

	sbom, err := tsbom.ParseSBOM(fileName)
	if err != nil {
		fmt.Println("Error parsing SBOM:", err)
		return err
	}

	var rootNode *tview.TreeNode
	if primary {
		rootRef := sbom.Dependencies[0].Ref
		rootNode = tsbom.BuildSBOMTreeNode(rootRef, sbom, 0, dependencies, primary, doesNodenameToBeShort)
	} else {
		// rootNode = tview.NewTreeNode("SBOM").SetExpanded(false)
		rootNode = tview.NewTreeNode("SBOM")
		for _, dep := range sbom.Dependencies {
			childNode := tsbom.BuildSBOMTreeNode(dep.Ref, sbom, 0, dependencies, primary, doesNodenameToBeShort)
			rootNode.AddChild(childNode)
		}
	}
	treeView := tview.NewTreeView().SetRoot(rootNode).SetCurrentNode(rootNode)

	app := tview.NewApplication()
	treeView.SetSelectedFunc(func(node *tview.TreeNode) {
		if node.IsExpanded() {
			node.Collapse()
		} else {
			node.Expand()
		}
	})

	if err := app.SetRoot(treeView, true).Run(); err != nil {
		log.Fatalf("Error running application: %v\n", err)
	}
	return nil
}
