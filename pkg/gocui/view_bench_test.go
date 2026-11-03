package gocui

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkViewDraw(b *testing.B) {
	for _, size := range []struct {
		name          string
		width, height int
	}{
		{"80x24", 80, 24},
		{"160x48", 160, 48},
	} {
		for _, content := range []struct {
			name, text string
		}{
			{"full", strings.Repeat(strings.Repeat("diff content ", 20)+"\n", 100)},
			{"short", "short content"},
			{"empty", ""},
			{"unicode", strings.Repeat(strings.Repeat("界é ", 40)+"\n", 100)},
		} {
			b.Run(size.name+"/"+content.name, func(b *testing.B) {
				saved := Screen
				g, err := NewGui(NewGuiOpts{
					OutputMode: OutputNormal,
					Headless:   true,
					Width:      size.width,
					Height:     size.height,
				})
				if !assert.NoError(b, err) {
					b.FailNow()
				}
				b.Cleanup(func() {
					g.Close()
					Screen = saved
				})
				v := NewView("diff", 0, 0, size.width-1, size.height-1, OutputNormal)
				v.SetContent(content.text)
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

func BenchmarkGuiFlush(b *testing.B) {
	for _, size := range []struct {
		name          string
		width, height int
	}{
		{"80x24", 80, 24},
		{"160x48", 160, 48},
	} {
		b.Run(size.name, func(b *testing.B) {
			saved := Screen
			g, err := NewGui(NewGuiOpts{
				OutputMode: OutputNormal,
				Headless:   true,
				Width:      size.width,
				Height:     size.height,
			})
			if !assert.NoError(b, err) {
				b.FailNow()
			}
			b.Cleanup(func() {
				g.Close()
				Screen = saved
			})
			list, err := g.SetView("commits", 0, 0, size.width/3, size.height-1, 0)
			if !assert.ErrorIs(b, err, ErrUnknownView) {
				b.FailNow()
			}
			list.Title = "Commits"
			list.Highlight = true
			list.SelBgColor = ColorBlue
			list.SetContent(strings.Repeat("\x1b[33mabc1234\x1b[0m Commit subject\n", 100))
			diff, err := g.SetView("diff", size.width/3+1, 0, size.width-1, size.height-1, 0)
			if !assert.ErrorIs(b, err, ErrUnknownView) {
				b.FailNow()
			}
			diff.Title = "Diff"
			diff.SetContent(strings.Repeat("\x1b[32m+ added content\x1b[0m\n\x1b[31m- removed content 界é\x1b[0m\n", 100))
			if !assert.NoError(b, g.flush()) {
				b.FailNow()
			}
			position := 0
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				list.SetCursorY(position % (size.height - 2))
				diff.SetOriginY(position % 100)
				if err := g.flush(); err != nil {
					b.Fatal(err)
				}
				position++
			}
		})
	}
}
