package filetree

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/stretchr/testify/assert"
)

func TestTreeViewIndexFollowsVisibility(t *testing.T) {
	for _, treeMode := range []bool{false, true} {
		for _, showRoot := range []bool{false, true} {
			t.Run(fmt.Sprintf("tree=%t/root=%t", treeMode, showRoot), func(t *testing.T) {
				view := newTreeView[models.File]()
				files := []*models.File{{Path: "a/b/one"}, {Path: "a/b/two"}, {Path: "c/three"}, {Path: "c/four"}}
				build := func() *Node[models.File] {
					cmp := NewNodeSortOrder("mixed", true)
					if treeMode {
						return BuildTreeFromFiles(files, showRoot, cmp)
					}
					return BuildFlatTreeFromFiles(files, showRoot, cmp)
				}
				view.tree = build()
				check := func() {
					t.Helper()
					want := view.tree.Flatten(view.collapsedPaths)[1:]
					assert.Equal(t, want, view.getNodes())
					assert.Equal(t, len(want), view.length())
					for i, node := range want {
						assert.Same(t, node, view.getNode(i))
						assert.Equal(t, view.tree.GetVisualDepthAtIndex(i+1, view.collapsedPaths), view.GetVisualDepth(i))
						index, found := view.GetIndexForPath(node.path)
						assert.True(t, found)
						assert.Equal(t, i, index)
					}
					index, found := view.GetIndexForPath("missing")
					assert.False(t, found)
					assert.Equal(t, len(want)-1, index)
					index, found = view.GetIndexForPath(view.tree.path)
					assert.True(t, found)
					assert.Equal(t, -1, index)
					assert.Nil(t, view.getNode(-1))
					assert.Nil(t, view.getNode(len(want)))
					assert.Equal(t, -1, view.GetVisualDepth(-1))
					assert.Equal(t, -1, view.GetVisualDepth(len(want)))
				}
				check()
				check()
				view.ToggleCollapsed("a/b")
				check()
				view.ExpandToPath("a/b/one")
				check()
				view.CollapsedPaths().Collapse("c")
				check()
				view.CollapsedPaths().ExpandToPath("c/three")
				check()
				view.CollapsedPaths().ToggleCollapsed("c")
				check()
				view.ExpandAll()
				check()
				view.CollapsedPaths().Collapse("a/b")
				check()
				view.CollapsedPaths().ExpandAll()
				check()
				view.collapsedPaths = NewCollapsedPaths()
				view.collapsedPaths.Collapse("c")
				check()
				files = []*models.File{{Path: "replacement"}}
				view.tree = build()
				check()
				files = nil
				view.tree = build()
				check()
			})
		}
	}
}
