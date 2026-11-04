package gocui

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSearchRepeatedPrefix(b *testing.B) {
	for _, query := range []string{"a", strings.Repeat("a", 31) + "b", strings.Repeat("a", 127) + "b"} {
		b.Run(fmt.Sprintf("query=%d", len(query)), func(b *testing.B) {
			v := NewView("search", 0, 0, 80, 24, OutputNormal)
			v.SetContent(strings.Repeat(strings.Repeat("a", 200)+"\n", 2000))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				v.UpdateSearchResults(query, nil)
			}
		})
	}
}
