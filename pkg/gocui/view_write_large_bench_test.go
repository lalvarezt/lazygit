package gocui

import "testing"

func BenchmarkViewWriteLargeDiff(b *testing.B) {
	for _, fixture := range []struct {
		name, line string
	}{
		{"colored", "\x1b[32m+added content\x1b[0m\tmore content\x1b[K\n"},
		{"metadata", osc1717("1;a;1;;file.go") + "+added content\n"},
	} {
		b.Run(fixture.name, func(b *testing.B) {
			line := []byte(fixture.line)
			b.SetBytes(int64(len(line) * 50000))
			b.ReportAllocs()
			for b.Loop() {
				v := NewView("diff", 0, 0, 160, 48, OutputNormal)
				for range 50000 {
					_, _ = v.Write(line)
				}
			}
		})
	}
}
