package filetree

type nodeSelection interface {
	Len() int
	GetSelectionRange() (int, int)
}

func selectedNodes[T, N any](cursor nodeSelection, root *Node[T], collapsedPaths *CollapsedPaths, wrap func(*Node[T]) N) ([]N, int, int) {
	if cursor.Len() == 0 {
		return nil, 0, 0
	}

	startIdx, endIdx := cursor.GetSelectionRange()
	nodes := make([]N, 0, endIdx-startIdx+1)
	index := -1 // the root is not a displayed row
	var visit func(*Node[T])
	visit = func(node *Node[T]) {
		if index >= startIdx {
			nodes = append(nodes, wrap(node))
		}
		index++
		if !collapsedPaths.IsCollapsed(node.path) {
			for _, child := range node.Children {
				if index > endIdx {
					break
				}
				visit(child)
			}
		}
	}
	visit(root)
	return nodes, startIdx, endIdx
}
