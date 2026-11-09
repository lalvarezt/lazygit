# Regression audit against the previous snapshot

The audit compares the previous README snapshot, 818a17b1d21a701f43fe51474064445e7ee1266f, with the production checkpoint cc2e9b8b0. It covers the 111-commit range through the audit's initial head, 8626cc6c9, and the subsequent graph and ancestry refinements, including model-owned ancestry.

## Confirmed findings

The editor coordinate lookup copied cells through a generic comparison callback. Fixup 8626cc6c9 uses indexed cell access while preserving the shared coordinate lookup. All six vertical-movement probes improve.

The graph scheduler rounded its work budget down and ran large single-CPU workloads on the caller. Fixup 7eb64e4fe rounds the existing 128-pipe budget up and retains direct execution only within that budget. This is the approved scheduling policy. The earlier allocation and 96-pipe proposals are superseded.

The upstream branch-marker fix at 89c6f0ce5 added a complete unmerged-history map to every commit viewport render. The graph-only benchmark excluded this caller work. Interleaved probes confirm a substantial regression and history-sized allocation even when a local-main ancestry walk cannot affect output. Fixups fe856e684 and b920143e7 skip the map when no loaded, unmerged main tip exists and skip the ancestry calculation when no eligible branch marker is visible. Existing marker output is preserved.

The permanent benchmark then exposed the remaining visible-marker, unmerged-main case. Commit cc2e9b8b0 moves the derived ancestry sets onto the repository model. Commit, subcommit, branch, and main-branch configuration updates recompute the affected sets in the same UI thread update that replaces their inputs. Rendering reads the prepared set. Normal and rebase-only refreshes, both subcommit loaders, configuration reloads, and repository-state restoration use this interface. Main-branch names are copied, and each repository owns its own derived state.

The matched 18-case lifetime comparison confirms that repeated rendering no longer allocates a history-sized map. The ancestry algorithm still runs when model inputs change. Its time and allocations are measured separately in [the lifetime evidence](performance/ancestry-lifetime.json), so the moved work remains visible.

Preserving the current search match on its line adds a small measured cost to nearest-match lookup. Interleaved samples confirm it. This upstream guard fixes losing the second match on a line during refresh; its behavior is retained. The full integration suite passes.

## Measurements

The [187-case comparison](performance/snapshot-regression-audit.json) retains five samples per side against the previous snapshot, including allocated bytes and allocation counts. It combines the original 169-case snapshot baseline with 18 added commit-viewport cases and fresh current-checkpoint samples. Provenance identifies each input run, fixture hashes, revisions, and commands. The [interactive chart](performance/snapshot-regression-audit.html) exposes every measured row.

The new viewport benchmark renders 40 rows from three loaded-history sizes, with absent, merged, or unmerged local-main tips, and with another branch marker absent or visible. Graph rendering is disabled so it cannot hide the marker bookkeeping cost. Model setup is outside the timed loop. A typed function call is selected before timing so the renderer interfaces can use identical portable fixtures. The synthetic unmerged main tip has only a merged parent; its known ancestry is supplied outside timing. Native model updates are benchmarked separately.

The README remains anchored to fork base 8e3ac76991f5, separately from this previous-snapshot audit. Its expanded comparison contains 187 cases against cc2e9b8b0. No slower result is filtered out.

The 100,000-commit, merged-main, visible-marker case has a slower median in the broad snapshot run. A separate nine-sample, alternating before/after recheck pins both binaries to the same performance core. It reduces the difference to about 4%, with no history-sized allocation. Both runs remain recorded; the audit does not claim every timing is faster than the previous snapshot.

## Standards and spec review

The standards axis found no additional documented-standard violation in the inspected performance changes. It checked root and gocui contribution guidance, navigation, SGR parsing, incremental search, flat ordering, cached selections, graph scheduling, and benchmark fixture transformations. No release docs under docs/ were edited.

The spec axis confirmed the commit-render allocation problem. It found no further confirmed mismatch in the inspected editor, flat-builder, range-selection, or streaming-search requirements. Prepared-query reuse and integration of the deferred upstream PR remain untouched.

The static passes prioritized custom performance changes. Broad upstream command/controller changes were inventory-screened and selected control flows were inspected, rather than exhaustively reviewed line by line. Full integration testing covers the resulting branch. This audit is not proof that every possible input is regression-free.

An async search-replacement boundary remains an unconfirmed hypothesis, not a finding: old and incoming buffers may contain matches at different positions while Enter is pressed before the first swap. Existing slow-stream and replacement integration tests pass; that exact timing has not been separately reproduced. No speculative fix was added.

## Validation

Formatting, build, all unit tests, lint, and the full integration suite pass. Initial race checks cover gocui, filetree, graph, branchstack, and presentation; the model-owned change also passes race checks for types, context, presentation, and controller helpers. Model tests cover either input arrival order, branch-tip changes, local and subcommit replacement, changed status and parent data, configuration changes, copied configuration slices, zero state, and repository isolation. Color-sensitive checks remove NO_COLOR and set FORCE_COLOR=1 and CLICOLOR_FORCE=1.

Reviewable fixups remain separate. The one-line benchmark registration was retargeted to the README comparison commit with explicit approval. No upstream changes were merged or PRs created during this audit.

## Model-update work

These native model-update fixtures rebuild ancestry for an unmerged main tip in an all-unmerged history. The main tip has no parents. They measure updating local commits, not loading git history or rendering a viewport. Values are five-sample medians.

| Loaded commits | Update time, milliseconds | Allocated bytes |
| --- | ---: | ---: |
| 1000 | 0.066 | 109024 |
| 10000 | 0.719 | 873536 |
| 100000 | 7.757 | 6989760 |

This cost occurs when the model inputs are replaced. Scrolling and selection changes reuse the derived result.
