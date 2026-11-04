package filetree

type treeView[T any] struct {
	tree           *Node[T]
	collapsedPaths *CollapsedPaths
}

func newTreeView[T any]() treeView[T] {
	return treeView[T]{collapsedPaths: NewCollapsedPaths()}
}

func (self *treeView[T]) getNode(index int) *Node[T] {
	return self.tree.GetNodeAtIndex(index+1, self.collapsedPaths)
}

func (self *treeView[T]) getNodes() []*Node[T] {
	return self.tree.Flatten(self.collapsedPaths)[1:]
}

func (self *treeView[T]) length() int {
	return self.tree.Size(self.collapsedPaths) - 1
}

func (self *treeView[T]) ExpandToPath(path string) {
	self.collapsedPaths.ExpandToPath(path)
}

func (self *treeView[T]) GetIndexForPath(path string) (int, bool) {
	index, found := self.tree.GetIndexForPath(path, self.collapsedPaths)
	return index - 1, found
}

func (self *treeView[T]) IsCollapsed(path string) bool {
	return self.collapsedPaths.IsCollapsed(path)
}

func (self *treeView[T]) ToggleCollapsed(path string) {
	self.collapsedPaths.ToggleCollapsed(path)
}

func (self *treeView[T]) ExpandAll() {
	self.collapsedPaths.ExpandAll()
}

func (self *treeView[T]) CollapsedPaths() *CollapsedPaths {
	return self.collapsedPaths
}

func (self *treeView[T]) GetVisualDepth(index int) int {
	return self.tree.GetVisualDepthAtIndex(index+1, self.collapsedPaths)
}
