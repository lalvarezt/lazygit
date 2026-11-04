package gocui

import (
	"strings"

	"github.com/rivo/uniseg"
)

type lineSearcher struct {
	normalize func(string) string
	graphemes []string
	width     int
}

func newLineSearcher(query string) lineSearcher {
	normalize := strings.ToLower
	if containsUpcaseChar(query) {
		normalize = func(s string) string { return s }
	}
	return lineSearcher{
		normalize: normalize,
		graphemes: stringToGraphemes(normalize(query)),
		width:     uniseg.StringWidth(query),
	}
}

func (s lineSearcher) positions(line []cell, y int) []SearchPosition {
	var result []SearchPosition
	x := 0
	for startIdx, cell := range line {
		found := true
		for i, c := range s.graphemes {
			if len(line)-1 < startIdx+i || s.normalize(line[startIdx+i].chr) != c {
				found = false
				break
			}
		}
		if found {
			result = append(result, SearchPosition{XStart: x, XEnd: x + s.width, Y: y})
		}
		x += cell.width
	}
	return result
}
