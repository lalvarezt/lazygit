package gocui

import (
	"strings"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

// Checkpoints follow line transitions, when no breakable whitespace remains in
// the pending word. Its cells stay contiguous after outputCells in the completed
// layout, so restarting only copies that word.
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
	pending           []TextAreaCell // current visual line, including any incomplete word
	result            []TextAreaCell // completed lines
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
		result:            make([]TextAreaCell, 0, len(content)+estimatedBreaks),
		softLineBreaks:    make([]int, 0, estimatedBreaks),
	}
	p.parse()
	return p.result, p.softLineBreaks
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
		p.pending = append(p.pending, c)
		p.consume(c)
	}
	p.emit(len(p.pending))
}

func (p *textAreaParser) consume(c TextAreaCell) {
	if c.char == "\n" {
		p.emit(len(p.pending))
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
		p.lastWhitespace = len(p.pending)
	} else if p.width > 0 && p.lineWidth > p.width && p.lastWhitespace >= 0 {
		p.emit(p.lastWhitespace)
		contentIndex := p.pending[0].contentIndex
		p.row++
		p.result = append(p.result, TextAreaCell{char: "\n", width: 1, contentIndex: contentIndex, x: 0, y: p.row})
		if p.softLineBreaks != nil {
			p.softLineBreaks = append(p.softLineBreaks, contentIndex)
		}
		p.lastWhitespace = -1
		p.lineWidth = 0
		for _, c := range p.pending {
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
	x := 0
	for i := range count {
		p.pending[i].x = x
		p.pending[i].y = p.row
		x += p.pending[i].width
	}
	p.result = append(p.result, p.pending[:count]...)
	remaining := copy(p.pending, p.pending[count:])
	clear(p.pending[remaining:])
	p.pending = p.pending[:remaining]
}

func (p *textAreaParser) checkpoint() {
	if !p.checkpointing {
		return
	}
	p.checkpoints = append(p.checkpoints, textAreaCheckpoint{
		nextByte:       p.nextByte,
		graphemeState:  p.graphemeState,
		outputCells:    len(p.result),
		pendingCells:   len(p.pending),
		lineWidth:      p.lineWidth,
		row:            p.row,
		footnotePrefix: p.footnote.lineStr.String(),
		footnoteFailed: p.footnote.didFailToMatch,
		trailer:        p.trailer,
	})
}
