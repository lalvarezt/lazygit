package gocui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawSearchAfterStreamingAndReplacement(t *testing.T) {
	for _, wrap := range []bool{false, true} {
		t.Run(fmt.Sprintf("wrap=%t", wrap), func(t *testing.T) {
			WithSimulationScreen(t, 14, 6)
			v := NewView("diff", 0, 0, 11, 5, OutputNormal)
			v.Wrap = wrap
			v.InclusionGutterMarker = "✓"
			v.SetInclusionGutter(true, []bool{true, false, true})
			v.Search("aa", nil)
			var input strings.Builder
			check := func(query string) {
				t.Helper()
				v.UpdateSearchResults(query, nil)
				want := NewView("diff", 0, 0, 11, 5, OutputNormal)
				want.Wrap = wrap
				want.InclusionGutterMarker = "✓"
				want.SetInclusionGutter(true, []bool{true, false, true})
				want.SetContent(input.String())
				want.Search(query, nil)
				assert.Equal(t, want.searcher.searchPositions, v.searcher.searchPositions)
				for _, origin := range []int{0, 1, 3} {
					v.SetOrigin(2, origin)
					want.SetOrigin(2, origin)
					for index := range v.searcher.searchPositions {
						v.searcher.currentSearchIndex = index
						want.searcher.currentSearchIndex = index
						assert.Equal(t, drawSearchHighlights(want), drawSearchHighlights(v))
					}
					if len(v.searcher.searchPositions) == 0 {
						assert.Equal(t, drawSearchHighlights(want), drawSearchHighlights(v))
					}
				}
			}
			for _, chunk := range []string{"aaaa 界é\n", "long line aaaa aaaa\n", "aaaa", " and é\n", "tail aa"} {
				input.WriteString(chunk)
				v.WriteString(chunk)
				check("aa")
			}
			check("é")
			check("absent")
			v.BeginOffscreenRender()
			v.WriteString("界é aaa\nreplacement aaa")
			v.SwapInOffscreenRender()
			input.Reset()
			input.WriteString("界é aaa\nreplacement aaa")
			check("aa")
		})
	}
}
