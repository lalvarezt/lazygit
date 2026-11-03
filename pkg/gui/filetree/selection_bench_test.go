package filetree

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/stretchr/testify/assert"
)

func BenchmarkRestoreFileSelection(b *testing.B) {
	for _, count := range []int{100, 1000, 5000} {
		for _, removed := range []bool{false, true} {
			b.Run(fmt.Sprintf("files=%d/removed=%t", count, removed), func(b *testing.B) {
				previous := make([]*FileNode, count+1)
				current := make([]*FileNode, count+1)
				for i := range count {
					oldPath := fmt.Sprintf("deleted/file-%05d", i)
					newPath := fmt.Sprintf("other/file-%05d", i)
					previous[i] = NewFileNode(&Node[models.File]{path: oldPath, File: &models.File{Path: oldPath}})
					current[i] = NewFileNode(&Node[models.File]{path: newPath, File: &models.File{Path: newPath}})
				}
				survivor := NewFileNode(&Node[models.File]{path: "survivor", File: &models.File{Path: "survivor"}})
				previous[count], current[count] = survivor, survivor
				if !removed {
					previous[0] = current[0]
				}
				view := &FileTreeViewModel{}
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					view.findNewSelectedIdx(previous, current)
				}
			})
		}
	}
}

func BenchmarkRestoreAfterSingleRemoval(b *testing.B) {
	for _, count := range []int{100, 1000, 5000} {
		b.Run(fmt.Sprintf("files=%d", count), func(b *testing.B) {
			previous := make([]*FileNode, count+1)
			current := make([]*FileNode, count)
			previous[0] = NewFileNode(&Node[models.File]{path: "deleted", File: &models.File{Path: "deleted"}})
			for i := range current {
				path := fmt.Sprintf("other/file-%05d", i)
				current[i] = NewFileNode(&Node[models.File]{path: path, File: &models.File{Path: path}})
				previous[i+1] = current[i]
			}
			view := &FileTreeViewModel{}
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				view.findNewSelectedIdx(previous, current)
			}
		})
	}
}

func BenchmarkRefreshAfterDirectoryRemoval(b *testing.B) {
	for _, count := range []int{100, 1000, 5000} {
		b.Run(fmt.Sprintf("files=%d", count), func(b *testing.B) {
			before := make([]*models.File, 0, 2*count)
			after := make([]*models.File, count)
			for i := range count {
				before = append(before, &models.File{Path: fmt.Sprintf("deleted/file-%05d", i)})
				after[i] = &models.File{Path: fmt.Sprintf("other/file-%05d", i)}
			}
			before = append(before, after...)
			cmn := common.NewDummyCommon()
			var view *FileTreeViewModel
			b.ReportAllocs()
			for b.Loop() {
				b.StopTimer()
				files := before
				view = NewFileTreeViewModel(func() []*models.File { return files }, cmn, false)
				view.SetTree()
				view.SetSelection(0)
				files = after
				b.StartTimer()
				view.SetTree()
			}
			assert.Equal(b, "other/file-00000", view.GetSelectedPath())
		})
	}
}
