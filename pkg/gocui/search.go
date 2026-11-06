package gocui

import (
	"strings"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

type lineSearcher struct {
	caseSensitive bool
	graphemes     []string
	width         int
	prefix        []int
	starts        []int
}

func newLineSearcher(query string) lineSearcher {
	caseSensitive := containsUpcaseChar(query)
	normalized := query
	if !caseSensitive {
		normalized = strings.ToLower(query)
	}
	s := lineSearcher{
		caseSensitive: caseSensitive,
		graphemes:     stringToGraphemes(normalized),
		width:         uniseg.StringWidth(query),
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
	for i := range line {
		chr := line[i].chr
		if !s.caseSensitive {
			if len(chr) == 1 && chr[0] < utf8.RuneSelf {
				if chr[0] >= 'A' && chr[0] <= 'Z' {
					chr = string([]byte{chr[0] + ('a' - 'A')})
				}
			} else {
				chr = strings.ToLower(chr)
			}
		}
		for matched > 0 && chr != s.graphemes[matched] {
			matched = s.prefix[matched-1]
		}
		if chr == s.graphemes[matched] {
			s.starts[i%len(s.graphemes)] = x
			matched++
		}
		if matched == len(s.graphemes) {
			start := s.starts[(i-matched+1)%len(s.graphemes)]
			result = append(result, SearchPosition{XStart: start, XEnd: start + s.width, Y: y})
			matched = s.prefix[matched-1]
		}
		x += line[i].width
	}
	return result
}
