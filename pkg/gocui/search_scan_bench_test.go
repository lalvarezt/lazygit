package gocui

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSearchRepeatedPrefix(b *testing.B) {
	for _, query := range []string{"a", strings.Repeat("a", 31) + "b", strings.Repeat("a", 127) + "b"} {
		b.Run(fmt.Sprintf("query=%d", len(query)), func(b *testing.B) {
			v := NewView("search", 0, 0, 80, 24, OutputNormal)
			v.SetContent(strings.Repeat(strings.Repeat("a", 200)+"\n", 2000))
			v.UpdateSearchResults(query, nil)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				v.UpdateSearchResults(query, nil)
			}
		})
	}
}

func BenchmarkSearchOrdinaryText(b *testing.B) {
	for _, query := range []string{"match", "missing", "MATCH"} {
		b.Run(query, func(b *testing.B) {
			v := NewView("search", 0, 0, 80, 24, OutputNormal)
			v.SetContent(strings.Repeat("line of a diff, most of which does not match\n", 2000))
			v.UpdateSearchResults(query, nil)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				v.UpdateSearchResults(query, nil)
			}
		})
	}
}

func BenchmarkNearestSearchPosition(b *testing.B) {
	for _, count := range []int{100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("matches=%d", count), func(b *testing.B) {
			v := NewView("search", 0, 0, 80, 24, OutputNormal)
			v.searcher.searchPositions = make([]SearchPosition, count)
			for i := range count {
				v.searcher.searchPositions[i] = SearchPosition{XStart: i % 2, XEnd: i%2 + 1, Y: i / 2}
			}
			v.oy = count/2 - 1
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				v.nearestSearchPosition()
			}
		})
	}
}

func BenchmarkScrollSearchPosition(b *testing.B) {
	for _, count := range []int{100, 1000, 10000, 100000} {
		for _, up := range []bool{false, true} {
			b.Run(fmt.Sprintf("matches=%d/up=%t", count, up), func(b *testing.B) {
				v := NewView("search", 0, 0, 80, 24, OutputNormal)
				v.searcher.searchPositions = make([]SearchPosition, count)
				for i := range count {
					v.searcher.searchPositions[i] = SearchPosition{Y: i}
				}
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					if up {
						v.oy = 0
						v.searcher.currentSearchIndex = count - 1
						v.selectVisibleSearchResultAfterScrollUp()
					} else {
						v.oy = count - 10
						v.searcher.currentSearchIndex = 0
						v.selectVisibleSearchResultAfterScrollDown()
					}
				}
			})
		}
	}
}
