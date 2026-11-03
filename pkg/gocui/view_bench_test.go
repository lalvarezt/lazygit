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
