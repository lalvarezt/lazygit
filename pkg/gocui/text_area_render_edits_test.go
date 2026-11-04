package gocui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEditorMatchesWriter(t *testing.T, got *View) {
	t.Helper()
	want := NewView("writer", 0, 0, 20, 5, OutputNormal)
	want.Editable = got.Editable
	want.Wrap = got.Wrap
	want.AutoRenderHyperLinks = got.AutoRenderHyperLinks
	want.buf.ei.screenColMax = got.buf.ei.screenColMax
	want.SetContent(got.TextArea.GetContent())
	if !assert.Len(t, got.buf.lines, len(want.buf.lines)) {
		return
	}
	for i := range want.buf.lines {
		assert.Equal(t, want.buf.lines[i].cells, got.buf.lines[i].cells)
		assert.Equal(t, want.buf.lines[i].metadataCollected, got.buf.lines[i].metadataCollected)
		assert.Equal(t, want.buf.lines[i].asWritten, got.buf.lines[i].asWritten)
	}
	assert.Equal(t, want.buf.wx, got.buf.wx)
	assert.Equal(t, want.buf.wy, got.buf.wy)
	assert.Equal(t, want.buf.pendingNewline, got.buf.pendingNewline)
	assert.Equal(t, want.buf.ei.screenRow, got.buf.ei.screenRow)
	assert.Equal(t, want.buf.ei.screenCol, got.buf.ei.screenCol)
	assert.Equal(t, want.ViewLinesHeight(), got.ViewLinesHeight())
}

func TestRenderTextAreaAfterEdits(t *testing.T) {
	for _, editable := range []bool{false, true} {
		v := NewView("editor", 0, 0, 20, 5, OutputNormal)
		v.Editable = editable
		v.Wrap = true
		v.TextArea.AutoWrap = true
		v.TextArea.AutoWrapWidth = 8
		for _, text := range []string{"first line\n", "a body with many words\n", "界é👩‍💻\n", "\n", "Signed-off-by: Person <person@example.com>", "\n\n", "a final paragraph"} {
			for _, r := range text {
				v.TextArea.TypeCharacter(string(r))
				v.RenderTextArea()
				assertEditorMatchesWriter(t, v)
			}
		}
		v.TextArea.SetCursor2D(0, 1)
		v.TextArea.DeleteToEndOfLine()
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
		v.TextArea.TypeString("one")
		v.TextArea.TypeString("two")
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
		v.TextArea.AutoWrapWidth = 5
		v.TextArea.updateCells()
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
		v.SetContentWidth(4)
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
		v.Editable = !v.Editable
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
		v.AutoRenderHyperLinks = true
		v.TextArea.TypeString("\nhttps://example.com")
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
		v.AutoRenderHyperLinks = false
		v.TextArea.TypeString("\tcontrol\r\n")
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
	}
}

func TestRenderTextAreaRestoresExternallyChangedBuffer(t *testing.T) {
	for _, mutate := range []func(*View){
		func(v *View) { v.WriteString("external") },
		func(v *View) { v.SetContent("replacement") },
		func(v *View) { v.Clear() },
		func(v *View) { v.Reset() },
		func(v *View) { v.OverwriteLinesAndClearEverythingElse(1, 0, "overwritten") },
		func(v *View) {
			v.BeginOffscreenRender()
			v.WriteString("offscreen")
			v.SwapInOffscreenRender()
		},
		func(v *View) {
			other := NewView("other", 0, 0, 20, 5, OutputNormal)
			other.SetContent("copied")
			v.CopyContent(other)
		},
	} {
		v := NewView("editor", 0, 0, 20, 5, OutputNormal)
		v.TextArea.TypeString("first\nsecond\nthird")
		v.RenderTextArea()
		mutate(v)
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
	}
}

func TestEditedTextAreaPreservesCopiedContent(t *testing.T) {
	v := NewView("editor", 0, 0, 20, 5, OutputNormal)
	v.TextArea.TypeString("first\n界 é\nlast")
	v.RenderTextArea()
	copy := NewView("copy", 0, 0, 20, 5, OutputNormal)
	copy.CopyContent(v)
	want := copy.Buffer()
	v.TextArea.SetCursor2D(0, 1)
	v.TextArea.DeleteToEndOfLine()
	v.TextArea.TypeString("changed")
	v.RenderTextArea()
	assert.Equal(t, want, copy.Buffer())
	assertEditorMatchesWriter(t, v)
}

func TestEditedTextAreaPreservesMultipleSnapshots(t *testing.T) {
	for i, text := range []string{
		strings.Repeat("abcdefgh", 40),
		strings.Repeat("界é👩‍💻", 40),
		"first\n" + strings.Repeat("body text ", 40) + "\nlast",
	} {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			v := NewView("editor", 0, 0, 20, 5, OutputNormal)
			v.Wrap = true
			v.TextArea.AutoWrap = true
			v.TextArea.AutoWrapWidth = 72
			v.TextArea.TypeString(text)
			v.RenderTextArea()
			type snapshot struct {
				view       *View
				buffer     string
				viewBuffer string
			}
			var snapshots []snapshot
			for step := range 18 {
				if step%3 == 0 {
					v.ViewLinesHeight()
					copy := NewView(fmt.Sprintf("copy%d", step), 0, 0, 20, 5, OutputNormal)
					copy.Wrap = true
					copy.CopyContent(v)
					snapshots = append(snapshots, snapshot{copy, copy.Buffer(), copy.ViewBuffer()})
				}
				switch step % 3 {
				case 0:
					v.TextArea.TypeString("界x")
				case 1:
					v.TextArea.SetCursor2D(3, step%4)
					v.TextArea.DeleteToEndOfLine()
					v.TextArea.TypeString("replacement")
				case 2:
					v.TextArea.TypeString("\nnew row\n")
				}
				v.RenderTextArea()
				assertEditorMatchesWriter(t, v)
				for _, snapshot := range snapshots {
					assert.Equal(t, snapshot.buffer, snapshot.view.Buffer())
					assert.Equal(t, snapshot.viewBuffer, snapshot.view.ViewBuffer())
				}
			}
		})
	}
}
