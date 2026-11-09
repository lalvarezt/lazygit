package graph

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/gookit/color"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/xo/terminfo"
)

func graphViewportFixture(branches bool) ([]*models.Commit, [][]Pipe) {
	pool := &utils.StringPool{}
	commits := generateCommits(pool, 500)
	if !branches {
		for i := range commits {
			commits[i] = models.NewCommit(pool, models.NewCommitOpts{
				Hash: fmt.Sprintf("commit-%d", i), Parents: []string{fmt.Sprintf("commit-%d", i+1)},
			})
		}
	}
	return commits, GetPipeSets(commits, func(*models.Commit) *style.TextStyle { return &style.FgDefault })
}

func TestRenderAuxMatchesSequentialRows(t *testing.T) {
	oldColorLevel := color.ForceSetColorLevel(terminfo.ColorLevelMillions)
	t.Cleanup(func() { color.ForceSetColorLevel(oldColorLevel) })
	oldProcs := runtime.GOMAXPROCS(0)
	t.Cleanup(func() { runtime.GOMAXPROCS(oldProcs) })
	for _, branches := range []bool{false, true} {
		commits, pipes := graphViewportFixture(branches)
		for _, count := range []int{0, 1, 7, 40, 120} {
			for _, symbols := range []SymbolSet{BoxDrawingSymbols, BranchDrawingSymbols} {
				for _, selected := range []*string{nil, commits[100+count/2].HashPtr()} {
					want := make([]string, count)
					for i := range want {
						var previous *models.Commit
						if i > 0 {
							previous = commits[100+i-1]
						}
						want[i] = renderPipeSet(pipes[100+i], selected, previous, symbols)
					}
					for _, procs := range []int{1, 4, 16} {
						runtime.GOMAXPROCS(procs)
						assert.Equal(t, want, RenderAux(pipes[100:100+count], commits[100:100+count], selected, symbols))
					}
				}
			}
		}
	}
}

func BenchmarkGraphViewport(b *testing.B) {
	oldColorLevel := color.ForceSetColorLevel(terminfo.ColorLevelMillions)
	defer color.ForceSetColorLevel(oldColorLevel)
	for _, branches := range []bool{false, true} {
		commits, pipes := graphViewportFixture(branches)
		for _, count := range []int{1, 8, 40, 120} {
			b.Run(fmt.Sprintf("branches=%t/rows=%d", branches, count), func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					RenderAux(pipes[100:100+count], commits[100:100+count], nil, BoxDrawingSymbols)
				}
			})
		}
	}
}
