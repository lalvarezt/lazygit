package filetree

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/stretchr/testify/assert"
)

func TestFlatTreesPreserveHierarchicalLeafOrder(t *testing.T) {
	paths := make([]string, 0, 68)
	paths = append(paths, "Z/top", "a/one", "A/two", "root", "z/three", "a/nested/four", ".hidden", "a0/file")
	for i := range 30 {
		paths = append(paths, fmt.Sprintf("Folder%02d/leaf", i), fmt.Sprintf("folder%02d/leaf", i))
	}
	for _, sortOrder := range []string{"mixed", "filesFirst", "foldersFirst"} {
		for _, caseSensitive := range []bool{false, true} {
			for _, showRoot := range []bool{false, true} {
				for _, count := range []int{0, 1, len(paths)} {
					t.Run(fmt.Sprintf("order=%s/case=%t/root=%t/count=%d", sortOrder, caseSensitive, showRoot, count), func(t *testing.T) {
						files := make([]*models.File, count)
						commits := make([]*models.CommitFile, count)
						for i, path := range paths[:count] {
							files[i] = &models.File{Path: path, Tracked: i%2 == 0, HasMergeConflicts: i%3 == 0}
							commits[i] = &models.CommitFile{Path: path}
						}
						fileOrder := NewNodeSortOrder(sortOrder, caseSensitive)
						wantFiles := BuildTreeFromFiles(files, showRoot, fileOrder).GetLeaves()
						statusGroup := func(file *models.File) int {
							group := 0
							if !file.HasMergeConflicts {
								group += 2
							}
							if !file.Tracked {
								group++
							}
							return group
						}
						slices.SortStableFunc(wantFiles, func(a, b *Node[models.File]) int {
							return cmp.Compare(statusGroup(a.File), statusGroup(b.File))
						})
						assert.Equal(t, wantFiles, BuildFlatTreeFromFiles(files, showRoot, fileOrder).Children)
						commitOrder := NewNodeSortOrder(sortOrder, caseSensitive)
						wantCommits := BuildTreeFromCommitFiles(commits, showRoot, commitOrder).GetLeaves()
						assert.Equal(t, wantCommits, BuildFlatTreeFromCommitFiles(commits, showRoot, commitOrder).Children)
					})
				}
			}
		}
	}
}

func BenchmarkBuildFlatFileTrees(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		for _, deep := range []bool{false, true} {
			for _, commitFiles := range []bool{false, true} {
				b.Run(fmt.Sprintf("files=%d/deep=%t/commits=%t", count, deep, commitFiles), func(b *testing.B) {
					files := make([]*models.File, count)
					commits := make([]*models.CommitFile, count)
					for i := range files {
						path := fmt.Sprintf("dir/file-%05d", count-i)
						if deep {
							path = fmt.Sprintf("root/dir-%05d/sub/file-%05d", i, count-i)
						}
						files[i] = &models.File{Path: path, Tracked: i%2 == 0}
						commits[i] = &models.CommitFile{Path: path}
					}
					fileOrder := NewNodeSortOrder("mixed", false)
					commitOrder := NewNodeSortOrder("mixed", false)
					b.ReportAllocs()
					b.ResetTimer()
					for b.Loop() {
						if commitFiles {
							BuildFlatTreeFromCommitFiles(commits, true, commitOrder)
						} else {
							BuildFlatTreeFromFiles(files, true, fileOrder)
						}
					}
				})
			}
		}
	}
}
