package gocui

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSearchDrawViewport(b *testing.B) {
	for _, count := range []int{100, 10000} {
		for _, fixture := range []struct {
			name, line, query string
		}{
			{"sparse", "some text with a needle at the end\n", "needle"},
			{"dense", strings.Repeat("e ", 45) + "\n", "e"},
			{"overlap", strings.Repeat("a", 90) + "\n", "aa"},
			{"unicode", strings.Repeat("界é ", 25) + "\n", "é"},
			{"miss", "some text without the requested term\n", "needle"},
		} {
			for _, wrap := range []bool{false, true} {
				for _, position := range []string{"start", "middle", "end"} {
					b.Run(fmt.Sprintf("lines=%d/%s/wrap=%t/%s", count, fixture.name, wrap, position), func(b *testing.B) {
						saved := Screen
						g, err := NewGui(NewGuiOpts{OutputMode: OutputNormal, Headless: true, Width: 80, Height: 24})
						if err != nil {
							b.Fatal(err)
						}
						b.Cleanup(func() {
							g.Close()
							Screen = saved
						})
						v := NewView("diff", 0, 0, 79, 23, OutputNormal)
						v.Wrap = wrap
						v.SetContent(strings.Repeat(fixture.line, count))
						v.Search(fixture.query, nil)
						height := v.ViewLinesHeight()
						origin := 0
						if position == "middle" {
							origin = height / 2
						} else if position == "end" {
							origin = max(0, height-v.InnerHeight())
						}
						v.SetOriginY(origin)
						v.draw(true)
						b.ReportAllocs()
						b.ResetTimer()
						for b.Loop() {
							v.draw(true)
						}
					})
				}
			}
		}
	}
}

func BenchmarkSearchDrawHorizontalScroll(b *testing.B) {
	for _, columns := range []int{1000, 100000} {
		for _, rows := range []int{1, 22} {
			for _, position := range []string{"start", "middle", "end"} {
				b.Run(fmt.Sprintf("columns=%d/rows=%d/%s", columns, rows, position), func(b *testing.B) {
					saved := Screen
					g, err := NewGui(NewGuiOpts{OutputMode: OutputNormal, Headless: true, Width: 80, Height: 24})
					if err != nil {
						b.Fatal(err)
					}
					b.Cleanup(func() {
						g.Close()
						Screen = saved
					})
					v := NewView("diff", 0, 0, 79, 23, OutputNormal)
					v.SetContent(strings.Repeat(strings.Repeat("e ", columns/2)+"\n", rows))
					v.Search("e", nil)
					origin := 0
					if position == "middle" {
						origin = columns / 2
					} else if position == "end" {
						origin = columns - v.InnerWidth()
					}
					v.SetOriginX(origin)
					v.draw(true)
					b.ReportAllocs()
					b.ResetTimer()
					for b.Loop() {
						v.draw(true)
					}
				})
			}
		}
	}
}
