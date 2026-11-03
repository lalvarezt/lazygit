package context

import (
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/stretchr/testify/assert"
)

func TestFilteredListLengthFollowsModelAndFilter(t *testing.T) {
	items := []*models.Branch{{Name: "alpha"}, {Name: "beta"}, {Name: "alpine"}}
	list := NewFilteredListViewModel(func() []*models.Branch { return items }, func(branch *models.Branch) []string { return []string{branch.Name} })
	assert.Equal(t, 3, list.Len())

	list.SetFilter("alp", false)
	assert.Equal(t, 2, list.Len())
	list.SetSelection(1)
	assert.Equal(t, "alpine", list.GetSelected().Name)

	items = append(items, &models.Branch{Name: "alphabet"})
	list.ReApplyFilter(false)
	assert.Equal(t, 3, list.Len())
	list.SetSelection(2)
	assert.Equal(t, "alphabet", list.GetSelected().Name)

	list.ClearFilter()
	assert.Equal(t, 4, list.Len())
	assert.Equal(t, "alphabet", list.GetSelected().Name)
	items = items[:1]
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, "alpha", list.GetSelected().Name)

	list.SetFilter("missing", false)
	assert.Zero(t, list.Len())
	assert.Nil(t, list.GetSelected())
	list.ClearFilter()
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, "alpha", list.GetSelected().Name)
}
