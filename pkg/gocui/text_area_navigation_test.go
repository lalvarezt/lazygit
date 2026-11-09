package gocui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sequentialCursorPosition(area *TextArea, x, y int) int {
	x, y = max(x, 0), max(y, 0)
	for i, c := range area.cells {
		if y == 0 && (x <= 0 || c.char == "\n") {
			area.cursor = area.cellCursorToContentCursor(i)
			if c.char == "\n" {
				area.moveLeftFromSoftLineBreak()
			}
			return area.cursor
		}
		if c.char == "\n" {
			y--
		} else if y == 0 {
			x -= c.width
		}
	}
	if y == 0 {
		return len(area.content)
	}
	return area.cursor
}

func TestSetCursor2DMatchesSequentialNavigation(t *testing.T) {
	for _, content := range []string{
		"", "abc", "\n\n", "abc\n", "abc\n\ndef\nlast",
		"abc def ghi jkl", "界é abc 界\n\u200bnext", "\ttext\r\nlast",
		"paragraph text\n\nSigned-off-by: Some User <user@example.com>",
	} {
		for _, width := range []int{0, 1, 4, 12} {
			t.Run(fmt.Sprintf("%q/width=%d", content, width), func(t *testing.T) {
				area := &TextArea{content: content, AutoWrap: width > 0, AutoWrapWidth: width}
				area.updateCells()
				for _, cursor := range []int{0, len(content)} {
					for y := -1; y <= len(area.cells)+1; y++ {
						for x := -1; x <= 14; x++ {
							area.cursor = cursor
							reference := *area
							want := sequentialCursorPosition(&reference, x, y)
							area.SetCursor2D(x, y)
							assert.Equal(t, want, area.cursor, "cursor=%d, x=%d, y=%d", cursor, x, y)
						}
					}
				}
			})
		}
	}
}

func BenchmarkTextAreaVerticalMovement(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		for _, bottom := range []bool{false, true} {
			b.Run(fmt.Sprintf("lines=%d/bottom=%t", count, bottom), func(b *testing.B) {
				area := &TextArea{content: strings.Repeat("some text for a line\n", count)}
				area.updateCells()
				row := 2
				if bottom {
					row = count - 2
				}
				area.SetCursor2D(5, row)
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					area.MoveCursorUp()
					area.MoveCursorDown()
				}
			})
		}
	}
}

func BenchmarkTextAreaVerticalKeypress(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("lines=%d", count), func(b *testing.B) {
			v := NewView("description", 0, 0, 80, 24, OutputNormal)
			v.Editable = true
			v.TextArea.TypeString(strings.Repeat("some text for a line\n", count))
			v.TextArea.SetCursor2D(5, count-2)
			v.RenderTextArea()
			up, down := NewKeyName(KeyArrowUp), NewKeyName(KeyArrowDown)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				SimpleEditor(v, up)
				SimpleEditor(v, down)
			}
		})
	}
}

func BenchmarkTextAreaLineStart(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprintf("lines=%d", count), func(b *testing.B) {
			area := &TextArea{content: strings.Repeat("some text for a line\n", count)}
			area.updateCells()
			cursor := len(area.content) - 3
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				area.cursor = cursor
				area.GoToStartOfLine()
			}
		})
	}
}
