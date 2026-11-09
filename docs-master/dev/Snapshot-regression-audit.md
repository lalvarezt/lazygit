# Regression audit against the previous snapshot

The audit compares the previous README snapshot, 818a17b1d21a701f43fe51474064445e7ee1266f, with the production checkpoint b920143e7. It covers the 111-commit range through the audit's initial head, 8626cc6c9, and the subsequent graph and ancestry refinements.

## Confirmed findings

The editor coordinate lookup copied cells through a generic comparison callback. Fixup 8626cc6c9 uses indexed cell access while preserving the shared coordinate lookup. All six vertical-movement probes improve.

The graph scheduler rounded its work budget down and ran large single-CPU workloads on the caller. Fixup 7eb64e4fe rounds the existing 128-pipe budget up and retains direct execution only within that budget. This is the approved scheduling policy. The earlier allocation and 96-pipe proposals are superseded.

The upstream branch-marker fix at 89c6f0ce5 added a complete unmerged-history map to every commit viewport render. The graph-only benchmark excluded this caller work. Interleaved probes confirm a substantial regression and history-sized allocation even when a local-main ancestry walk cannot affect output. Fixups fe856e684 and b920143e7 skip the map when no loaded, unmerged main tip exists and skip the ancestry calculation when no eligible branch marker is visible. Existing marker output is preserved.

A remaining regression is now visible in the permanent benchmark: when an eligible marker is visible and a local-main tip is unmerged, the history map is rebuilt on every render. The early returns do not fix this lifetime problem. Moving the derived result to model updates requires a model/renderer interface change and a new commit. That design decision is pending. The README retains these slower cases.

Preserving the current search match on its line adds a small measured cost to nearest-match lookup. Interleaved samples confirm it. This upstream guard fixes losing the second match on a line during refresh; its behavior is retained. The full integration suite passes.

## Measurements

The [187-case comparison](performance/snapshot-regression-audit.json) retains five samples per side against the previous snapshot, including allocated bytes and allocation counts. It combines the original 169-case snapshot run with 18 added commit-viewport cases and fresh current-checkpoint samples. Provenance identifies each input run, fixture hashes, revisions, and commands. The [interactive chart](performance/snapshot-regression-audit.html) exposes every measured row.

The new viewport benchmark renders 40 rows from three loaded-history sizes, with absent, merged, or unmerged local-main tips, and with another branch marker absent or visible. Graph rendering is disabled so it cannot hide the marker bookkeeping cost. Model setup is outside the timed loop.

The README remains anchored to fork base 8e3ac76991f5, separately from this previous-snapshot audit. Its expanded comparison contains 187 cases against b920143e7. No slower result is filtered out.

## Standards and spec review

The standards axis found no additional documented-standard violation in the inspected performance changes. It checked root and gocui contribution guidance, navigation, SGR parsing, incremental search, flat ordering, cached selections, graph scheduling, and benchmark fixture transformations. No release docs under docs/ were edited.

The spec axis confirmed the commit-render allocation problem. It found no further confirmed mismatch in the inspected editor, flat-builder, range-selection, or streaming-search requirements. Prepared-query reuse and integration of the deferred upstream PR remain untouched.

The static passes prioritized custom performance changes. Broad upstream command/controller changes were inventory-screened and selected control flows were inspected, rather than exhaustively reviewed line by line. Full integration testing covers the resulting branch. This audit is not proof that every possible input is regression-free.

An async search-replacement boundary remains an unconfirmed hypothesis, not a finding: old and incoming buffers may contain matches at different positions while Enter is pressed before the first swap. Existing slow-stream and replacement integration tests pass; that exact timing has not been separately reproduced. No speculative fix was added.

## Validation

Formatting, build, all unit tests, lint, the full integration suite, and race tests for gocui, filetree, graph, branchstack, and presentation pass. Color-sensitive checks remove NO_COLOR and set FORCE_COLOR=1 and CLICOLOR_FORCE=1.

Reviewable fixups remain separate. The one-line benchmark registration was retargeted to the README comparison commit with explicit approval. No upstream changes were merged or PRs created during this audit.
