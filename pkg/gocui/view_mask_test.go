package gocui

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawMaskedWideCharactersClearsUncoveredColumns(t *testing.T) {
	for _, test := range []struct {
		name    string
		content string
		mask    string
		wrap    bool
		want    []string
	}{
		{name: "CJK", content: "界", mask: "*", want: []string{"* ****", "      ", "      "}},
		{name: "emoji", content: "👩‍💻", mask: "*", want: []string{"* ****", "      ", "      "}},
		{name: "adjacent wide characters", content: "界界", mask: "*", want: []string{"* * **", "      ", "      "}},
		{name: "mixed widths", content: "a界b", mask: "*", want: []string{"** ***", "      ", "      "}},
		{name: "multiple lines", content: "界\n界", mask: "*", want: []string{"* ****", "* ****", "      "}},
		{name: "wrapped content", content: "界界界界", mask: "*", wrap: true, want: []string{"* * * ", "* ****", "      "}},
		{name: "Unicode mask", content: "界", mask: "•", want: []string{"• ••••", "      ", "      "}},
	} {
		t.Run(test.name, func(t *testing.T) {
			WithSimulationScreen(t, 8, 5)
			v := NewView("masked", 0, 0, 7, 4, OutputNormal)
			v.FgColor = ColorGreen
			v.BgColor = ColorBlue
			v.Mask = test.mask
			v.SetContent("abcdef\nabcdef\nabcdef")
			v.draw(true)

			v.SetContent(test.content)
			v.Wrap = test.wrap
			v.draw(true)

			for y, want := range test.want {
				var row strings.Builder
				for x := range 6 {
					ch, style, _ := Screen.Get(x+1, y+1)
					row.WriteString(ch)
					assert.Equal(t, getTcellStyle(oldStyle{fg: v.FgColor, bg: v.BgColor, outputMode: v.outMode}), style)
				}
				assert.Equal(t, want, row.String())
			}
		})
	}
}
