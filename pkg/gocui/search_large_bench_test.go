package gocui

import (
	"strings"
	"testing"
)

func BenchmarkSearchLargeDiff(b *testing.B) {
	for _, query := range []string{"a", "performance-match", "missing", "MATCH"} {
		b.Run("lines=50000/query="+query, func(b *testing.B) {
			v := NewView("diff", 0, 0, 160, 48, OutputNormal)
			v.SetContent(strings.Repeat("+performance-match diff content with 界 and é and ordinary text\n", 50000))
			v.UpdateSearchResults(query, nil)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				v.UpdateSearchResults(query, nil)
			}
		})
	}
}
