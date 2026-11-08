package gocui

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamingSearchMatchesCompleteSearch(t *testing.T) {
	for _, wrap := range []bool{false, true} {
		v := NewView("diff", 0, 0, 24, 40, OutputTrue)
		v.Wrap = wrap
		v.Search("needle", nil)
		var input strings.Builder
		for _, chunk := range []string{
			"needle first\n", "a long line without the term to be wrapped around the edge\n",
			"needle", " and another needle\n", "\x1b[32mneedle 界 é\x1b[m\n",
			"old needle\rnew content\n", "tail needle",
		} {
			input.WriteString(chunk)
			v.WriteString(chunk)
			v.RefreshSearch()
			want := NewView("diff", 0, 0, 24, 40, OutputTrue)
			want.Wrap = wrap
			want.SetContent(input.String())
			want.Search("needle", nil)
			assert.Equal(t, want.searcher.searchPositions, v.searcher.searchPositions)
		}
		for _, width := range []int{12, 40, 24} {
			v.x1 = width
			v.RewrapContent()
			v.RefreshSearch()
			want := NewView("diff", 0, 0, width, 40, OutputTrue)
			want.Wrap = wrap
			want.SetContent(input.String())
			want.Search("needle", nil)
			assert.Equal(t, want.searcher.searchPositions, v.searcher.searchPositions)
		}
		v.BeginOffscreenRender()
		v.WriteString("new needle\n")
		v.SwapInOffscreenRender()
		v.RefreshSearch()
		assert.Equal(t, []SearchPosition{{XStart: 4, XEnd: 10, Y: 0}}, v.searcher.searchPositions)
		v.UpdateSearchResults("new", nil)
		assert.Equal(t, []SearchPosition{{XStart: 0, XEnd: 3, Y: 0}}, v.searcher.searchPositions)
	}
}

func BenchmarkIncrementalSearchedDiff(b *testing.B) {
	chunk := []byte(strings.Repeat("line with a needle and more diff content\n", 128))
	for _, wrap := range []bool{false, true} {
		name := "plain"
		if wrap {
			name = "wrapped"
		}
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				v := NewView("diff", 0, 0, 24, 40, OutputNormal)
				v.Wrap = wrap
				v.Search("needle", nil)
				for range 64 {
					_, _ = v.Write(chunk)
					v.RefreshSearch()
				}
			}
		})
	}
}
