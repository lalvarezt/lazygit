package gocui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifyCellsWritten(t *testing.T) {
	for _, test := range []struct {
		columns, startColumn, width, wantRow, wantColumn int
	}{
		{80, 1, 80, 1, 81},
		{80, 81, 1, 2, 2},
		{80, 1, 160, 2, 81},
		{80, 80, 2, 2, 2},
		{80, 40, 100000, 1251, 40},
		{1, 1, 3, 3, 2},
		{0, 1, 1000, 1, 1},
		{80, 81, 0, 1, 81},
		{80, 1, -1, 1, 1},
	} {
		t.Run(fmt.Sprintf("%d/%d/%d", test.columns, test.startColumn, test.width), func(t *testing.T) {
			ei := escapeInterpreter{screenColMax: test.columns, screenCol: test.startColumn, screenRow: 1}
			ei.notifyCellsWritten(test.width)
			assert.Equal(t, test.wantRow, ei.screenRow)
			assert.Equal(t, test.wantColumn, ei.screenCol)
		})
	}
}

func BenchmarkNotifyCellsWritten(b *testing.B) {
	for _, width := range []int{1, 80, 4096, 100000} {
		b.Run(fmt.Sprintf("columns=%d", width), func(b *testing.B) {
			b.ReportAllocs()
			ei := escapeInterpreter{screenColMax: 80}
			for b.Loop() {
				ei.screenRow = 1
				ei.screenCol = 1
				ei.notifyCellsWritten(width)
			}
		})
	}
}
