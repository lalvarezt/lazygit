package gocui

import (
	"strings"

	"github.com/rivo/uniseg"
)

type lineSearcher struct {
	normalize func(string) string
	graphemes []string
	width     int
	prefix    []int
	starts    []int
}

func newLineSearcher(query string) lineSearcher {
	normalize := strings.ToLower
	if containsUpcaseChar(query) {
		normalize = func(s string) string { return s }
	}
	s := lineSearcher{
		normalize: normalize,
		graphemes: stringToGraphemes(normalize(query)),
		width:     uniseg.StringWidth(query),
	}
	s.prefix = make([]int, len(s.graphemes))
	s.starts = make([]int, len(s.graphemes))
	for i, matched := 1, 0; i < len(s.graphemes); i++ {
		for matched > 0 && s.graphemes[i] != s.graphemes[matched] {
			matched = s.prefix[matched-1]
		}
		if s.graphemes[i] == s.graphemes[matched] {
			matched++
		}
		s.prefix[i] = matched
	}
	return s
}

func (s lineSearcher) positions(line []cell, y int) []SearchPosition {
	var result []SearchPosition
	if len(line) < len(s.graphemes) {
		return result
	}
	x := 0
	matched := 0
	for i, cell := range line {
		s.starts[i%len(s.graphemes)] = x
		chr := s.normalize(cell.chr)
		for matched > 0 && chr != s.graphemes[matched] {
			matched = s.prefix[matched-1]
		}
		if chr == s.graphemes[matched] {
			matched++
		}
		if matched == len(s.graphemes) {
			start := s.starts[(i-matched+1)%len(s.graphemes)]
			result = append(result, SearchPosition{XStart: start, XEnd: start + s.width, Y: y})
			matched = s.prefix[matched-1]
		}
		x += cell.width
	}
	return result
}
