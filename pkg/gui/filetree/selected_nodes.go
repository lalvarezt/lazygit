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
	nodes := []N{}
	for i := startIdx; i <= endIdx; i++ {
		nodes = append(nodes, wrap(root.GetNodeAtIndex(i+1, collapsedPaths)))
	}
	return nodes, startIdx, endIdx
}
