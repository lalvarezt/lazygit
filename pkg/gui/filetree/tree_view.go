package filetree

type treeView[T any] struct {
	tree           *Node[T]
	collapsedPaths *CollapsedPaths

	indexedTree      *Node[T]
	indexedCollapsed *CollapsedPaths
	indexedRevision  uint64
	nodes            []*Node[T]
	depths           []int
}

func newTreeView[T any]() treeView[T] {
	return treeView[T]{collapsedPaths: NewCollapsedPaths()}
}

func (self *treeView[T]) getNode(index int) *Node[T] {
	nodes := self.getNodes()
	if index < 0 || index >= len(nodes) {
		return nil
	}
	return nodes[index]
}

func (self *treeView[T]) getNodes() []*Node[T] {
	if self.tree == nil {
		return nil
	}
	if self.indexedTree == self.tree && self.indexedCollapsed == self.collapsedPaths && self.indexedRevision == self.collapsedPaths.revision {
		return self.nodes
	}

	clear(self.nodes)
	self.nodes = self.nodes[:0]
	self.depths = self.depths[:0]
	var visit func(*Node[T], int)
	visit = func(node *Node[T], depth int) {
		if depth >= 0 {
			self.nodes = append(self.nodes, node)
			self.depths = append(self.depths, depth)
		}
		if len(node.Children) > 0 && !self.collapsedPaths.IsCollapsed(node.path) {
			for _, child := range node.Children {
				visit(child, depth+1)
			}
		}
	}
	visit(self.tree, -1)
	self.indexedTree = self.tree
	self.indexedCollapsed = self.collapsedPaths
	self.indexedRevision = self.collapsedPaths.revision
	return self.nodes
}

func (self *treeView[T]) length() int {
	if self.tree == nil {
		return -1
	}
	return len(self.getNodes())
}

func (self *treeView[T]) ExpandToPath(path string) {
	self.collapsedPaths.ExpandToPath(path)
}

func (self *treeView[T]) GetIndexForPath(path string) (int, bool) {
	if self.tree.path == path {
		return -1, true
	}
	nodes := self.getNodes()
	for index, node := range nodes {
		if node.path == path {
			return index, true
		}
	}
	return len(nodes) - 1, false
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
	if index < 0 || index >= len(self.getNodes()) {
		return -1
	}
	return self.depths[index]
}
