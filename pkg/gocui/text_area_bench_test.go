package gocui

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkPasteText(b *testing.B) {
	for _, size := range []int{1000, 10000, 100000} {
		b.Run(fmt.Sprintf("bytes=%d", size), func(b *testing.B) {
			pasted := strings.Repeat("some pasted text\n", size/17)
			b.SetBytes(int64(len(pasted)))
			b.ReportAllocs()
			for b.Loop() {
				area := &TextArea{content: "prefix suffix", cursor: len("prefix "), AutoWrapWidth: 72, AutoWrap: true}
				area.TypeString(pasted)
			}
		})
	}
}
