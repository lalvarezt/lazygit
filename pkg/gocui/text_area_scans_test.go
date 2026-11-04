package gocui

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextAreaContentScansAcrossEdits(t *testing.T) {
	contents := []string{""}
	for size, start := 0, 0; size < 5; size++ {
		end := len(contents)
		for _, prefix := range contents[start:end] {
			for _, chr := range []string{"a", "\n", "\t", "\r", "\x1b"} {
				contents = append(contents, prefix+chr)
			}
		}
		start = end
	}
	for _, old := range contents {
		for start := 0; start <= len(old); start++ {
			for end := start; end <= len(old); end++ {
				for _, inserted := range []string{"", "b", "\n", "\n\n", "\t", "\r", "\x1b"} {
					area := textAreaWithScannedContent(old)
					area.replaceContent(start, end, inserted)
					assertContentScans(t, area)
				}
			}
		}
	}
}

func TestTextAreaContentScansAcrossDeferredEdits(t *testing.T) {
	rng := rand.New(rand.NewSource(92837))
	for _, initial := range []string{
		strings.Repeat("a", 300),
		"subject\n\nbody\n\nName: https://example.com\n\n",
		"body\n\n" + strings.Repeat("\n", 300),
		strings.Repeat("\n", 300),
		"before\tbetween\rafter\x1b[31mred",
		"界é👩‍💻\n\nbody",
	} {
		area := textAreaWithScannedContent(initial)
		for range 1000 {
			for range 1 + rng.Intn(4) {
				start := rng.Intn(len(area.content) + 1)
				end := start + rng.Intn(len(area.content)-start+1)
				text := []string{"", "a", "\n", "\n\n", "\t", "\r", "\x1b", "界", "\xff"}[rng.Intn(9)]
				area.replaceContent(start, end, text)
			}
			trailerStart, firstControl := assertContentScans(t, area)
			area.layout.content = area.content
			area.layout.trailerStart = trailerStart
			area.layout.firstControl = firstControl
			area.hasChanges = false
		}
	}
}

func TestRenderTextAreaAfterControlEdits(t *testing.T) {
	for _, editable := range []bool{false, true} {
		v := NewView("editor", 0, 0, 20, 5, OutputNormal)
		v.Editable = editable
		v.Wrap = true
		v.TextArea.AutoWrap = true
		v.TextArea.AutoWrapWidth = 8
		v.RenderTextArea()
		assertEditorMatchesWriter(t, v)
		for _, content := range []string{
			"first\tline\n\nbody\rafter\x1b[31mred",
			"\t\r\x1b",
			"ordinary text",
		} {
			v.TextArea.TypeString(content)
			v.RenderTextArea()
			assertEditorMatchesWriter(t, v)
			for {
				control := strings.IndexAny(v.TextArea.content, "\t\r\x1b")
				if control < 0 {
					break
				}
				v.TextArea.replaceContent(control, control+1, "")
				v.TextArea.cursor = len(v.TextArea.content)
				v.TextArea.updateCells()
				v.RenderTextArea()
				assertEditorMatchesWriter(t, v)
			}
			v.TextArea.replaceContent(0, 0, "\t")
			v.TextArea.replaceContent(0, 1, "\n\n")
			v.TextArea.updateCells()
			v.RenderTextArea()
			assertEditorMatchesWriter(t, v)
			v.TextArea.AutoWrapWidth++
			v.TextArea.updateCells()
			v.RenderTextArea()
			assertEditorMatchesWriter(t, v)
			v.TextArea.Clear()
			v.RenderTextArea()
			assertEditorMatchesWriter(t, v)
		}
	}
}

func textAreaWithScannedContent(content string) *TextArea {
	return &TextArea{
		content: content,
		layout: textAreaLayout{
			ready:        true,
			content:      content,
			trailerStart: referenceTrailerStart(content),
			firstControl: strings.IndexAny(content, "\t\r\x1b"),
		},
	}
}

func referenceTrailerStart(content string) int {
	if i := strings.LastIndex(strings.TrimRight(content, "\n"), "\n\n"); i >= 0 {
		return i + 2
	}
	return 0
}

func assertContentScans(t *testing.T, area *TextArea) (int, int) {
	t.Helper()
	trailerStart, firstControl := area.scanContent()
	if !assert.Equal(t, referenceTrailerStart(area.content), trailerStart,
		"old=%q new=%q dirty=%d", area.layout.content, area.content, area.changedFromByte) {
		t.FailNow()
	}
	if !assert.Equal(t, strings.IndexAny(area.content, "\t\r\x1b"), firstControl,
		"old=%q new=%q dirty=%d", area.layout.content, area.content, area.changedFromByte) {
		t.FailNow()
	}
	return trailerStart, firstControl
}
