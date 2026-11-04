package filetree

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/stretchr/testify/assert"
)

func TestTreeRangeSelection(t *testing.T) {
	type rangeView interface {
		types.IListCursor
		SetTree()
		Len() int
		ToggleCollapsed(string)
		GetSelectedItemIds() ([]string, int, int)
	}
	for _, commitFiles := range []bool{false, true} {
		for _, treeMode := range []bool{false, true} {
			for _, showRoot := range []bool{false, true} {
				t.Run(fmt.Sprintf("commits=%t/tree=%t/root=%t", commitFiles, treeMode, showRoot), func(t *testing.T) {
					cmn := common.NewDummyCommon()
					cmn.UserConfig().Gui.ShowRootItemInFileTree = showRoot
					paths := []string{"a/b/one", "a/b/two", "c/three", "c/four", "five"}
					var view rangeView
					var allIds func() []string
					if commitFiles {
						files := make([]*models.CommitFile, 0, len(paths))
						for _, path := range paths {
							files = append(files, &models.CommitFile{Path: path})
						}
						model := NewCommitFileTreeViewModel(func() []*models.CommitFile { return files }, cmn, treeMode)
						view = model
						allIds = func() []string {
							nodes := model.GetAllItems()
							ids := make([]string, 0, len(nodes))
							for _, node := range nodes {
								ids = append(ids, node.ID())
							}
							return ids
						}
					} else {
						files := make([]*models.File, 0, len(paths))
						for _, path := range paths {
							files = append(files, &models.File{Path: path})
						}
						model := NewFileTreeViewModel(func() []*models.File { return files }, cmn, treeMode)
						view = model
						allIds = func() []string {
							nodes := model.GetAllItems()
							ids := make([]string, 0, len(nodes))
							for _, node := range nodes {
								ids = append(ids, node.ID())
							}
							return ids
						}
					}
					view.SetTree()
					for _, collapsed := range []bool{false, true} {
						if collapsed {
							view.ToggleCollapsed("a/b")
						}
						ids := allIds()
						for _, selection := range [][2]int{{0, 0}, {0, len(ids) - 1}, {len(ids) - 1, 1}, {2, 1}} {
							view.SetSelection(selection[0])
							view.ToggleStickyRange()
							view.SetSelectedLineIdx(selection[1])
							selected, start, end := view.GetSelectedItemIds()
							assert.Equal(t, min(selection[0], selection[1]), start)
							assert.Equal(t, max(selection[0], selection[1]), end)
							assert.Equal(t, ids[start:end+1], selected)
						}
					}
				})
			}
		}
	}
}
