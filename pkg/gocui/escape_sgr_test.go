package gocui

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompleteSGRMatchesStreamingInterpreter(t *testing.T) {
	sequences := []string{
		"\x1b[m", "\x1b[0m", "\x1b[1;31m", "\x1b[22;39;49m",
		"\x1b[38;2;125;150;200;48;2;40;50;60m", "\x1b[38;5;123m",
		"\x1b[90;107m", "\x1b[1;;31m", "\x1b[;31m", "\x1b[38;2;1m",
		"\x1b[999m", "\x1b[999999999999999999999999999m",
	}
	for _, mode := range []OutputMode{OutputNormal, Output256, OutputTrue} {
		fast, slow := newEscapeInterpreter(mode), newEscapeInterpreter(mode)
		for _, sequence := range sequences {
			assert.Equal(t, len(sequence), fast.parseSGR([]byte(sequence)), sequence)
			for i := range sequence {
				_, err := slow.parseOne([]byte{sequence[i]})
				assert.NoError(t, err)
			}
			assert.Equal(t, slow, fast, sequence)
		}
	}
}

func TestCompleteSGRLeavesStreamingCasesUntouched(t *testing.T) {
	for _, sequence := range []string{
		"", "text", "\x1b[", "\x1b[31", "\x1b[2K", "\x1b[?25m",
		"\x1b[38:2:1:2:3m", "\x1b[31ḿ", "\x1b[" + strings.Repeat("1;", 20) + "1m",
		"\x1b[" + strings.Repeat("1", 256) + "m",
	} {
		ei := newEscapeInterpreter(OutputTrue)
		assert.Zero(t, ei.parseSGR([]byte(sequence)), sequence)
		assert.Equal(t, newEscapeInterpreter(OutputTrue), ei)
	}
}

func TestCompleteSGRUnicodeBoundary(t *testing.T) {
	for _, suffix := range []struct {
		name, text string
		complete   bool
	}{
		{"box", "│", true},
		{"wide", "中文", true},
		{"emoji", "🙂", true},
		{"flag", "🇫🇷", true},
		{"prepend", "\u0600x", true},
		{"separator", "\u2028", true},
		{"invalid", "\xff", true},
		{"invalid-continuation", "\xe2x", true},
		{"combining", "\u0301", false},
		{"before-combining-block", "\u02ff", true},
		{"combining-block-start", "\u0300", false},
		{"combining-block-end", "\u036f", false},
		{"after-combining-block", "\u0370", true},
		{"variation-selector", "\ufe0f", false},
		{"joiner", "\u200d🙂", false},
		{"spacing-mark", "\u0903", false},
		{"enclosing-mark", "\u20dd", false},
		{"supplementary-mark", "\U0001d165", false},
		{"non-joiner", "\u200c", false},
		{"emoji-modifier", "\U0001f3fb", false},
		{"emoji-tag", "\U000e0020", false},
		{"partial-box", "\xe2", false},
		{"partial-emoji", "\xf0\x9f", false},
		{"partial-combining", "\xcc", false},
	} {
		t.Run(suffix.name, func(t *testing.T) {
			for _, mode := range []OutputMode{OutputNormal, Output256, OutputTrue} {
				fast, slow := newEscapeInterpreter(mode), newEscapeInterpreter(mode)
				for _, sequence := range []string{
					"\x1b[1;38;2;125;150;200;48;2;40;50;60m",
					"\x1b[22;39m", "\x1b[38;5;123m", "\x1b[1;;31m", "\x1b[m",
				} {
					if !suffix.complete {
						assert.Zero(t, fast.parseSGR([]byte(sequence+suffix.text)))
						assert.Equal(t, slow, fast)
						continue
					}
					assert.Equal(t, len(sequence), fast.parseSGR([]byte(sequence+suffix.text)))
					for i := range sequence {
						_, err := slow.parseOne([]byte{sequence[i]})
						assert.NoError(t, err)
					}
					assert.Equal(t, slow, fast)
				}
			}
		})
	}
}

func TestSGRWriteChunkBoundaries(t *testing.T) {
	input := "\x1b[1;38;2;125;150;200mtext\x1b[m\x1b[48;5;123mbackground\x1b[0m\n"
	want := NewView("diff", 0, 0, 80, 40, OutputTrue)
	want.WriteString(input)
	for split := range len(input) {
		got := NewView("diff", 0, 0, 80, 40, OutputTrue)
		got.WriteString(input[:split])
		got.WriteString(input[split:])
		assert.Equal(t, want.buf.lines, got.buf.lines, split)
	}
}
