package gocui

import (
	"regexp"
	"slices"
	"strings"

	"github.com/rivo/uniseg"
)

const (
	WHITESPACES     = " \t"
	WORD_SEPARATORS = "*?_+-.[]~=/&;!#$%^(){}<>"
)

type TextAreaCell struct {
	char         string // string because it could be a multi-rune grapheme cluster
	width        int
	x, y         int // cell coordinates
	contentIndex int // byte index into the original content
}

// returns the cursor x,y position after this cell
func (c *TextAreaCell) nextCursorXY() (int, int) {
	if c.char == "\n" {
		return 0, c.y + 1
	}
	return c.x + c.width, c.y
}

type textAreaLayout struct {
	checkpoints    []textAreaCheckpoint
	trailerStart   int
	firstControl   int // first tab, carriage return, or escape byte; -1 if absent
	content        string
	width          int
	ready          bool
	revision       uint64
	unchangedRows  int
	unchangedCells int
}

type TextArea struct {
	content         string
	layout          textAreaLayout
	hasChanges      bool
	changedFromByte int
	cells           []TextAreaCell
	cursor          int // position in content, as an index into the byte array
	overwrite       bool
	clipboard       string
	AutoWrap        bool
	AutoWrapWidth   int
}

// Returns the indices in content where soft line breaks occur due to auto-wrapping to the given width.
func AutoWrapContent(content string, autoWrapWidth int) []int {
	_, softLineBreakIndices := contentToCells(content, autoWrapWidth)
	return softLineBreakIndices
}

// startOfTrailerBlock returns the byte index into content at which the trailer
// block begins, i.e. the start of the last paragraph (the run of lines at the
// end of the message that is separated from the body by a blank line). Trailers
// are only looked for from this index onwards, so that a trailer-looking line in
// the middle of the message body isn't mistaken for a trailer. Trailing blank
// lines are ignored, and a message that consists of a single paragraph is
// treated as its own trailer block.
func startOfTrailerBlock(content string) int {
	return lastBlankLineEnd(strings.TrimRight(content, "\n"))
}

func lastBlankLineEnd(content string) int {
	if blankLine := strings.LastIndex(content, "\n\n"); blankLine >= 0 {
		return blankLine + len("\n\n")
	}
	return 0
}

func (self *TextArea) scanContent() (trailerStart, firstControl int) {
	if !self.layout.ready || !self.hasChanges {
		return startOfTrailerBlock(self.content), strings.IndexAny(self.content, "\t\r\x1b")
	}

	end := len(strings.TrimRight(self.content, "\n"))
	start := 0
	if self.layout.trailerStart <= self.changedFromByte && self.layout.trailerStart <= end {
		trailerStart = self.layout.trailerStart
		// A separator can overlap the edit boundary. Trailing blank lines
		// can also become interior text when non-newline content is appended.
		unchanged := min(self.changedFromByte, len(strings.TrimRight(self.layout.content, "\n")))
		start = max(0, trailerStart-1, unchanged-1)
		start = min(start, end)
	}
	if next := lastBlankLineEnd(self.content[start:end]); next > 0 {
		trailerStart = start + next
	}

	firstControl = self.layout.firstControl
	if firstControl < 0 || firstControl >= self.changedFromByte {
		firstControl = strings.IndexAny(self.content[self.changedFromByte:], "\t\r\x1b")
		if firstControl >= 0 {
			firstControl += self.changedFromByte
		}
	}
	return trailerStart, firstControl
}

var footNoteRe = regexp.MustCompile(`^\[\d+\]:\s*$`)

type footNoteMatcher struct {
	lineStr        strings.Builder
	didFailToMatch bool
}

func (self *footNoteMatcher) addCharacter(chr string) {
	if self.didFailToMatch {
		// don't bother tracking the rune if we know it can't possibly match any more
		return
	}

	if self.lineStr.Len() == 0 && chr != "[" {
		// fail early if the first rune of a line isn't a '['; this is mainly to avoid a (possibly
		// expensive) regex match
		self.didFailToMatch = true
		return
	}

	self.lineStr.WriteString(chr)
}

func (self *footNoteMatcher) isFootNote() bool {
	if self.didFailToMatch {
		return false
	}

	if footNoteRe.MatchString(self.lineStr.String()) {
		// it's a footnote, so treat spaces as non-breaking. It's important not to reset the matcher
		// here, because there could be multiple spaces after a footnote.
		return true
	}

	// no need to check again for this line
	self.didFailToMatch = true
	return false
}

func (self *footNoteMatcher) reset() {
	self.lineStr.Reset()
	self.didFailToMatch = false
}

type trailerMatcher struct {
	didFailToMatch   bool
	didMatch         bool
	keyContainsDash  bool
	keyEndsWithColon bool
}

func (self *trailerMatcher) addCharacter(chr string) {
	if self.didFailToMatch || self.didMatch {
		return
	}

	if len(chr) != 1 {
		// Trailers are all ASCII, so if we get a non-ASCII UTF-8 character (or even a multi-rune
		// grapheme cluster), we can fail early.
		self.didFailToMatch = true
		return
	}

	self.keyContainsDash = self.keyContainsDash || chr == "-"
	self.keyEndsWithColon = chr == ":"
}

func (self *trailerMatcher) isTrailer(remainingContent string) bool {
	if self.didFailToMatch {
		return false
	}

	if self.didMatch {
		return true
	}

	remainingContent = strings.TrimLeft(remainingContent, WHITESPACES)
	if self.keyEndsWithColon && (self.keyContainsDash ||
		strings.HasPrefix(remainingContent, "http://") || strings.HasPrefix(remainingContent, "https://")) {
		self.didMatch = true
		return true
	}

	self.didFailToMatch = true
	return false
}

func (self *trailerMatcher) reset() {
	self.didFailToMatch = false
	self.didMatch = false
	self.keyContainsDash = false
	self.keyEndsWithColon = false
}

func (self *TextArea) checkpointBeforeEdit(width, trailerStart int) (textAreaCheckpoint, int) {
	if !self.layout.ready || self.layout.width != width || !self.hasChanges || len(self.layout.checkpoints) == 0 {
		return textAreaCheckpoint{graphemeState: -1}, -1
	}
	dirty := self.changedFromByte
	if self.layout.trailerStart != trailerStart {
		dirty = min(dirty, self.layout.trailerStart, trailerStart)
	}
	cellIndex, _ := slices.BinarySearchFunc(self.cells, dirty, func(c TextAreaCell, index int) int { return c.contentIndex - index })
	// Grapheme parsing stores lookahead. Keep it inside unchanged text, even
	// when an edit merges clusters. Soft breaks can share a source position.
	safeByte, lastIndex, distinct := 0, dirty, 0
	for i := cellIndex - 1; i >= 0 && distinct < 2; i-- {
		if self.cells[i].contentIndex == lastIndex {
			continue
		}
		safeByte, lastIndex = self.cells[i].contentIndex, self.cells[i].contentIndex
		distinct++
	}
	if distinct < 2 {
		safeByte = 0
	}
	index, found := slices.BinarySearchFunc(self.layout.checkpoints, safeByte, func(c textAreaCheckpoint, index int) int { return c.nextByte - index })
	if !found {
		index--
	}
	index = max(index, 0)
	return self.layout.checkpoints[index], index
}

func (self *TextArea) updateCells() {
	width := self.AutoWrapWidth
	if !self.AutoWrap {
		width = -1
	}
	if self.layout.ready && self.layout.width == width && self.layout.content == self.content {
		self.hasChanges = false
		return
	}
	trailerStart, firstControl := self.scanContent()
	checkpoint, index := self.checkpointBeforeEdit(width, trailerStart)
	oldCells, oldCheckpoints := self.cells, self.layout.checkpoints
	checkpoints := oldCheckpoints[:0]
	if index >= 0 {
		checkpoints = oldCheckpoints[:index+1]
	} else {
		checkpoints = append(checkpoints, checkpoint)
	}
	remainingBytes := len(self.content) - checkpoint.nextByte
	remainingCells := remainingBytes
	if width > 0 {
		remainingCells += (checkpoint.lineWidth + remainingBytes) / width
	}
	p := textAreaParser{
		content:           self.content,
		width:             width,
		trailerBlockStart: trailerStart,
		nextByte:          checkpoint.nextByte,
		graphemeState:     checkpoint.graphemeState,
		row:               checkpoint.row,
		lineWidth:         checkpoint.lineWidth,
		lastWhitespace:    -1,
		trailer:           checkpoint.trailer,
		cells:             slices.Grow(oldCells[:checkpoint.outputCells+checkpoint.pendingCells], remainingCells),
		pendingStart:      checkpoint.outputCells,
		positioned:        checkpoint.pendingCells,
		checkpointing:     true,
		checkpoints:       checkpoints,
		ownedText:         true,
	}
	p.footnote.didFailToMatch = checkpoint.footnoteFailed
	p.footnote.lineStr.WriteString(checkpoint.footnotePrefix)
	p.parse()
	self.cells = p.cells
	if len(oldCells) > len(self.cells) {
		clear(oldCells[len(self.cells):])
	}
	if len(oldCheckpoints) > len(p.checkpoints) {
		clear(oldCheckpoints[len(p.checkpoints):])
	}
	self.layout.checkpoints = p.checkpoints
	self.layout.unchangedRows = checkpoint.row
	self.layout.unchangedCells = checkpoint.outputCells + checkpoint.pendingCells
	self.layout.content, self.layout.width, self.layout.ready = self.content, width, true
	self.layout.trailerStart = trailerStart
	self.layout.firstControl = firstControl
	self.hasChanges = false
	self.layout.revision++
}

func (self *TextArea) replaceContent(start, end int, text string) {
	if !self.hasChanges {
		self.changedFromByte = start
	} else {
		self.changedFromByte = min(self.changedFromByte, start)
	}
	self.hasChanges = true
	self.content = self.content[:start] + text + self.content[end:]
}

func (self *TextArea) typeCharacter(ch string) {
	widthToDelete := 0
	if self.overwrite && !self.atEnd() {
		s, _, _, _ := uniseg.FirstGraphemeClusterInString(self.content[self.cursor:], -1)
		widthToDelete = len(s)
	}

	self.replaceContent(self.cursor, self.cursor+widthToDelete, ch)
	self.cursor += len(ch)
}

func (self *TextArea) TypeCharacter(ch string) {
	self.typeCharacter(ch)
	self.updateCells()
}

func (self *TextArea) BackSpaceChar() {
	if self.cursor == 0 {
		return
	}

	cellCursor := self.contentCursorToCellCursor(self.cursor)
	widthToDelete := len(self.cells[cellCursor-1].char)

	oldCursor := self.cursor
	self.cursor -= widthToDelete
	self.replaceContent(self.cursor, oldCursor, "")

	self.updateCells()
}

func (self *TextArea) DeleteChar() {
	if self.atEnd() {
		return
	}

	s, _, _, _ := uniseg.FirstGraphemeClusterInString(self.content[self.cursor:], -1)
	widthToDelete := len(s)
	self.replaceContent(self.cursor, self.cursor+widthToDelete, "")
	self.updateCells()
}

func (self *TextArea) MoveCursorLeft() {
	if self.cursor == 0 {
		return
	}

	cellCursor := self.contentCursorToCellCursor(self.cursor)
	self.cursor -= len(self.cells[cellCursor-1].char)
}

func (self *TextArea) MoveCursorRight() {
	if self.cursor == len(self.content) {
		return
	}

	s, _, _, _ := uniseg.FirstGraphemeClusterInString(self.content[self.cursor:], -1)
	self.cursor += len(s)
}

func (self *TextArea) newCursorForMoveLeftWord() int {
	if self.cursor == 0 {
		return 0
	}
	if self.atLineStart() {
		return self.cursor - 1
	}

	cellCursor := self.contentCursorToCellCursor(self.cursor)
	for cellCursor > 0 && (self.isSoftLineBreak(cellCursor-1) || strings.Contains(WHITESPACES, self.cells[cellCursor-1].char)) {
		cellCursor--
	}
	separators := false
	for cellCursor > 0 && strings.Contains(WORD_SEPARATORS, self.cells[cellCursor-1].char) {
		cellCursor--
		separators = true
	}
	if !separators {
		for cellCursor > 0 && self.cells[cellCursor-1].char != "\n" && !strings.Contains(WHITESPACES+WORD_SEPARATORS, self.cells[cellCursor-1].char) {
			cellCursor--
		}
	}

	return self.cellCursorToContentCursor(cellCursor)
}

func (self *TextArea) MoveLeftWord() {
	self.cursor = self.newCursorForMoveLeftWord()
}

func (self *TextArea) newCursorForMoveRightWord() int {
	if self.atEnd() {
		return self.cursor
	}
	if self.atLineEnd() {
		return self.cursor + 1
	}

	cellCursor := self.contentCursorToCellCursor(self.cursor)
	for cellCursor < len(self.cells) && (self.isSoftLineBreak(cellCursor) || strings.Contains(WHITESPACES, self.cells[cellCursor].char)) {
		cellCursor++
	}
	separators := false
	for cellCursor < len(self.cells) && strings.Contains(WORD_SEPARATORS, self.cells[cellCursor].char) {
		cellCursor++
		separators = true
	}
	if !separators {
		for cellCursor < len(self.cells) && self.cells[cellCursor].char != "\n" && !strings.Contains(WHITESPACES+WORD_SEPARATORS, self.cells[cellCursor].char) {
			cellCursor++
		}
	}

	return self.cellCursorToContentCursor(cellCursor)
}

func (self *TextArea) MoveRightWord() {
	self.cursor = self.newCursorForMoveRightWord()
}

func (self *TextArea) MoveCursorUp() {
	x, y := self.GetCursorXY()
	self.SetCursor2D(x, y-1)
}

func (self *TextArea) MoveCursorDown() {
	x, y := self.GetCursorXY()
	self.SetCursor2D(x, y+1)
}

func (self *TextArea) GetContent() string {
	var b strings.Builder
	for _, c := range self.cells {
		b.WriteString(c.char)
	}
	return b.String()
}

func (self *TextArea) GetUnwrappedContent() string {
	return self.content
}

func (self *TextArea) ToggleOverwrite() {
	self.overwrite = !self.overwrite
}

func (self *TextArea) atEnd() bool {
	return self.cursor == len(self.content)
}

func (self *TextArea) DeleteToStartOfLine() {
	// copying vim's logic: if you're at the start of the line, you delete the newline
	// character and go to the end of the previous line
	if self.atLineStart() {
		if self.cursor == 0 {
			return
		}

		self.replaceContent(self.cursor-1, self.cursor, "")
		self.cursor--
		self.updateCells()
		return
	}

	// otherwise, if we're at a soft line start, skip left past the soft line
	// break, so we'll end up deleting the previous line. This seems like the
	// only reasonable behavior in this case, as you can't delete just the soft
	// line break.
	if self.atSoftLineStart() {
		self.cursor--
	}

	// otherwise, you delete everything up to the start of the current line, without
	// deleting the newline character
	newlineIndex := self.closestNewlineOnLeft()
	self.clipboard = self.content[newlineIndex+1 : self.cursor]
	self.replaceContent(newlineIndex+1, self.cursor, "")
	self.updateCells()
	self.cursor = newlineIndex + 1
}

func (self *TextArea) DeleteToEndOfLine() {
	if self.atEnd() {
		return
	}

	// if we're at the end of the line, delete just the newline character
	if self.atLineEnd() {
		self.replaceContent(self.cursor, self.cursor+1, "")
		self.updateCells()
		return
	}

	// otherwise, if we're at a soft line end, skip right past the soft line
	// break, so we'll end up deleting the next line. This seems like the
	// only reasonable behavior in this case, as you can't delete just the soft
	// line break.
	if self.atSoftLineEnd() {
		self.cursor++
	}

	lineEndIndex := self.closestNewlineOnRight()
	self.clipboard = self.content[self.cursor:lineEndIndex]
	self.replaceContent(self.cursor, lineEndIndex, "")
	self.updateCells()
}

func (self *TextArea) GoToStartOfLine() {
	if self.atSoftLineStart() {
		return
	}

	newlineIndex := self.closestNewlineOnLeft()
	self.cursor = newlineIndex + 1
}

func (self *TextArea) closestNewlineOnLeft() int {
	cellCursor := self.contentCursorToCellCursor(self.cursor)

	newlineCellIndex := -1

	for i, c := range self.cells[0:cellCursor] {
		if c.char == "\n" {
			newlineCellIndex = i
		}
	}

	if newlineCellIndex == -1 {
		return -1
	}

	newlineContentIndex := self.cells[newlineCellIndex].contentIndex
	if self.content[newlineContentIndex] != '\n' {
		newlineContentIndex--
	}
	return newlineContentIndex
}

func (self *TextArea) GoToEndOfLine() {
	if self.atEnd() {
		return
	}

	self.cursor = self.closestNewlineOnRight()

	self.moveLeftFromSoftLineBreak()
}

func (self *TextArea) closestNewlineOnRight() int {
	cellCursor := self.contentCursorToCellCursor(self.cursor)

	for i, c := range self.cells[cellCursor:] {
		if c.char == "\n" {
			return self.cellCursorToContentCursor(cellCursor + i)
		}
	}

	return len(self.content)
}

func (self *TextArea) moveLeftFromSoftLineBreak() {
	// If the end of line is a soft line break, we need to move left by one so
	// that we end up at the last whitespace before the line break. Otherwise
	// we'd be at the start of the next line, since the newline character
	// doesn't really exist in the real content.
	if self.cursor < len(self.content) && self.content[self.cursor] != '\n' {
		self.cursor--
	}
}

func (self *TextArea) atLineStart() bool {
	return self.cursor == 0 ||
		(len(self.content) > self.cursor-1 && self.content[self.cursor-1] == '\n')
}

func (self *TextArea) isSoftLineBreak(cellCursor int) bool {
	cell := self.cells[cellCursor]
	return cell.char == "\n" && self.content[cell.contentIndex] != '\n'
}

func (self *TextArea) atSoftLineStart() bool {
	cellCursor := self.contentCursorToCellCursor(self.cursor)
	return cellCursor == 0 ||
		(len(self.cells) > cellCursor-1 && self.cells[cellCursor-1].char == "\n")
}

func (self *TextArea) atLineEnd() bool {
	return self.atEnd() ||
		(len(self.content) > self.cursor && self.content[self.cursor] == '\n')
}

func (self *TextArea) atSoftLineEnd() bool {
	cellCursor := self.contentCursorToCellCursor(self.cursor)
	return cellCursor == len(self.cells) ||
		(len(self.cells) > cellCursor+1 && self.cells[cellCursor+1].char == "\n")
}

func (self *TextArea) BackSpaceWord() {
	newCursor := self.newCursorForMoveLeftWord()
	if newCursor == self.cursor {
		return
	}

	clipboard := self.content[newCursor:self.cursor]
	if clipboard != "\n" {
		self.clipboard = clipboard
	}
	self.replaceContent(newCursor, self.cursor, "")
	self.cursor = newCursor
	self.updateCells()
}

func (self *TextArea) ForwardDeleteWord() {
	newCursor := self.newCursorForMoveRightWord()
	if newCursor == self.cursor {
		return
	}

	clipboard := self.content[self.cursor:newCursor]
	if clipboard != "\n" {
		self.clipboard = clipboard
	}
	self.replaceContent(self.cursor, newCursor, "")
	self.updateCells()
}

func (self *TextArea) Yank() {
	self.TypeString(self.clipboard)
}

func (self *TextArea) contentCursorToCellCursor(origCursor int) int {
	idx, _ := slices.BinarySearchFunc(self.cells, origCursor, func(cell TextAreaCell, cursor int) int {
		return cell.contentIndex - cursor
	})
	for idx < len(self.cells)-1 && self.cells[idx+1].contentIndex == origCursor {
		idx++
	}
	return idx
}

func (self *TextArea) cellCursorToContentCursor(cellCursor int) int {
	if cellCursor >= len(self.cells) {
		return len(self.content)
	}

	return self.cells[cellCursor].contentIndex
}

func (self *TextArea) GetCursorXY() (int, int) {
	if len(self.cells) == 0 {
		return 0, 0
	}
	cellCursor := self.contentCursorToCellCursor(self.cursor)
	if cellCursor >= len(self.cells) {
		return self.cells[len(self.cells)-1].nextCursorXY()
	}
	if cellCursor > 0 && self.cells[cellCursor].char == "\n" {
		return self.cells[cellCursor-1].nextCursorXY()
	}
	cell := self.cells[cellCursor]
	return cell.x, cell.y
}

// takes an x,y position and maps it to a 1D cursor position
func (self *TextArea) SetCursor2D(x int, y int) {
	x, y = max(x, 0), max(y, 0)
	lastRow := 0
	if len(self.cells) > 0 {
		_, lastRow = self.cells[len(self.cells)-1].nextCursorXY()
	}
	if y > lastRow {
		return
	}
	newCursor, _ := slices.BinarySearchFunc(self.cells, y, func(c TextAreaCell, row int) int {
		cellRow := c.y
		// A soft break stores the following row's coordinates, but ends the
		// preceding row for cursor placement.
		if c.char == "\n" && self.content[c.contentIndex] != '\n' {
			cellRow--
		}
		return cellRow - row
	})
	for newCursor < len(self.cells) && x > 0 && self.cells[newCursor].char != "\n" {
		x -= self.cells[newCursor].width
		newCursor++
	}
	self.cursor = self.cellCursorToContentCursor(newCursor)
	if newCursor < len(self.cells) && self.cells[newCursor].char == "\n" {
		self.moveLeftFromSoftLineBreak()
	}
}

func (self *TextArea) Clear() {
	self.layout.checkpoints = nil
	self.layout.trailerStart = 0
	self.layout.firstControl = -1
	self.layout.content = ""
	self.layout.ready = false
	self.hasChanges = false
	self.layout.unchangedRows = 0
	self.layout.unchangedCells = 0
	self.layout.revision++
	self.content = ""
	self.cells = nil
	self.cursor = 0
}

func (self *TextArea) TypeString(str string) {
	if !self.overwrite {
		self.TypeCharacter(str)
		return
	}

	state := -1
	for str != "" {
		var chr string
		chr, str, _, state = uniseg.FirstGraphemeClusterInString(str, state)
		self.typeCharacter(chr)
	}

	self.updateCells()
}
