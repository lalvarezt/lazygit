package filetree

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
)

func BenchmarkBuildCommitFileTree(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("files=%d", count), func(b *testing.B) {
			files := make([]*models.CommitFile, count)
			for i := range files {
				files[i] = &models.CommitFile{Path: fmt.Sprintf("directory/file-%05d", count-i)}
			}
			cmp := NewNodeSortOrder("mixed", false)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				BuildTreeFromCommitFiles(files, true, cmp)
			}
		})
	}
}
