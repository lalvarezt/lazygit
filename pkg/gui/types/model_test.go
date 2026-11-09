package types_test

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestModelBranchMarkerUpdates(t *testing.T) {
	pool := &utils.StringPool{}
	commits := []*models.Commit{
		models.NewCommit(pool, models.NewCommitOpts{Hash: "tip", Parents: []string{"side", "base"}}),
		models.NewCommit(pool, models.NewCommitOpts{Hash: "side", Parents: []string{"base"}}),
		models.NewCommit(pool, models.NewCommitOpts{Hash: "base", Status: models.StatusMerged}),
	}
	for _, branchesFirst := range []bool{false, true} {
		t.Run(fmt.Sprintf("branchesFirst=%t", branchesFirst), func(t *testing.T) {
			model := &types.Model{}
			assert.Empty(t, model.CommitsMergedIntoLocalMainBranch().ToSlice())
			assert.Empty(t, model.SubCommitsMergedIntoLocalMainBranch().ToSlice())
			model.SetMainBranchNames([]string{"main"})
			branches := []*models.Branch{{Name: "main", CommitHash: "tip"}}
			if branchesFirst {
				model.SetBranches(branches)
			}
			model.SetCommits(commits)
			model.SetSubCommits(commits[1:])
			if !branchesFirst {
				model.SetBranches(branches)
			}
			assert.ElementsMatch(t, []string{"tip", "side"}, model.CommitsMergedIntoLocalMainBranch().ToSlice())
			assert.Empty(t, model.SubCommitsMergedIntoLocalMainBranch().ToSlice())

			model.SetBranches([]*models.Branch{{Name: "main", CommitHash: "side"}})
			assert.ElementsMatch(t, []string{"side"}, model.CommitsMergedIntoLocalMainBranch().ToSlice())
			assert.ElementsMatch(t, []string{"side"}, model.SubCommitsMergedIntoLocalMainBranch().ToSlice())

			model.SetSubCommits(commits[2:])
			assert.Empty(t, model.SubCommitsMergedIntoLocalMainBranch().ToSlice())
			assert.ElementsMatch(t, []string{"side"}, model.CommitsMergedIntoLocalMainBranch().ToSlice())

			model.SetMainBranchNames([]string{"develop"})
			assert.Empty(t, model.CommitsMergedIntoLocalMainBranch().ToSlice())
			model.SetMainBranchNames([]string{"main"})
			assert.ElementsMatch(t, []string{"side"}, model.CommitsMergedIntoLocalMainBranch().ToSlice())
			model.SetCommits(commits[2:])
			assert.Empty(t, model.CommitsMergedIntoLocalMainBranch().ToSlice())
		})
	}
}

func TestModelBranchMarkersUseUpdatedCommitData(t *testing.T) {
	pool := &utils.StringPool{}
	commit := models.NewCommit(pool, models.NewCommitOpts{Hash: "tip", Parents: []string{"parent"}})
	parent := models.NewCommit(pool, models.NewCommitOpts{Hash: "parent"})
	model := &types.Model{}
	model.SetMainBranchNames([]string{"main"})
	model.SetBranches([]*models.Branch{{Name: "main", CommitHash: "tip"}})
	commits := []*models.Commit{commit, parent}
	model.SetCommits(commits)
	assert.ElementsMatch(t, []string{"tip", "parent"}, model.CommitsMergedIntoLocalMainBranch().ToSlice())
	commit.Status = models.StatusMerged
	model.SetCommits(commits)
	assert.Empty(t, model.CommitsMergedIntoLocalMainBranch().ToSlice())
	commits[0] = models.NewCommit(pool, models.NewCommitOpts{Hash: "tip"})
	model.SetCommits(commits)
	assert.ElementsMatch(t, []string{"tip"}, model.CommitsMergedIntoLocalMainBranch().ToSlice())
}

func TestModelBranchMarkerConfigurationIsRepoLocal(t *testing.T) {
	pool := &utils.StringPool{}
	commits := []*models.Commit{models.NewCommit(pool, models.NewCommitOpts{Hash: "tip"})}
	branches := []*models.Branch{{Name: "main", CommitHash: "tip"}}
	first, second := &types.Model{}, &types.Model{}
	names := []string{"main"}
	first.SetMainBranchNames(names)
	second.SetMainBranchNames([]string{"develop"})
	names[0] = "develop"
	for _, model := range []*types.Model{first, second} {
		model.SetCommits(commits)
		model.SetBranches(branches)
	}
	assert.ElementsMatch(t, []string{"tip"}, first.CommitsMergedIntoLocalMainBranch().ToSlice())
	assert.Empty(t, second.CommitsMergedIntoLocalMainBranch().ToSlice())
}

func BenchmarkModelBranchMarkerUpdate(b *testing.B) {
	for _, count := range []int{1000, 10000, 100000} {
		b.Run(fmt.Sprintf("commits=%d", count), func(b *testing.B) {
			pool := &utils.StringPool{}
			commits := make([]*models.Commit, count)
			for i := range commits {
				commits[i] = models.NewCommit(pool, models.NewCommitOpts{Hash: fmt.Sprintf("commit-%d", i)})
			}
			model := &types.Model{}
			model.SetMainBranchNames([]string{"main"})
			model.SetBranches([]*models.Branch{{Name: "main", CommitHash: commits[count-1].Hash()}})
			b.ReportAllocs()
			for b.Loop() {
				model.SetCommits(commits)
			}
		})
	}
}
