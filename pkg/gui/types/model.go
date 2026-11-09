package types

import (
	"slices"

	"github.com/jesseduffield/generics/set"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/branchstack"
)

func (self *Model) SetCommits(commits []*models.Commit) {
	self.Commits = commits
	self.commitsMergedLocally = *branchstack.CommitsMergedIntoLocalMainBranch(commits, self.Branches, self.mainBranchNames)
}

func (self *Model) SetSubCommits(commits []*models.Commit) {
	self.SubCommits = commits
	self.subCommitsMergedLocally = *branchstack.CommitsMergedIntoLocalMainBranch(commits, self.Branches, self.mainBranchNames)
}

func (self *Model) SetBranches(branches []*models.Branch) {
	self.Branches = branches
	self.SetCommits(self.Commits)
	self.SetSubCommits(self.SubCommits)
}

func (self *Model) SetMainBranchNames(names []string) {
	if slices.Equal(self.mainBranchNames, names) {
		return
	}
	self.mainBranchNames = slices.Clone(names)
	self.SetCommits(self.Commits)
	self.SetSubCommits(self.SubCommits)
}

func (self *Model) CommitsMergedIntoLocalMainBranch() *set.Set[string] {
	return &self.commitsMergedLocally
}

func (self *Model) SubCommitsMergedIntoLocalMainBranch() *set.Set[string] {
	return &self.subCommitsMergedLocally
}
