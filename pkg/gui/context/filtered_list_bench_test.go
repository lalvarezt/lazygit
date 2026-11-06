package context

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
)

func BenchmarkFilteredListNavigation(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("items=%d", count), func(b *testing.B) {
			items := make([]*models.Branch, count)
			for i := range items {
				items[i] = &models.Branch{Name: fmt.Sprintf("branch-%d", i)}
			}
			list := NewFilteredListViewModel(func() []*models.Branch { return items }, func(branch *models.Branch) []string { return []string{branch.Name} })
			list.SetFilter("branch", false)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				list.SetSelectedLineIdx(count / 2)
				list.MoveSelectedLine(1)
				list.GetSelected()
				list.Len()
			}
		})
	}
}
