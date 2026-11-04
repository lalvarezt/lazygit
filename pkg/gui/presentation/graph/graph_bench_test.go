package graph

import (
	"fmt"
	"testing"

	"github.com/gookit/color"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation/authors"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/xo/terminfo"
)

func BenchmarkRenderCommitGraphVariants(b *testing.B) {
	oldColorLevel := color.ForceSetColorLevel(terminfo.ColorLevelMillions)
	defer color.ForceSetColorLevel(oldColorLevel)

	for _, count := range []int{50, 500} {
		for _, symbols := range []struct {
			name string
			set  SymbolSet
		}{
			{name: "box", set: BoxDrawingSymbols},
			{name: "branch", set: BranchDrawingSymbols},
		} {
			for _, selected := range []bool{false, true} {
				b.Run(fmt.Sprintf("commits=%d/symbols=%s/selected=%t", count, symbols.name, selected), func(b *testing.B) {
					commits := generateCommits(&utils.StringPool{}, count)
					var selectedHash *string
					if selected {
						selectedHash = commits[count/2].HashPtr()
					}
					getStyle := func(commit *models.Commit) *style.TextStyle {
						return authors.AuthorStyle(commit.AuthorName)
					}
					b.ReportAllocs()
					b.ResetTimer()
					for b.Loop() {
						RenderCommitGraph(commits, selectedHash, getStyle, symbols.set)
					}
				})
			}
		}
	}
}
