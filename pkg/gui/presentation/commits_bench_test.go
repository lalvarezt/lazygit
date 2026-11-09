package presentation

import (
	"fmt"
	"testing"
	"time"

	"github.com/jesseduffield/generics/set"
	"github.com/jesseduffield/lazygit/pkg/commands/git_commands"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/jesseduffield/lazygit/pkg/gui/branchstack"
	"github.com/jesseduffield/lazygit/pkg/gui/presentation/graph"
	"github.com/jesseduffield/lazygit/pkg/utils"
)

func BenchmarkCommitViewport(b *testing.B) {
	for _, count := range []int{1000, 10000, 100000} {
		for _, main := range []string{"absent", "merged", "unmerged"} {
			for _, markers := range []bool{false, true} {
				b.Run(fmt.Sprintf("commits=%d/main=%s/markers=%t", count, main, markers), func(b *testing.B) {
					c := common.NewDummyCommon()
					c.UserConfig().Git.MainBranches = []string{"main"}
					pool := &utils.StringPool{}
					commits := make([]*models.Commit, count)
					for i := range commits {
						opts := models.NewCommitOpts{Hash: fmt.Sprintf("%040d", i), Name: "subject", Parents: []string{fmt.Sprintf("%040d", i+1)}}
						if i == count-1 {
							opts.Status = models.StatusMerged
						}
						commits[i] = models.NewCommit(pool, opts)
					}
					tip := "not-loaded"
					if main == "merged" {
						tip = commits[count-1].Hash()
					}
					if main == "unmerged" {
						tip = commits[count-2].Hash()
					}
					branches := []*models.Branch{{Name: "main", CommitHash: tip}, {Name: "feature", CommitHash: commits[0].Hash()}}
					if markers {
						branches = append(branches, &models.Branch{Name: "other", CommitHash: commits[count/2].Hash()})
					}
					picked := set.New[string]()
					bisect := git_commands.NewNullBisectInfo()
					now := time.Unix(1700000000, 0)
					mergedLocally := branchstack.CommitsMergedIntoLocalMainBranch(commits, branches, c.UserConfig().Git.MainBranches)
					b.ReportAllocs()
					b.ResetTimer()
					for b.Loop() {
						GetCommitListDisplayStrings(c, commits, branches, mergedLocally, "feature", false, false, picked, "", "", "", "", now, false, nil, count/2, count/2+40, false, graph.BoxDrawingSymbols, bisect)
					}
				})
			}
		}
	}
}
