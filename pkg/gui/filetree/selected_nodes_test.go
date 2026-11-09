package filetree

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestSelectedTreeRangesFollowVisibility(t *testing.T) {
	for _, treeMode := range []bool{false, true} {
		for _, showRoot := range []bool{false, true} {
			t.Run(fmt.Sprintf("tree=%t/root=%t", treeMode, showRoot), func(t *testing.T) {
				c := common.NewDummyCommon()
				c.UserConfig().Gui.ShowRootItemInFileTree = showRoot
				files := []*models.File{{Path: "a/b/one"}, {Path: "a/b/two"}, {Path: "c/three"}, {Path: "c/four"}}
				commits := make([]*models.CommitFile, len(files))
				for i, file := range files {
					commits[i] = &models.CommitFile{Path: file.Path}
				}
				working := NewFileTreeViewModel(func() []*models.File { return files }, c, treeMode)
				committed := NewCommitFileTreeViewModel(func() []*models.CommitFile { return commits }, c, treeMode)
				working.SetTree()
				committed.SetTree()
				check := func() {
					t.Helper()
					wantWorking := working.GetRoot().Raw().Flatten(working.CollapsedPaths())[1:]
					wantCommitted := committed.GetRoot().Raw().Flatten(committed.CollapsedPaths())[1:]
					for start := range wantWorking {
						for end := start; end < len(wantWorking); end++ {
							working.SetSelection(start)
							working.ToggleStickyRange()
							working.SetSelectedLineIdx(end)
							selected, first, last := working.GetSelectedItems()
							assert.Equal(t, start, first)
							assert.Equal(t, end, last)
							for i, node := range selected {
								assert.Same(t, wantWorking[start+i], node.Raw())
							}
							assert.Len(t, selected, end-start+1)
							committed.SetSelection(start)
							committed.ToggleStickyRange()
							committed.SetSelectedLineIdx(end)
							selectedCommits, first, last := committed.GetSelectedItems()
							assert.Equal(t, start, first)
							assert.Equal(t, end, last)
							for i, node := range selectedCommits {
								assert.Same(t, wantCommitted[start+i], node.Raw())
							}
							assert.Len(t, selectedCommits, end-start+1)
						}
					}
				}
				check()
				working.ToggleCollapsed(InternalTreePathForFilePath("a/b", showRoot))
				committed.ToggleCollapsed(InternalTreePathForFilePath("a/b", showRoot))
				check()
				files, commits = files[:2], commits[:2]
				working.SetTree()
				committed.SetTree()
				check()
				files, commits = nil, nil
				working.SetTree()
				committed.SetTree()
				selected, _, _ := working.GetSelectedItems()
				assert.Nil(t, selected)
				selectedCommits, _, _ := committed.GetSelectedItems()
				assert.Nil(t, selectedCommits)
			})
		}
	}
}

func BenchmarkSmallTreeRangeSelection(b *testing.B) {
	for _, count := range []int{100, 10000} {
		for _, cold := range []bool{false, true} {
			for _, commitFiles := range []bool{false, true} {
				b.Run(fmt.Sprintf("files=%d/cold=%t/commits=%t", count, cold, commitFiles), func(b *testing.B) {
					files := make([]*models.File, count)
					commits := make([]*models.CommitFile, count)
					for i := range files {
						path := fmt.Sprintf("dir/file-%05d", i)
						files[i] = &models.File{Path: path}
						commits[i] = &models.CommitFile{Path: path}
					}
					c := common.NewDummyCommon()
					working := NewFileTreeViewModel(func() []*models.File { return files }, c, true)
					committed := NewCommitFileTreeViewModel(func() []*models.CommitFile { return commits }, c, true)
					working.SetTree()
					committed.SetTree()
					working.SetSelection(working.Len() - 2)
					committed.SetSelection(committed.Len() - 2)
					b.ReportAllocs()
					b.ResetTimer()
					for b.Loop() {
						if commitFiles {
							if cold {
								committed.ICommitFileTree.(*CommitFileTree).indexedTree = nil
							}
							committed.GetSelectedItems()
						} else {
							if cold {
								working.IFileTree.(*FileTree).indexedTree = nil
							}
							working.GetSelectedItems()
						}
					}
				})
			}
		}
	}
}
