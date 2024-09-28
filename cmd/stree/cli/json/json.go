package json

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
	"github.com/viveksahu26/stree/cmd/stree/cli/options"
	"github.com/viveksahu26/stree/pkg/tjson"
)

func JsonCmd(jo options.JsonOptions, args []string) error {
	fileName := args[0]
	fmt.Println("fileName: ", fileName)
	root, err := tjson.ParseJSON(fileName)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}

	rootNode := tjson.BuildTreeNode(root)
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
