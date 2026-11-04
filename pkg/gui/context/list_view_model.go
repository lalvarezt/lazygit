package context

import (
	"github.com/jesseduffield/lazygit/pkg/gui/context/traits"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/samber/lo"
)

type HasID interface {
	ID() string
}

type ListViewModel[T HasID] struct {
	*traits.ListCursor
	getModel func() []T
	getRange func(int, int) []T
}

func NewListViewModel[T HasID](getModel func() []T) *ListViewModel[T] {
	return newListViewModel(getModel, func() int { return len(getModel()) })
}

func newListViewModel[T HasID](getModel func() []T, getLength func() int) *ListViewModel[T] {
	self := &ListViewModel[T]{
		getModel: getModel,
		getRange: func(start, end int) []T { return getModel()[start:end] },
	}

	self.ListCursor = traits.NewListCursor(getLength)

	return self
}

func (self *ListViewModel[T]) GetSelected() T {
	if self.Len() == 0 {
		return Zero[T]()
	}

	index := self.GetSelectedLineIdx()
	return self.getRange(index, index+1)[0]
}

func (self *ListViewModel[T]) GetSelectedItemId() string {
	if self.Len() == 0 {
		return ""
	}

	return self.GetSelected().ID()
}

func (self *ListViewModel[T]) GetSelectedItems() ([]T, int, int) {
	if self.Len() == 0 {
		return nil, -1, -1
	}

	startIdx, endIdx := self.GetSelectionRange()

	return self.getRange(startIdx, endIdx+1), startIdx, endIdx
}

func (self *ListViewModel[T]) GetSelectedItemIds() ([]string, int, int) {
	selectedItems, startIdx, endIdx := self.GetSelectedItems()

	ids := lo.Map(selectedItems, func(item T, _ int) string {
		return item.ID()
	})

	return ids, startIdx, endIdx
}

func (self *ListViewModel[T]) GetItems() []T {
	return self.getModel()
}

func Zero[T any]() T {
	return *new(T)
}

func (self *ListViewModel[T]) GetItem(index int) types.HasUrn {
	item := self.getRange(index, index+1)[0]
	return any(item).(types.HasUrn)
}
