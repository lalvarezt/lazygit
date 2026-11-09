package filetree

type nodeSelection interface {
	Len() int
	GetSelectionRange() (int, int)
}

func selectedNodes[N any](cursor nodeSelection, getRange func(int, int) []N) ([]N, int, int) {
	if cursor.Len() == 0 {
		return nil, 0, 0
	}

	startIdx, endIdx := cursor.GetSelectionRange()
	return getRange(startIdx, endIdx+1), startIdx, endIdx
}
