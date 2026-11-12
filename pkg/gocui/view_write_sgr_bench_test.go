package gocui

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkViewWriteSGR(b *testing.B) {
	for _, test := range []struct {
		name, text string
	}{
		{"ascii", "+content"},
		{"unicode", "│中文🙂"},
		{"combining", "\u0301content"},
		{"unique", "│content"},
	} {
		for _, lines := range []int{1, 2000} {
			b.Run(fmt.Sprintf("%s/lines-%d", test.name, lines), func(b *testing.B) {
				var input strings.Builder
				for i := range lines {
					red, blue := 125, 200
					if test.name == "unique" {
						red = i % 256
						blue = i / 256
					}
					fmt.Fprintf(&input, "\x1b[38;2;%d;150;%dm%s\x1b[0m\n", red, blue, test.text)
				}
				text := input.String()
				b.ReportAllocs()
				b.SetBytes(int64(len(text)))
				for b.Loop() {
					v := NewView("diff", 0, 0, 101, 41, OutputTrue)
					v.SetContent(text)
				}
			})
		}
	}
}
