package gui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/gocui"
)

func BenchmarkCommitDescriptionInput(b *testing.B) {
	for _, size := range []int{256, 1024, 4096} {
		b.Run(fmt.Sprintf("bytes=%d", size), func(b *testing.B) {
			text := strings.Repeat("body text\n", size/10)
			keys := make([]gocui.Key, 0, len(text))
			for _, r := range text {
				key := gocui.NewKeyRune(r)
				if r == '\n' {
					key = gocui.NewKeyName(gocui.KeyEnter)
				}
				keys = append(keys, key)
			}
			gui := &Gui{}
			b.SetBytes(int64(len(text)))
			b.ReportAllocs()
			for b.Loop() {
				v := gocui.NewView("description", 0, 0, 80, 24, gocui.OutputNormal)
				v.TextArea.AutoWrap = true
				v.TextArea.AutoWrapWidth = 72
				for _, key := range keys {
					gui.commitDescriptionEditor(v, key)
				}
			}
		})
	}
}
