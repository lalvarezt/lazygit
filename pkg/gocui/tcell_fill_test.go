package gocui

import (
	"testing"
	"unicode/utf8"

	"github.com/gdamore/tcell/v3"
	"github.com/stretchr/testify/assert"
)

func TestTcellFillAreaRuneEncoding(t *testing.T) {
	var buffer tcell.CellBuffer
	buffer.Resize(1, 1)
	for _, r := range []rune{-1, 0, ' ', '~', utf8.RuneSelf - 1, utf8.RuneSelf, 'é', '界', utf8.MaxRune, utf8.MaxRune + 1} {
		buffer.FillArea(0, 0, 1, 1, r, tcell.StyleDefault)
		content, style, width := buffer.Get(0, 0)
		assert.Equal(t, string(r), content, "rune %U", r)
		assert.Equal(t, tcell.StyleDefault, style)
		assert.Equal(t, 1, width)
	}
}

func TestTcellFillAreaASCIIAllocations(t *testing.T) {
	var buffer tcell.CellBuffer
	buffer.Resize(1, 1)
	for r := rune(0); r < utf8.RuneSelf; r++ {
		allocs := testing.AllocsPerRun(10, func() {
			buffer.FillArea(0, 0, 1, 1, r, tcell.StyleDefault)
		})
		assert.Zero(t, allocs, "rune %U", r)
	}
}

func BenchmarkTcellFillArea(b *testing.B) {
	for _, size := range []struct {
		name          string
		width, height int
	}{
		{"1x1", 1, 1},
		{"160x48", 160, 48},
	} {
		for _, fill := range []struct {
			name string
			r    rune
		}{
			{"ascii", ' '},
			{"unicode", 'é'},
		} {
			b.Run(size.name+"/"+fill.name, func(b *testing.B) {
				var buffer tcell.CellBuffer
				buffer.Resize(size.width, size.height)
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					buffer.FillArea(0, 0, size.width, size.height, fill.r, tcell.StyleDefault)
				}
			})
		}
	}
}
