package gocui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderTextAreaMatchesWriter(t *testing.T) {
	for _, editable := range []bool{false, true} {
		for _, autoLinks := range []bool{false, true} {
			for _, wrapWidth := range []int{0, 8} {
				t.Run(fmt.Sprintf("editable=%t/links=%t/wrap=%d", editable, autoLinks, wrapWidth), func(t *testing.T) {
					got := NewView("editor", 0, 0, 20, 5, OutputNormal)
					want := NewView("writer", 0, 0, 20, 5, OutputNormal)
					for _, v := range []*View{got, want} {
						v.Editable = editable
						v.AutoRenderHyperLinks = autoLinks
						v.FgColor = ColorRed
						v.BgColor = ColorBlue
						v.TextArea.AutoWrap = wrapWidth > 0
						v.TextArea.AutoWrapWidth = wrapWidth
					}
					for _, text := range []string{
						"some text and more text",
						"界 é 👩‍💻\nnext line\n",
						"\n\n",
						"",
						"one\ttwo\nthree\rFOUR",
						"one\r\ntwo",
						"\x1b[31mred\x1b[0m\n",
						"\x1b]8;;https://example.com\x1b\\link\x1b]8;;\x1b\\",
						osc1717("1;a;1;;file.go") + "added\n",
						"https://example.com/one\n",
						"invalid \xff byte",
					} {
						got.TextArea.Clear()
						got.TextArea.TypeString(text)
						got.RenderTextArea()
						want.SetContent(got.TextArea.GetContent())
						assert.Equal(t, want.buf.lines, got.buf.lines, "content %q", text)
						assert.Equal(t, want.buf.wx, got.buf.wx)
						assert.Equal(t, want.buf.wy, got.buf.wy)
						assert.Equal(t, want.buf.pendingNewline, got.buf.pendingNewline)
						assert.Equal(t, want.buf.ei.screenRow, got.buf.ei.screenRow)
						assert.Equal(t, want.buf.ei.screenCol, got.buf.ei.screenCol)

						got.WriteString(" appended\n界\tend")
						want.WriteString(" appended\n界\tend")
						assert.Equal(t, want.buf.lines, got.buf.lines, "append after %q", text)

						got.SetWritePos(0, 0)
						want.SetWritePos(0, 0)
						got.WriteString(strings.Repeat("overwritten", 50))
						want.WriteString(strings.Repeat("overwritten", 50))
						assert.Equal(t, want.buf.lines, got.buf.lines, "overwrite after %q", text)
					}
				})
			}
		}
	}
}
