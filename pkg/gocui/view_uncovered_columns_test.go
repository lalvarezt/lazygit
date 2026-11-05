package gocui

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawClearsUncoveredColumnsAfterMalformedEscape(t *testing.T) {
	for _, highlight := range []bool{false, true} {
		t.Run(map[bool]string{false: "plain", true: "highlighted"}[highlight], func(t *testing.T) {
			WithSimulationScreen(t, 8, 5)
			v := NewView("escape", 0, 0, 7, 4, OutputNormal)
			v.FgColor = ColorGreen
			v.BgColor = ColorBlue
			v.Highlight = highlight
			v.SelBgColor = ColorYellow
			v.SetContent("abcdef")
			v.draw(true)

			v.SetContent("\x1b[12é")
			v.draw(true)

			var row strings.Builder
			for x := range 6 {
				ch, _, _ := Screen.Get(x+1, 1)
				row.WriteString(ch)
			}
			assert.Equal(t, "[1 ;é ", row.String())
			_, style, _ := Screen.Get(3, 1)
			assert.Equal(t, getTcellStyle(oldStyle{fg: v.FgColor, bg: v.BgColor, outputMode: v.outMode}), style)
		})
	}
}

func TestDrawClearsClippedWideCharacters(t *testing.T) {
	for _, test := range []struct {
		name    string
		left    int
		content string
	}{
		{name: "CJK", left: -2, content: "界"},
		{name: "emoji", left: -2, content: "👩‍💻"},
		{name: "after narrow character", left: -3, content: "a界"},
	} {
		t.Run(test.name, func(t *testing.T) {
			WithSimulationScreen(t, 8, 5)
			outsideStyle := getTcellStyle(oldStyle{fg: ColorRed, bg: ColorYellow, outputMode: OutputNormal})
			Screen.Fill('x', outsideStyle)
			v := NewView("clipped", test.left, 0, test.left+7, 4, OutputNormal)
			v.FgColor = ColorGreen
			v.BgColor = ColorBlue
			v.SetContent("abcdef")
			v.draw(true)

			v.SetContent(test.content)
			v.draw(true)

			ch, style, width := Screen.Get(0, 1)
			assert.Equal(t, " ", ch)
			assert.Equal(t, 1, width)
			assert.Equal(t, getTcellStyle(oldStyle{fg: v.FgColor, bg: v.BgColor, outputMode: v.outMode}), style)
			ch, style, _ = Screen.Get(test.left+7, 1)
			assert.Equal(t, "x", ch)
			assert.Equal(t, outsideStyle, style)
		})
	}
}
