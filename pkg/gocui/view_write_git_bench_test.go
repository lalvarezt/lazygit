package gocui

import "testing"

func BenchmarkViewWriteGitColoredDiff(b *testing.B) {
	line := []byte("\x1b[32m+\x1b[m\x1b[32mline 050000 NAV-BIG-B performance content with wide 界 and é\x1b[m\n")
	b.ReportAllocs()
	b.SetBytes(int64(len(line) * 50000))
	for b.Loop() {
		v := NewView("diff", 0, 0, 160, 48, OutputNormal)
		for range 50000 {
			_, _ = v.Write(line)
		}
	}
}
