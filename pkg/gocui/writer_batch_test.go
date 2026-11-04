package gocui

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteMatchesSmallChunks(t *testing.T) {
	for _, test := range []struct {
		name     string
		existing string
		input    string
	}{
		{name: "plain", input: "some printable ASCII text\nmore text\n"},
		{name: "long line", input: strings.Repeat("diff content ", 100)},
		{name: "Unicode boundaries", input: "before 界 é 👩‍💻 after\n"},
		{name: "combining mark after ASCII", input: "abcde\u0301 tail x\ufe0f rest\n"},
		{name: "invalid UTF-8", input: "before\xff\xfeafter\n"},
		{name: "combining mark after SGR final", input: "\x1b[32m\u0301visible\x1b[0m"},
		{name: "character set designation", input: "\x1b(BASCII text\n"},
		{name: "SGR", input: "\x1b[32mcolored text\x1b[0m plain text\n"},
		{name: "metadata", input: osc1717("1;a;1;;file.go") + "+added content\n"},
		{name: "orphaned metadata", input: osc1717("1;a;1;;first.go") + osc1717("1;a;2;;second.go") + "visible content\n"},
		{name: "hyperlink", input: "\x1b]8;;https://example.com\x07linked text\x1b]8;;\x07 plain text\n"},
		{name: "tabs", input: "before\tmiddle\ttail\n"},
		{name: "carriage return", input: "long old content\rshort\x1b[K\n"},
		{name: "CRLF", input: "first line\r\nsecond line\r\n"},
		{name: "cursor forward", input: "left side\x1b[5Cright side\n"},
		{name: "cursor down", input: "first line\x1b[3;1Hthird line\n"},
		{name: "soft wraps", input: strings.Repeat("a", 55) + "\x1b[5;1Hafter cursor advance\n"},
		{name: "overwrite", existing: "keep the trailing cells", input: "new text"},
	} {
		t.Run(test.name, func(t *testing.T) {
			batch := NewView("batch", 0, 0, 20, 10, OutputNormal)
			single := NewView("single", 0, 0, 20, 10, OutputNormal)
			if test.existing != "" {
				batch.writeString(test.existing)
				single.writeString(test.existing)
				batch.buf.seekWrite(0, 0)
				single.buf.seekWrite(0, 0)
			}

			batch.writeString(test.input)
			for _, chr := range stringToGraphemes(test.input) {
				single.writeString(chr)
			}

			assert.Equal(t, single.buf.lines, batch.buf.lines)
			assert.Equal(t, single.buf.wx, batch.buf.wx, "write column")
			assert.Equal(t, single.buf.wy, batch.buf.wy, "write row")
			assert.Equal(t, single.buf.pendingNewline, batch.buf.pendingNewline)
			assert.Equal(t, single.buf.ei.screenCol, batch.buf.ei.screenCol, "screen column")
			assert.Equal(t, single.buf.ei.screenRow, batch.buf.ei.screenRow, "screen row")
			assert.Equal(t, single.buf.ei.state, batch.buf.ei.state)
			assert.Equal(t, single.buf.ei.metadata.String(), batch.buf.ei.metadata.String())
			assert.Equal(t, single.buf.ei.metadataConsumed, batch.buf.ei.metadataConsumed)
			assert.Equal(t, single.buf.ei.hyperlink.String(), batch.buf.ei.hyperlink.String())
		})
	}
}
