# Reuse editor layout during cursor navigation

Implement this follow-up after the pending upstream integration, following
the [shared implementation guidance](README.md). Preserve editor cursor
semantics while avoiding work proportional to unrelated preceding text.

## Trace complete navigation operations

Inspect [`TextArea`](../../../pkg/gocui/text_area.go), the parser, the default
editor, and mouse-position handling. The incremental parser already stores
display coordinates and content offsets in `TextAreaCell`.

`MoveCursorUp` and `MoveCursorDown` obtain the current coordinates, then call
`SetCursor2D`. That method scans cells from the beginning. The same method
handles arbitrary pointer positions, so adjacent moves and arbitrary jumps
share an implementation despite having different starting knowledge.

`closestNewlineOnLeft` scans the preceding prefix and remembers its final
newline. Its callers include start-of-line navigation and deletion.

The [audit evidence](../performance/algorithm-context-audit.json) measures an
Up/Down pair over prepared, unwrapped cells. Movement near the end grows with
document length while movement near the beginning stays small. The timed
loop allocates no memory. It excludes rendering and callbacks, and its
largest document exceeds typical commit-message sizes.

## Compare algorithms using existing layout state

Evaluate these alternatives before choosing the interface:

- Locate the requested row using ordered cell coordinates, then inspect only
  that row. Verify coordinate monotonicity for every supported cell form
  before using binary search.
- For adjacent moves, start at the current cell position and traverse to the
  neighboring row. Keep arbitrary clicks and jumps efficient through the
  same navigation module.

Prefer existing cell metadata over another persistent row index unless
measurements justify the index's memory and invalidation costs. Keep byte
offsets, grapheme indices, and display columns distinct.

Evaluate the left-newline lookup separately. A backward walk can stop at the
first preceding newline. Preserve its existing handling of soft wraps and
the mapping to the original content. Share positional logic where it is
truly the same, without combining unrelated behavior changes.

## Preserve cursor and edit semantics

Exercise navigation through the same interfaces used by keyboard and mouse
callers. Cover:

- Empty content, empty lines, final newlines, and the start and end of text.
- Hard and soft line breaks, wrapping changes, and unusually long lines.
- Wide graphemes, combining characters, zero-width cells, and multiple cells
  associated with one content position.
- Columns beyond a row's text, negative coordinates, and rows outside the
  document, preserving current clamping or no-movement behavior.
- Repeated Up and Down, clicks, Home, and operations that delete to the start
  of a line when they use the changed helper.

Compare cursor byte offsets and displayed coordinates with the existing
implementation across generated content and target coordinates. Confirm
that moving the cursor leaves content unchanged. If the comparison reveals
a pre-existing correctness defect, discuss that finding before expanding
the change to fix it.

## Measure navigation and the complete key path

Vary document length, cursor depth, line length, and wrapping independently.
Include ordinary commit messages as well as large synthetic documents.
Measure repeated adjacent moves and arbitrary position jumps separately.

Keep parsing outside the isolated navigation benchmark. Also measure the
default editor's complete key path, including `RenderTextArea`, so a faster
lookup is not mistaken for the same improvement in visible response time.
Profile realistic editor use before choosing this work over other hotspots.

Record allocation counts and retained memory. A result should show reduced
prefix dependence without moving a full-document scan into another method
or adding an unnecessary cache.

Inspect the original navigation code and the fork's parser/layout commits
before choosing commit placement. Keep an inherited navigation improvement
distinct from any fixup that changes layout work already on this branch.
