package gocui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDrawSearchedView(b *testing.B) {
	for _, lines := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("lines=%d", lines), func(b *testing.B) {
			saved := Screen
			g, err := NewGui(NewGuiOpts{
				OutputMode: OutputNormal,
				Headless:   true,
				Width:      80,
				Height:     24,
			})
			if !assert.NoError(b, err) {
				b.FailNow()
			}
			b.Cleanup(func() {
				g.Close()
				Screen = saved
			})
			v := NewView("diff", 0, 0, 79, 23, OutputNormal)
			v.SetContent(strings.Repeat("some match text and another match\n", lines))
			v.Search("match", nil)
			v.SetOriginY(lines / 2)
			v.draw(true)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				v.draw(true)
			}
		})
	}
}
