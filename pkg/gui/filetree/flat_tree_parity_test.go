package filetree

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/common"
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

func TestFlatTreeOrderAcrossPathShapes(t *testing.T) {
	shapes := make([][]string, 0, 12)
	shapes = append(shapes,
		[]string{"same", "same"},
		[]string{"file", "file/child"},
		[]string{"file/child", "file"},
		[]string{"Kelvin/one", "İstanbul/two", "界/leaf", "UPPER/file", "alpha/z", "alpha/Key/last"},
	)
	for seed := range int64(8) {
		rng := rand.New(rand.NewSource(seed))
		paths := make([]string, 100)
		for i := range paths {
			paths[i] = fmt.Sprintf("dir-%d/sub-%d/file-%d", rng.Intn(8), rng.Intn(8), i)
		}
		rng.Shuffle(len(paths), func(i, j int) { paths[i], paths[j] = paths[j], paths[i] })
		shapes = append(shapes, paths)
	}
	for shape, paths := range shapes {
		for _, sortOrder := range []string{"mixed", "filesFirst", "foldersFirst"} {
			for _, caseSensitive := range []bool{false, true} {
				for _, showRoot := range []bool{false, true} {
					t.Run(fmt.Sprintf("shape=%d/order=%s/case=%t/root=%t", shape, sortOrder, caseSensitive, showRoot), func(t *testing.T) {
						files := make([]*models.CommitFile, len(paths))
						for i, path := range paths {
							files[i] = &models.CommitFile{Path: path}
						}
						order := NewNodeSortOrder(sortOrder, caseSensitive)
						want := BuildTreeFromCommitFiles(files, showRoot, order).GetLeaves()
						got := BuildFlatTreeFromCommitFiles(files, showRoot, order).Children
						assert.Equal(t, want, got)
						for i := range got {
							assert.Same(t, want[i].File, got[i].File)
						}
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

func BenchmarkFlatTreeRefresh(b *testing.B) {
	for _, count := range []int{100, 10000} {
		for _, commitFiles := range []bool{false, true} {
			b.Run(fmt.Sprintf("files=%d/commits=%t", count, commitFiles), func(b *testing.B) {
				files := make([]*models.File, count)
				commits := make([]*models.CommitFile, count)
				for i := range files {
					path := fmt.Sprintf("dir-%05d/sub/file", i)
					files[i] = &models.File{Path: path}
					commits[i] = &models.CommitFile{Path: path}
				}
				c := common.NewDummyCommon()
				working := NewFileTreeViewModel(func() []*models.File { return files }, c, false)
				committed := NewCommitFileTreeViewModel(func() []*models.CommitFile { return commits }, c, false)
				working.SetTree()
				committed.SetTree()
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					if commitFiles {
						committed.SetTree()
					} else {
						working.SetTree()
					}
				}
			})
		}
	}
}
