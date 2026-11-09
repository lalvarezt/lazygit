# Algorithm choices in their calling context

This audit examines the fork at `dd68e659c`, whose production code matches
`96e977c85`. It follows the discussion of
[upstream PR #6201](https://github.com/jesseduffield/lazygit/pull/6201): a
binary search for each cell misses the draw loop's ordered traversal.

The same question applies elsewhere: what does the caller already know,
retain, or process in order? The findings below identify missed opportunities
around existing improvements. They do not establish that those measured
improvements were regressions or that these paths dominate application CPU.
Production source is unchanged.

The two already-discussed follow-ups have a separate
[post-merge implementation plan](Post-merge-performance-plan.md).

Each additional finding also has a separate implementation plan:

- [Graph occupancy](performance-plans/Graph-occupancy.md)
- [Editor navigation](performance-plans/Editor-navigation.md)
- [Graph render scheduling](performance-plans/Graph-render-scheduling.md)
- [Flat file lists](performance-plans/Flat-file-lists.md)

## Graph occupancy has cheap lookups but repeated prefix scans

**Evidence: counted operations in the real algorithm.**

The fork replaced hashed occupancy with a dense column-state array in
`f2ab2c610`. That reduces membership-check costs inside
[`getNextPipes`](../../pkg/gui/presentation/graph/graph.go), but both
`getNextAvailablePosForContinuingPipe` and
`getNextAvailablePosForNewPipe` still start scanning at column zero on every
call.

Within a single row, `traversed`, `taken`, and `continuing` only gain occupied
positions or remain fixed. A column rejected by one of these lookups cannot
become available to the same lookup later in that row. Repeatedly discovering
the same occupied prefix therefore adds work without adding information.

Disposable instrumentation counted the current occupancy checks for a row
whose final incoming pipe reaches the current commit, and for an octopus
merge with no incoming pipes:

| Incoming columns or parent count | Continuing-column checks | New-parent checks |
| ---: | ---: | ---: |
| 4 | 6 | 6 |
| 16 | 120 | 120 |
| 64 | 2,016 | 2,016 |

The counts grow quadratically in these fixtures. Hash-to-array replacement
does not remove that repeated traversal.

The next experiment should give each lookup its own forward-only cursor for
one `getNextPipes` call. The two predicates differ, so the cursors must remain
separate. Discard them at the next graph row. Check output against the
existing layout for sparse columns, merges, unrelated histories, and both
symbol sets before measuring time.

This is the closest analogue to the search walker. Its practical priority
depends on graph width, not just commit count. The current graph benchmarks
vary commit count and symbols without independently controlling active width.
The operation counts do not establish how often real repositories encounter
these shapes.

## Editor navigation ignores coordinates the parser already stores

**Evidence: caller trace and a current-code scaling probe.**

The fork's incremental parser retains each cell's content offset and display
coordinates. However,
[`MoveCursorUp` and `MoveCursorDown`](../../pkg/gocui/text_area.go) obtain the
current coordinates, then call `SetCursor2D`, which walks cells from the
beginning. `closestNewlineOnLeft` also scans the full preceding prefix to
remember its final newline.

The parser and rendering improvements reduce edit work, but navigation still
rediscovers the position in an already laid-out buffer. A local cursor move
near the end can therefore depend on the length of the entire prefix.

The probe prepared unwrapped cells before timing and repeated one Up/Down
pair. The text was identical on each line. Three 100 ms samples produced
these medians with `GOMAXPROCS=1`:

| Buffer lines | Pair near the beginning | Pair near the end |
| ---: | ---: | ---: |
| 100 | 0.134 microseconds | 4.920 microseconds |
| 1,000 | 0.144 microseconds | 49.856 microseconds |
| 10,000 | 0.166 microseconds | 412.664 microseconds |

All cases allocate zero bytes during the timed operation. Zero allocation
does not imply that the operation is independent of document size.

A useful experiment would locate a target row using existing ordered cell
coordinates, then inspect only that row. Adjacent cursor moves may also use
their current position to reach the neighboring row. These are alternatives
to compare, not a reason to add another persistent layout cache immediately.
For the left-newline lookup, a backward walk can stop at the first newline.

Tests need to preserve wide and combining characters, zero-width cells, soft
line breaks, empty lines, end-of-buffer behavior, and clicks outside content.
The probe measures navigation alone. It excludes rendering and editor
callbacks, and the largest fixture is much longer than a typical commit
message. A complete key-handling benchmark and a realistic editor profile
should set the priority of an implementation.

## Graph parallelism uses CPU count without enough workload context

**Evidence: UI caller trace and an output-checked alternative.**

[`RenderAux`](../../pkg/gui/presentation/graph/graph.go) starts
`GOMAXPROCS` goroutines, divides rows into chunks, builds a slice per chunk,
and flattens those slices. When there are fewer rows than workers, some
workers receive empty chunks.

The UI calls it with a visible slice after
[`loadPipesets`](../../pkg/gui/presentation/commits.go) has obtained cached
graph layout. Our principal graph benchmark calls `RenderCommitGraph`, which
rebuilds layout and renders the whole fixture. That benchmark remains useful
for cold work, but it does not isolate the operation repeated while moving
through an existing viewport.

A disposable sequential implementation produced identical lines for both
symbol sets, selected and unselected matches, and each tested viewport size.
The timed comparison used box symbols and cached layout:

| Visible rows | `GOMAXPROCS` | Current parallel renderer | Sequential probe |
| ---: | ---: | ---: | ---: |
| 1 | 16 | 9.001 microseconds | 1.977 microseconds |
| 8 | 16 | 26.158 microseconds | 20.034 microseconds |
| 40 | 4 | 52.572 microseconds | 72.228 microseconds |
| 40 | 16 | 103.711 microseconds | 82.223 microseconds |
| 120 | 16 | 294.723 microseconds | 360.269 microseconds |

Parallel rendering helps some larger cases. A blanket switch to sequential
rendering would lose those gains. Conversely, hardware CPU count alone does
not determine useful parallelism for a small viewport.

The next experiment should separate cold layout, cached viewport rendering,
and selected-commit redraws. Vary visible rows, graph width, symbols, and
worker count. Candidate implementations include a sequential small-work
path, fewer workers, and writing directly into disjoint ranges of one output
slice. Select thresholds from those workloads, not from these short samples.

This scheduling code predates our graph improvements. The finding limits
what our combined benchmark can say about their interactive benefit; it
does not attribute the inherited scheduling policy to those commits.

## Flat file lists still construct and discard directory structure

**Evidence: source trace only. No alternative measured.**

[`BuildFlatTreeFromFiles` and `BuildFlatTreeFromCommitFiles`](../../pkg/gui/filetree/build_tree.go)
build a complete tree, sort and compress it, collect its leaves, then retain
only those file nodes. Our indexed tree construction makes child lookup
faster, but flat-mode callers still pay for directory structures they do not
display.

This is a candidate for avoiding work at the caller's requested output
shape. It needs a flat-mode benchmark before an implementation is selected.
The current construction benchmark exercises `BuildTreeFromCommitFiles`.

A direct file ordering cannot simply sort path strings and assume parity.
The current leaf order includes hierarchical files-first, folders-first, or
mixed ordering, case sensitivity, and optional root-prefix behavior.
Working-tree flat mode also applies a stable conflict and tracked-status
ordering. Input files are not guaranteed to arrive sorted.

An experiment should compare a direct leaf builder with the current result
under every sorting mode, varied directory shapes, rename cases, and status
groups. It should include both panels and measure cold allocation and total
refresh time. Until then, the redundant intermediate structure is confirmed,
but its performance cost and a simpler compatible replacement remain open.

## Evidence and next decisions

[The evidence JSON](performance/algorithm-context-audit.json) contains every
timed sample, medians, environment metadata, commands, exact disposable test
sources, and occupancy-instrumentation edits. The
[interactive charts](performance/algorithm-context-audit.html) expose the
graph worker-count comparison and editor position dependence.

The timed probes use Go 1.27.3 on Linux amd64, `GOGC=100`, and
`GOMEMLIMIT=off`. Each case has three 100 ms samples. These are ranking
experiments, not stable scheduling thresholds or end-to-end latency claims.
CPU-frequency variation and benchmark ordering may affect comparisons.

The strongest algorithmic follow-up is the graph occupancy cursor. The
strongest caller-specific measurement is the editor navigation scan. Graph
scheduling deserves a separate viewport benchmark before choosing a policy.
Direct flat construction remains a lower-confidence candidate.

Choose any implementation separately with the user. A useful sequence is
to add the missing caller-shaped measurements, select one change, preserve
its behavior with existing interfaces, and compare complete operations.
Keep those decisions separate from the pending upstream merge.

## Reproduce the disposable probes

Use a checkout of the evidence JSON's `sourceRevision`. Keep the JSON
available outside that checkout if it predates this audit commit. The probes
use Go overlays and do not require editing production files.

For the timed probes, write `probes.graphBenchmark.content` and
`probes.editorBenchmark.content` to temporary files. Create a Go overlay whose
`Replace` map associates each probe's `replaces` path with its temporary
file. Use absolute paths in that map. Save it as
`.git/algorithm-context-audit/overlay.json`.

For occupancy counting, use `probes.occupancyTest.content` instead of the
graph benchmark file. Apply each `occupancyInstrumentation.changes` edit to
the pinned `graph.go`, and append its `append` field. Point a second overlay
at that instrumented file and the occupancy test. Save it as
`.git/algorithm-context-audit/occupancy-overlay.json`.

Run the exact argument arrays in `commands` with the recorded environment.
Use the uninstrumented overlay for timing. The occupancy counters deliberately
add work and are only for counting lookups.
