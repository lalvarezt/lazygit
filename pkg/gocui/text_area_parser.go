package gocui

import (
	"strings"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

// Checkpoints have no breakable whitespace in the pending word. Its cells
// remain an unchanged prefix with stable coordinates in the completed layout.
type textAreaCheckpoint struct {
	nextByte       int
	graphemeState  int
	outputCells    int
	pendingCells   int
	lineWidth      int
	row            int
	footnotePrefix string
	footnoteFailed bool
	trailer        trailerMatcher
}

type textAreaParser struct {
	content           string
	width             int
	trailerBlockStart int
	nextByte          int
	graphemeState     int
	row               int
	lineWidth         int
	lastWhitespace    int
	footnote          footNoteMatcher
	trailer           trailerMatcher
	cells             []TextAreaCell // completed lines followed by the current visual line
	pendingStart      int
	positioned        int // prefix of pending cells with stable coordinates
	softLineBreaks    []int
	checkpointing     bool
	checkpoints       []textAreaCheckpoint
	ownedText         bool
}

func contentToCells(content string, width int) ([]TextAreaCell, []int) {
	estimatedBreaks := 0
	if width > 0 {
		estimatedBreaks = len(content) / width
	}
	p := textAreaParser{
		content:           content,
		width:             width,
		trailerBlockStart: startOfTrailerBlock(content),
		graphemeState:     -1,
		lastWhitespace:    -1,
		cells:             make([]TextAreaCell, 0, len(content)+estimatedBreaks),
		softLineBreaks:    make([]int, 0, estimatedBreaks),
	}
	p.parse()
	return p.cells, p.softLineBreaks
}

func (p *textAreaParser) parse() {
	for p.nextByte < len(p.content) {
		chr, _, width, state := uniseg.FirstGraphemeClusterInString(p.content[p.nextByte:], p.graphemeState)
		c := TextAreaCell{char: chr, width: width, contentIndex: p.nextByte}
		p.nextByte += len(chr)
		p.graphemeState = state
		if p.ownedText {
			if len(chr) == 1 && chr[0] < utf8.RuneSelf {
				c.char = string([]byte{chr[0]})
			} else {
				c.char = strings.Clone(chr)
			}
		}
		p.cells = append(p.cells, c)
		p.consume(c)
		if p.checkpointing && p.lastWhitespace == -1 && p.nextByte-p.checkpoints[len(p.checkpoints)-1].nextByte >= 128 {
			p.checkpoint()
		}
	}
	p.emit(len(p.cells) - p.pendingStart)
}

func (p *textAreaParser) consume(c TextAreaCell) {
	if c.char == "\n" {
		p.emit(len(p.cells) - p.pendingStart)
		p.row++
		p.lastWhitespace = -1
		p.lineWidth = 0
		p.footnote.reset()
		p.trailer.reset()
		p.checkpoint()
		return
	}

	p.lineWidth += c.width
	inTrailerBlock := c.contentIndex >= p.trailerBlockStart
	wrapped := false
	if c.char == " " && !p.footnote.isFootNote() &&
		!(inTrailerBlock && p.trailer.isTrailer(p.content[p.nextByte:])) {
		p.lastWhitespace = len(p.cells) - p.pendingStart
	} else if p.width > 0 && p.lineWidth > p.width && p.lastWhitespace >= 0 {
		p.emit(p.lastWhitespace)
		contentIndex := p.cells[p.pendingStart].contentIndex
		p.row++
		p.cells = append(p.cells, TextAreaCell{})
		copy(p.cells[p.pendingStart+1:], p.cells[p.pendingStart:len(p.cells)-1])
		p.cells[p.pendingStart] = TextAreaCell{char: "\n", width: 1, contentIndex: contentIndex, x: 0, y: p.row}
		p.pendingStart++
		if p.softLineBreaks != nil {
			p.softLineBreaks = append(p.softLineBreaks, contentIndex)
		}
		p.lastWhitespace = -1
		p.lineWidth = 0
		for _, c := range p.cells[p.pendingStart:] {
			p.lineWidth += c.width
		}
		p.footnote.reset()
		p.trailer.reset()
		wrapped = true
	}

	p.footnote.addCharacter(c.char)
	if inTrailerBlock {
		p.trailer.addCharacter(c.char)
	}
	if wrapped {
		p.checkpoint()
	}
}

func (p *textAreaParser) emit(count int) {
	start := p.pendingStart + min(p.positioned, count)
	x := 0
	if start > p.pendingStart {
		last := p.cells[start-1]
		x = last.x + last.width
	}
	for i := start; i < p.pendingStart+count; i++ {
		p.cells[i].x = x
		p.cells[i].y = p.row
		x += p.cells[i].width
	}
	p.pendingStart += count
	p.positioned = max(0, p.positioned-count)
}

func (p *textAreaParser) checkpoint() {
	if !p.checkpointing {
		return
	}
	// URL-style trailers can match through whitespace by looking ahead into
	// text that a later edit may change. Dashed keys match on the key alone.
	if p.trailer.didMatch && !p.trailer.keyContainsDash {
		return
	}
	p.checkpoints = append(p.checkpoints, textAreaCheckpoint{
		nextByte:       p.nextByte,
		graphemeState:  p.graphemeState,
		outputCells:    p.pendingStart,
		pendingCells:   len(p.cells) - p.pendingStart,
		lineWidth:      p.lineWidth,
		row:            p.row,
		footnotePrefix: p.footnote.lineStr.String(),
		footnoteFailed: p.footnote.didFailToMatch,
		trailer:        p.trailer,
	})
}
