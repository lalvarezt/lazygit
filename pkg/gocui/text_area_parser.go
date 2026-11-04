package gocui

import "github.com/rivo/uniseg"

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
		return
	}

	p.lineWidth += c.width
	inTrailerBlock := c.contentIndex >= p.trailerBlockStart
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
	}

	p.footnote.addCharacter(c.char)
	if inTrailerBlock {
		p.trailer.addCharacter(c.char)
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
