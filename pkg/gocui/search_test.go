package gocui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gdamore/tcell/v3/color"
	"github.com/stretchr/testify/assert"
)

// writeLines writes the given lines to the view, as a task rendering content into it
// does: one line at a time.
func writeLines(v *View, lines ...string) {
	for _, line := range lines {
		fmt.Fprintf(v, "%s\n", line)
	}
}

func TestSearchStatusAfterTheMatchesChange(t *testing.T) {
	v := NewView("name", 0, 0, 40, 10, OutputNormal)
	writeLines(v, "match", "other", "match", "other", "match")

	v.Search("match", nil)
	_ = v.gotoNextMatch()
	_ = v.gotoNextMatch()
	index, total := v.GetSearchStatus()
	assert.Equal(t, 2, index)
	assert.Equal(t, 3, total)

	// The content is re-rendered with only the first of those matches left in it.
	v.Clear()
	writeLines(v, "match", "other", "other")

	index, total = v.GetSearchStatus()
	assert.Equal(t, 0, index)
	assert.Equal(t, 1, total)
}

func TestSearchPositionsFollowStreamedContent(t *testing.T) {
	v := NewView("name", 0, 0, 40, 10, OutputNormal)
	v.Search("match", nil)

	// A render arrives a line at a time, and the status describes all of it.
	writeLines(v, "other", "match", "other", "match")

	_, total := v.GetSearchStatus()
	assert.Equal(t, 2, total)
}

func TestSearchPositionsWithOverlappingPrefixes(t *testing.T) {
	for _, test := range []struct {
		name    string
		content string
		query   string
		want    []SearchPosition
	}{
		{name: "overlaps", content: "aaaa", query: "aa", want: []SearchPosition{{XStart: 0, XEnd: 2}, {XStart: 1, XEnd: 3}, {XStart: 2, XEnd: 4}}},
		{name: "fallback", content: "aaaabaaab", query: "aaab", want: []SearchPosition{{XStart: 1, XEnd: 5}, {XStart: 5, XEnd: 9}}},
		{name: "nested prefixes", content: "abababaca", query: "ababaca", want: []SearchPosition{{XStart: 2, XEnd: 9}}},
		{name: "wide cells", content: "界界界", query: "界界", want: []SearchPosition{{XStart: 0, XEnd: 4}, {XStart: 2, XEnd: 6}}},
		{name: "mixed-width fallback", content: "界a界a界b", query: "界a界b", want: []SearchPosition{{XStart: 3, XEnd: 9}}},
		{name: "mixed-width overlaps", content: "界a界a界", query: "界a界", want: []SearchPosition{{XStart: 0, XEnd: 5}, {XStart: 3, XEnd: 8}}},
		{name: "ASCII case folding", content: "xMaTcH MATCH match", query: "match", want: []SearchPosition{{XStart: 1, XEnd: 6}, {XStart: 7, XEnd: 12}, {XStart: 13, XEnd: 18}}},
		{name: "ASCII case sensitive", content: "xMaTcH MATCH match", query: "MATCH", want: []SearchPosition{{XStart: 7, XEnd: 12}}},
		{name: "combining marks", content: "ééé", query: "éé", want: []SearchPosition{{XStart: 0, XEnd: 2}, {XStart: 1, XEnd: 3}}},
		{name: "lowercase", content: "Éé", query: "é", want: []SearchPosition{{XStart: 0, XEnd: 1}, {XStart: 1, XEnd: 2}}},
		{name: "uppercase", content: "Éé", query: "É", want: []SearchPosition{{XStart: 0, XEnd: 1}}},
		{name: "case mapping", content: "Kkk", query: "kk", want: []SearchPosition{{XStart: 0, XEnd: 2}, {XStart: 1, XEnd: 3}}},
		{name: "line boundary", content: "ab\nba", query: "abba", want: []SearchPosition{}},
		{name: "missing suffix", content: strings.Repeat("a", 1000), query: strings.Repeat("a", 31) + "b", want: []SearchPosition{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			v := NewView("search", 0, 0, 80, 24, OutputNormal)
			v.SetContent(test.content)
			v.UpdateSearchResults(test.query, nil)
			assert.Equal(t, test.want, v.searcher.searchPositions)
		})
	}
}

func TestSearchModelPositionsAcrossLines(t *testing.T) {
	v := NewView("search", 0, 0, 80, 24, OutputNormal)
	v.SetContent("xx xx\nx\nother xx")
	v.UpdateSearchResults("xx", []SearchPosition{
		{XStart: 20, XEnd: 22, Y: 0},
		{XStart: 0, XEnd: 1, Y: 1},
		{XStart: 20, XEnd: 22, Y: 2},
	})
	assert.Equal(t, []SearchPosition{
		{XStart: 0, XEnd: 2, Y: 0},
		{XStart: 3, XEnd: 5, Y: 0},
		{XStart: 0, XEnd: 1, Y: 1},
		{XStart: 6, XEnd: 8, Y: 2},
	}, v.searcher.searchPositions)

	v.UpdateSearchResults("xx", []SearchPosition{})
	assert.Empty(t, v.searcher.searchPositions)
}

func TestSearchResultBufferReuse(t *testing.T) {
	for _, model := range []bool{false, true} {
		t.Run(fmt.Sprintf("model=%t", model), func(t *testing.T) {
			v := NewView("search", 0, 0, 80, 24, OutputNormal)
			v.SetContent("aaaa\naaaa")
			var modelResults []SearchPosition
			if model {
				modelResults = []SearchPosition{{Y: 0}, {Y: 1}}
			}
			v.UpdateSearchResults("a", modelResults)
			assert.Len(t, v.searcher.searchPositions, 8)
			firstResult := &v.searcher.searchPositions[0]
			capacity := cap(v.searcher.searchPositions)

			v.UpdateSearchResults("aa", modelResults)
			assert.Len(t, v.searcher.searchPositions, 6)
			assert.Same(t, firstResult, &v.searcher.searchPositions[0])

			v.Clear()
			writeLines(v, "aa", "aa")
			_, total := v.GetSearchStatus()
			assert.Equal(t, 2, total)
			assert.Same(t, firstResult, &v.searcher.searchPositions[0])

			v.UpdateSearchResults("missing", nil)
			assert.Empty(t, v.searcher.searchPositions)
			assert.Equal(t, capacity, cap(v.searcher.searchPositions))

			v.UpdateSearchResults("a", nil)
			assert.Len(t, v.searcher.searchPositions, 4)
			assert.Same(t, firstResult, &v.searcher.searchPositions[0])
		})
	}
}

func TestClearSearchReleasesResultBuffer(t *testing.T) {
	for _, emptyQuery := range []bool{false, true} {
		t.Run(fmt.Sprintf("emptyQuery=%t", emptyQuery), func(t *testing.T) {
			v := NewView("search", 0, 0, 80, 24, OutputNormal)
			v.SetContent("aaaa")
			v.UpdateSearchResults("a", nil)
			assert.Positive(t, cap(v.searcher.searchPositions))

			if emptyQuery {
				v.UpdateSearchResults("", nil)
			} else {
				v.ClearSearch()
			}
			assert.False(t, v.IsSearching())
			assert.Empty(t, v.searcher.searchPositions)
			assert.Zero(t, cap(v.searcher.searchPositions))

			v.UpdateSearchResults("aa", nil)
			assert.Equal(t, []SearchPosition{
				{XStart: 0, XEnd: 2},
				{XStart: 1, XEnd: 3},
				{XStart: 2, XEnd: 4},
			}, v.searcher.searchPositions)
		})
	}
}

func TestNearestSearchPosition(t *testing.T) {
	v := NewView("search", 0, 0, 80, 24, OutputNormal)
	assert.Zero(t, v.nearestSearchPosition())
	v.searcher.searchPositions = []SearchPosition{{Y: 2}, {Y: 2}, {Y: 5}, {Y: 5}, {Y: 9}}
	for _, test := range []struct{ line, want int }{
		{0, 0}, {2, 0}, {3, 1}, {5, 2}, {8, 3}, {9, 4}, {10, 4},
	} {
		v.cy = 1
		v.oy = test.line - v.cy
		assert.Equal(t, test.want, v.nearestSearchPosition(), "line %d", test.line)
	}
}

func TestSearchSelectionAfterScrolling(t *testing.T) {
	for _, test := range []struct {
		name      string
		up        bool
		origin    int
		current   int
		want      int
		highlight bool
	}{
		{name: "down to first visible", origin: 10, current: 0, want: 1},
		{name: "up to last visible", up: true, origin: 10, current: 6, want: 3},
		{name: "down into gap", origin: 25, current: 0, want: 3},
		{name: "up into gap", up: true, origin: 25, current: 6, want: 4},
		{name: "down past last result", origin: 100, current: 0, want: 6},
		{name: "up before first result", up: true, origin: 0, current: 6, want: 3},
		{name: "down stays visible", origin: 10, current: 2, want: 2},
		{name: "up stays visible", up: true, origin: 10, current: 2, want: 2},
		{name: "down stays nearest above gap", origin: 25, current: 3, want: 3},
		{name: "up stays nearest below gap", up: true, origin: 25, current: 4, want: 4},
		{name: "highlighted down", origin: 10, current: 0, want: 0, highlight: true},
		{name: "highlighted up", up: true, origin: 10, current: 6, want: 6, highlight: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			v := NewView("search", 0, 0, 80, 24, OutputNormal)
			v.searcher.searchPositions = []SearchPosition{{Y: 0}, {Y: 10}, {Y: 20}, {Y: 20}, {Y: 50}, {Y: 50}, {Y: 80}}
			v.searcher.currentSearchIndex = test.current
			v.oy = test.origin
			v.Highlight = test.highlight
			notifications := 0
			v.setRenderSearchStatus(func(_ *View, selected, total int) {
				notifications++
				assert.Equal(t, test.want, selected)
				assert.Equal(t, 7, total)
			})
			if test.up {
				v.selectVisibleSearchResultAfterScrollUp()
			} else {
				v.selectVisibleSearchResultAfterScrollDown()
			}
			assert.Equal(t, test.want, v.searcher.currentSearchIndex)
			assert.Equal(t, test.want != test.current, notifications == 1)
		})
	}
}

func TestSearchHighlights(t *testing.T) {
	for _, test := range []struct {
		name       string
		content    string
		search     string
		selected   int
		originX    int
		originY    int
		wrap       bool
		model      []SearchPosition
		wantColors []string
	}{
		{name: "no matches", content: "some text", search: "match", wantColors: []string{"........", "........"}},
		{name: "several matches", content: "xx xx\nxx", search: "xx", selected: 1, wantColors: []string{"yy.cc...", "yy......"}},
		{name: "overlapping matches", content: "aaa", search: "aa", selected: 1, wantColors: []string{"yyc.....", "........"}},
		{name: "scrolling", content: "other\nxx match\nmatch", search: "match", selected: 1, originX: 3, originY: 1, wantColors: []string{"yyyyy...", "cc......"}},
		{name: "wrapping", content: "12345678match", search: "match", wrap: true, wantColors: []string{"........", "ccccc..."}},
		{name: "Unicode columns", content: "界ématch", search: "match", wantColors: []string{"...ccccc", "........"}},
		{name: "abbreviated model match", content: "abc\nother", search: "abcdef", model: []SearchPosition{{Y: 0, XStart: 0, XEnd: 3}}, wantColors: []string{"ccc.....", "........"}},
		{name: "unrendered model match", content: "abc", search: "abcdef", model: []SearchPosition{{Y: 0, XStart: -1, XEnd: -1}}, wantColors: []string{"........", "........"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			WithSimulationScreen(t, 10, 4)
			v := NewView("diff", 0, 0, 9, 3, OutputNormal)
			v.Wrap = test.wrap
			v.SetContent(test.content)
			v.UpdateSearchResults(test.search, test.model)
			v.searcher.currentSearchIndex = test.selected
			v.SetOrigin(test.originX, test.originY)
			v.draw(true)

			for y, want := range test.wantColors {
				var colors strings.Builder
				for x := range 8 {
					_, style, _ := Screen.Get(x+1, y+1)
					switch style.GetBackground() {
					case color.Olive:
						colors.WriteByte('y')
					case color.Teal:
						colors.WriteByte('c')
					default:
						colors.WriteByte('.')
					}
				}
				assert.Equal(t, want, colors.String())
			}
		})
	}
}

func BenchmarkWriteToSearchedView(b *testing.B) {
	for b.Loop() {
		v := NewView("name", 0, 0, 100, 40, OutputNormal)
		v.Search("match", nil)
		for i := range 2000 {
			fmt.Fprintf(v, "line %d of a diff, most of which does not match\n", i)
		}
		v.GetSearchStatus()
	}
}
