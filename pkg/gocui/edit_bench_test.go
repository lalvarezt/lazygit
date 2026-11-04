package gocui

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkEditorInput(b *testing.B) {
	benchmarkEditorInput(b, "body text\n")
}

func BenchmarkEditorParagraphInput(b *testing.B) {
	benchmarkEditorInput(b, "body text ")
}

func BenchmarkEditorUnbrokenInput(b *testing.B) {
	benchmarkEditorInput(b, "a")
}

func BenchmarkEditorEditUnbrokenText(b *testing.B) {
	for _, size := range []int{4096, 100000, 1000000} {
		b.Run(fmt.Sprintf("bytes=%d", size), func(b *testing.B) {
			v := NewView("editor", 0, 0, 80, 24, OutputNormal)
			v.TextArea.AutoWrap = true
			v.TextArea.AutoWrapWidth = 72
			v.TextArea.TypeString(strings.Repeat("a", size))
			v.RenderTextArea()
			key := NewKeyRune('b')
			backspace := NewKeyName(KeyBackspace)
			b.ReportAllocs()
			for b.Loop() {
				SimpleEditor(v, key)
				SimpleEditor(v, backspace)
			}
		})
	}
}

func benchmarkEditorInput(b *testing.B, body string) {
	b.Helper()
	for _, size := range []int{256, 1024, 4096} {
		b.Run(fmt.Sprintf("bytes=%d", size), func(b *testing.B) {
			text := strings.Repeat(body, size/len(body))
			keys := make([]Key, 0, len(text))
			for _, r := range text {
				key := NewKeyRune(r)
				if r == '\n' {
					key = NewKeyName(KeyEnter)
				}
				keys = append(keys, key)
			}
			b.SetBytes(int64(len(text)))
			b.ReportAllocs()
			for b.Loop() {
				v := NewView("editor", 0, 0, 80, 24, OutputNormal)
				v.TextArea.AutoWrap = true
				v.TextArea.AutoWrapWidth = 72
				for _, key := range keys {
					SimpleEditor(v, key)
				}
			}
		})
	}
}
