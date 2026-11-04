package filetree

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/common"
)

func BenchmarkFileTreeRangeSelection(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("files=%d", count), func(b *testing.B) {
			files := make([]*models.File, count)
			for i := range files {
				files[i] = &models.File{Path: fmt.Sprintf("dir/file-%05d", i)}
			}
			view := NewFileTreeViewModel(func() []*models.File { return files }, common.NewDummyCommon(), true)
			view.SetTree()
			view.SetSelection(0)
			view.ToggleStickyRange()
			view.SetSelectedLineIdx(view.Len() - 1)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				view.GetSelectedItems()
			}
		})
	}
}
