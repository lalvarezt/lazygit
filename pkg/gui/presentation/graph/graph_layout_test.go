package graph

import (
	"fmt"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetNextPipesWithManyParents(t *testing.T) {
	hashPool := &utils.StringPool{}
	parents := make([]string, 100)
	for i := range parents {
		parents[i] = fmt.Sprintf("parent-%d", i)
	}
	commit := models.NewCommit(hashPool, models.NewCommitOpts{Hash: "merge", Parents: parents})
	previous := []Pipe{{fromHash: &StartCommitHash, toHash: commit.HashPtr(), kind: STARTS, style: &style.FgDefault}}
	getStyle := func(*models.Commit) *style.TextStyle { return &style.FgDefault }

	want := make([]Pipe, 1, len(commit.ParentPtrs())+1)
	want[0] = Pipe{fromHash: &StartCommitHash, toHash: commit.HashPtr(), kind: TERMINATES, style: &style.FgDefault}
	for i, parent := range commit.ParentPtrs() {
		want = append(want, Pipe{toPos: int16(i), fromHash: commit.HashPtr(), toHash: parent, kind: STARTS, style: &style.FgDefault})
	}
	assert.Equal(t, want, getNextPipes(previous, commit, getStyle))
}

func TestGetNextPipesWithSparseColumns(t *testing.T) {
	hashPool := &utils.StringPool{}
	commit := models.NewCommit(hashPool, models.NewCommitOpts{Hash: "commit", Parents: []string{"parent"}})
	previous := []Pipe{
		{fromHash: hashPool.Add("left"), toHash: hashPool.Add("later"), kind: STARTS, style: &style.FgDefault},
		{fromPos: 1024, toPos: 1024, fromHash: hashPool.Add("right"), toHash: commit.HashPtr(), kind: STARTS, style: &style.FgDefault},
	}
	getStyle := func(*models.Commit) *style.TextStyle { return &style.FgDefault }
	assert.Equal(t, []Pipe{
		{fromHash: hashPool.Add("left"), toHash: hashPool.Add("later"), kind: CONTINUES, style: &style.FgDefault},
		{fromPos: 1024, toPos: 1024, fromHash: hashPool.Add("right"), toHash: commit.HashPtr(), kind: TERMINATES, style: &style.FgDefault},
		{fromPos: 1024, toPos: 1024, fromHash: commit.HashPtr(), toHash: commit.ParentPtrs()[0], kind: STARTS, style: &style.FgDefault},
	}, getNextPipes(previous, commit, getStyle))
}
