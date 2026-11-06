package gocui

import (
	"strings"
	"testing"
)

func BenchmarkSearchLargeDiff(b *testing.B) {
	for _, fixture := range []struct {
		name, line string
	}{
		{"ordinary", "+performance-match diff content with 界 and é and ordinary text\n"},
		{"mixed-case", "+Performance-Match DiffContent With 界 and é and OrdinaryText\n"},
	} {
		for _, query := range []string{"a", "performance-match", "missing", "MATCH"} {
			b.Run("lines=50000/text="+fixture.name+"/query="+query, func(b *testing.B) {
				v := NewView("diff", 0, 0, 160, 48, OutputNormal)
				v.SetContent(strings.Repeat(fixture.line, 50000))
				v.UpdateSearchResults(query, nil)
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					v.UpdateSearchResults(query, nil)
				}
			})
		}
	}
}
