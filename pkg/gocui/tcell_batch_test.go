package gocui

import (
	"sync"
	"testing"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/stretchr/testify/assert"
)

func TestTcellLockCellsMatchesScreen(t *testing.T) {
	type renderedCell struct {
		text  string
		style tcell.Style
		width int
	}
	type putResult struct {
		rest  string
		width int
	}
	type cellWriter interface {
		Put(int, int, string, tcell.Style) (string, int)
		FillArea(int, int, int, int, rune, tcell.Style)
	}
	base := tcell.StyleDefault.Foreground(color.Green).Background(color.Blue)
	preserve := tcell.StyleDefault.Foreground(color.None).Background(color.None).Bold(true)
	var expectedCells [][]renderedCell
	var expectedResults []putResult
	for _, batched := range []bool{false, true} {
		t.Run(map[bool]string{false: "individual", true: "batched"}[batched], func(t *testing.T) {
			WithSimulationScreen(t, 8, 3)
			Screen.Fill('x', base)
			Screen.Show()

			var results []putResult
			write := func(writer cellWriter) {
				writer.FillArea(-1, 0, 4, 2, ' ', preserve)
				for _, op := range []struct {
					x, y int
					text string
				}{
					{0, 0, "érest"},
					{1, 0, "界rest"},
					{3, 0, "👩‍💻rest"},
					{7, 0, "界"},
					{0, 1, "\x1b]52;c;payload\aA"},
					{1, 1, ""},
					{2, 1, "\u0301a"},
					{-1, 1, "outside"},
					{8, 1, "outside"},
					{0, -1, "outside"},
					{0, 3, "outside"},
				} {
					rest, width := writer.Put(op.x, op.y, op.text, preserve)
					results = append(results, putResult{rest, width})
				}
			}
			if batched {
				func() {
					cells := Screen.LockCells()
					defer cells.Unlock()
					write(cells)
				}()
			} else {
				write(Screen)
			}

			for pass := range 2 {
				rendered := make([]renderedCell, 0, 24)
				for y := range 3 {
					for x := range 8 {
						text, style, width := Screen.Get(x, y)
						rendered = append(rendered, renderedCell{text, style, width})
					}
				}
				if batched {
					assert.Equal(t, expectedCells[pass], rendered)
					assert.Equal(t, expectedResults, results)
				} else {
					expectedCells = append(expectedCells, rendered)
					expectedResults = results
				}
				Screen.Show()
			}
		})
	}
}

func TestTcellLockCellsUnlocksOnPanic(t *testing.T) {
	WithSimulationScreen(t, 8, 3)
	assert.PanicsWithValue(t, "draw failed", func() {
		cells := Screen.LockCells()
		defer cells.Unlock()
		cells.Put(0, 0, "a", tcell.StyleDefault)
		panic("draw failed")
	})
	Screen.Put(1, 0, "b", tcell.StyleDefault)
	text, _, _ := Screen.Get(1, 0)
	assert.Equal(t, "b", text)
}

func TestTcellLockCellsConcurrentAccess(t *testing.T) {
	WithSimulationScreen(t, 8, 3)
	screen := Screen
	var workers sync.WaitGroup
	workers.Go(func() {
		for range 100 {
			screen.Put(7, 2, "x", tcell.StyleDefault)
			screen.Get(0, 0)
			screen.Show()
		}
	})
	for range 100 {
		func() {
			cells := screen.LockCells()
			defer cells.Unlock()
			cells.FillArea(0, 0, 8, 2, ' ', tcell.StyleDefault)
			for x := range 8 {
				cells.Put(x, 0, "a", tcell.StyleDefault)
			}
		}()
	}
	workers.Wait()
	text, _, _ := screen.Get(0, 0)
	assert.Equal(t, "a", text)
	text, _, _ = screen.Get(7, 2)
	assert.Equal(t, "x", text)
}

func TestTcellLockCellsAllocations(t *testing.T) {
	WithSimulationScreen(t, 8, 3)
	assert.Zero(t, testing.AllocsPerRun(100, func() {
		cells := Screen.LockCells()
		defer cells.Unlock()
		cells.Put(0, 0, "a", tcell.StyleDefault)
	}))
}
