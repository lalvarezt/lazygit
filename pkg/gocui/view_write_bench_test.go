package gocui

import (
	"strings"
	"testing"
)

func BenchmarkViewWrite(b *testing.B) {
	for _, test := range []struct {
		name, line string
	}{
		{"plain", " a line of diff content that fits within the view\n"},
		{"colored", "\x1b[32m+added content\x1b[0m\tmore content\x1b[K\n"},
		{"unicode", " a line with 界 and é and 👩‍💻\n"},
		{"metadata", osc1717("1;a;1;;file.go") + "+added content\n"},
	} {
		b.Run(test.name, func(b *testing.B) {
			line := []byte(test.line)
			b.ReportAllocs()
			b.SetBytes(int64(len(line) * 2000))
			for b.Loop() {
				v := NewView("diff", 0, 0, 100, 40, OutputNormal)
				for range 2000 {
					_, _ = v.Write(line)
				}
			}
		})
	}
}

func BenchmarkViewWriteLongLine(b *testing.B) {
	line := []byte(strings.Repeat("diff content ", 1000) + "\n")
	b.ReportAllocs()
	b.SetBytes(int64(len(line)))
	for b.Loop() {
		v := NewView("diff", 0, 0, 100, 40, OutputNormal)
		_, _ = v.Write(line)
	}
}
