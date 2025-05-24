package json

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
	"github.com/viveksahu26/stree/cmd/stree/cli/options"
	"github.com/viveksahu26/stree/pkg/tjson"
)

func JsonCmd(jo options.JsonOptions, args []string) error {
	fileName := args[0]
	doesNodenameToBeShort := jo.Name

	root, err := tjson.ParseJSON(fileName)
	if err != nil {
		return fmt.Errorf("parsing JSON: %w", err)
	}
	if root.Name == "" && len(root.Children) == 0 {
		return fmt.Errorf("empty or invalid JSON structure in %s", fileName)
	}

	// Print text-based tree if no output file or UI-only mode
	if jo.Out == "" {
		fmt.Print(tjson.PrintTree(root, "", true))
	}

	// Save to output file if specified
	if jo.Out != "" {
		output := tjson.PrintTree(root, "", true)
		if err := os.WriteFile(jo.Out, []byte(output), 0o644); err != nil {
			return fmt.Errorf("writing output to %s: %w", jo.Out, err)
		}
		return nil
	}

	// Build interactive UI
	rootNode := tjson.BuildTreeNode(root, 0, doesNodenameToBeShort)

	// Explicitly collapse all nodes (except root)
	rootNode.CollapseAll()
	rootNode.SetExpanded(true) // Keep root expanded to show direct dependencies

	treeView := tview.
		NewTreeView().
		SetRoot(rootNode).
		SetCurrentNode(rootNode)

	// Handle node selection (click or Enter) to toggle expansion
	treeView.SetSelectedFunc(func(node *tview.TreeNode) {
		if node.IsExpanded() {
			node.Collapse()
		} else {
			node.Expand()
		}
	})

	// Add title bar
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText(fmt.Sprintf("Dependency Tree: %s (Ctrl+Q to quit)", fileName)), 1, 1, false).
		AddItem(treeView, 0, 1, true)

	app := tview.NewApplication()
	if err := app.SetRoot(flex, true).Run(); err != nil {
		return fmt.Errorf("running application: %w", err)
	}
	return nil
}
