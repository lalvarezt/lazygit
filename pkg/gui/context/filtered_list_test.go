package context

import (
	"testing"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/context/traits"
	"github.com/stretchr/testify/assert"
)

func TestFilteredListSelectionRanges(t *testing.T) {
	items := []*models.Branch{{Name: "alpha"}, {Name: "beta"}, {Name: "alpine"}, {Name: "alphabet"}}
	list := NewFilteredListViewModel(func() []*models.Branch { return items }, func(branch *models.Branch) []string { return []string{branch.Name} })
	list.SetFilter("alp", false)
	list.SetSelectionRangeAndMode(2, 1, traits.RangeSelectModeSticky)
	selected, start, end := list.GetSelectedItems()
	assert.Equal(t, []*models.Branch{items[2], items[3]}, selected)
	assert.Equal(t, 1, start)
	assert.Equal(t, 2, end)
	assert.Same(t, items[3], list.GetSelected())
	assert.Same(t, items[2], list.GetItem(1))
	assert.Equal(t, []*models.Branch{items[0], items[2], items[3]}, list.GetItems())

	list.ClearFilter()
	list.SetSelectionRangeAndMode(1, 2, traits.RangeSelectModeSticky)
	selected, start, end = list.GetSelectedItems()
	assert.Equal(t, items[1:3], selected)
	assert.Equal(t, 1, start)
	assert.Equal(t, 2, end)
	assert.Same(t, items[1], list.GetSelected())
	assert.Same(t, items[2], list.GetItem(2))

	list.SetFilter("missing", false)
	selected, start, end = list.GetSelectedItems()
	assert.Nil(t, selected)
	assert.Equal(t, -1, start)
	assert.Equal(t, -1, end)
	assert.Empty(t, list.GetSelectedItemId())
}

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
