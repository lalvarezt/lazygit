package gocui

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextAreaLayoutAcrossGraphemeBoundaries(t *testing.T) {
	for _, width := range []int{5, 72} {
		area := &TextArea{AutoWrap: true, AutoWrapWidth: width}
		area.TypeString(strings.Repeat("words in a paragraph ", 10))
		for _, text := range []string{
			"e", "\u0301", "\u0308", " ",
			"👩", "\u200d", "💻", " ",
			"🇦", "🇧", "🇨", "🇩", "🇪", " ",
			"👍", "🏻", " ", "☀", "\ufe0f", " ",
			"\r", "\n", "\xff", "\n\n", "Signed-off-by: Person", "\n", "\n", "body",
		} {
			area.TypeString(text)
			assertTextAreaLayout(t, area)
		}
		for range 20 {
			area.BackSpaceChar()
			assertTextAreaLayout(t, area)
		}
	}
}

func assertTextAreaLayout(t *testing.T, area *TextArea) {
	t.Helper()
	width := -1
	if area.AutoWrap {
		width = area.AutoWrapWidth
	}
	want, _ := contentToCells(area.content, width)
	if len(want) == 0 {
		assert.Empty(t, area.cells)
	} else {
		assert.Equal(t, want, area.cells)
	}
	reference := &TextArea{content: area.content, cells: want, cursor: area.cursor}
	x, y := area.GetCursorXY()
	wantX, wantY := reference.GetCursorXY()
	assert.Equal(t, wantX, x)
	assert.Equal(t, wantY, y)
	assert.Equal(t, reference.GetContent(), area.GetContent())
}

func TestTextAreaLayoutAfterEdits(t *testing.T) {
	area := &TextArea{AutoWrap: true, AutoWrapWidth: 12}
	area.TypeString("subject\n\na body with several words\n\nSigned-off-by: Someone <name@example.com>\n")
	assertTextAreaLayout(t, area)
	area.TypeString("\nnew last paragraph")
	assertTextAreaLayout(t, area)
	area.BackSpaceWord()
	assertTextAreaLayout(t, area)
	area.GoToStartOfLine()
	area.DeleteToStartOfLine()
	assertTextAreaLayout(t, area)
	area.TypeString("界 é 👩‍💻\n[1]: a footnote that should stay on one line\n")
	assertTextAreaLayout(t, area)
	area.AutoWrapWidth = 5
	area.updateCells()
	assertTextAreaLayout(t, area)
	area.AutoWrap = false
	area.updateCells()
	assertTextAreaLayout(t, area)
	area.Clear()
	area.TypeString("carriage\r\nreturn\tand tabs")
	assertTextAreaLayout(t, area)
}

func TestTextAreaLayoutThroughEditorKeys(t *testing.T) {
	rng := rand.New(rand.NewSource(491))
	keys := []Key{
		NewKeyRune('a'), NewKeyRune(' '), NewKeyRune('界'), NewKeyRune('é'),
		NewKeyName(KeyEnter), NewKeyName(KeyBackspace), NewKeyName(KeyDelete),
		NewKeyName(KeyArrowLeft), NewKeyName(KeyArrowRight),
		NewKeyName(KeyArrowUp), NewKeyName(KeyArrowDown),
		NewKeyName(KeyHome), NewKeyName(KeyEnd), NewKeyName(KeyInsert),
		NewKeyStrMod("u", ModCtrl), NewKeyStrMod("k", ModCtrl),
		NewKeyStrMod("w", ModCtrl), NewKeyStrMod("y", ModCtrl),
	}
	v := NewView("editor", 0, 0, 30, 8, OutputNormal)
	v.Editable = true
	v.TextArea.AutoWrap = true
	v.TextArea.AutoWrapWidth = 12
	v.TextArea.TypeString("subject\n\na body with words\n\nSigned-off-by: Someone\n")
	for range 1000 {
		key := keys[rng.Intn(len(keys))]
		assert.True(t, SimpleEditor(v, key))
		assertTextAreaLayout(t, v.TextArea)
	}
}
