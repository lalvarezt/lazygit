package gocui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffLineContentsWithCRLF(t *testing.T) {
	for _, test := range []struct {
		name   string
		chunks []string
		want   []string
	}{
		{
			name:   "one line",
			chunks: []string{"first\r\n"},
			/* EXPECTED:
			want: []string{"first"},
			ACTUAL: */
			want: []string{"first\r"},
		},
		{
			name:   "lines in one write",
			chunks: []string{"first\r\nsecond\r\n"},
			/* EXPECTED:
			want: []string{"first", "second"},
			ACTUAL: */
			want: []string{"first", "second\r"},
		},
		{
			name:   "lines in separate writes",
			chunks: []string{"first\r\n", "second\r\n"},
			/* EXPECTED:
			want: []string{"first", "second"},
			ACTUAL: */
			want: []string{"first\r", "second\r"},
		},
		{
			name:   "tabs",
			chunks: []string{"first\tvalue\r\n"},
			/* EXPECTED:
			want: []string{"first\tvalue"},
			ACTUAL: */
			want: []string{"first\tvalue\r"},
		},
		{
			name:   "bare carriage return before CRLF",
			chunks: []string{"before\rafter\r\n"},
			/* EXPECTED:
			want: []string{"before\rafter"},
			ACTUAL: */
			want: []string{"before\rafter\r"},
		},
		{
			name:   "bare carriage return before LF",
			chunks: []string{"before\rafter\n"},
			want:   []string{"before\rafter"},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v := NewView("diff", 0, 0, 40, 10, OutputNormal)
			for _, chunk := range test.chunks {
				v.writeString(chunk)
			}
			contents := v.DiffLineContents()
			text := make([]string, len(contents))
			for i, line := range contents {
				text[i] = line.Text
			}
			assert.Equal(t, test.want, text)
		})
	}
}
