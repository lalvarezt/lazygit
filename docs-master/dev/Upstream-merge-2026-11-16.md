# Upstream integration on 2026-11-16

The integration preserves the fork's performance implementations and adds
upstream's fixes. Search drawing combines a binary seek per visible row with
upstream's walker within the row. Both overlapping-match priority and gutter
highlighting follow upstream's corrected behavior.

## Pinned inputs

- Task-start fork: `149b5091ab10df3f72cc6e3902ad88f6d912799d`.
- Immediate merge parent after the runner repair: `c468e9da0`.
- Fetched upstream master: `54cc6af1d77b5a6cec3dc9ad0b203afdb4ad3451`.
- Merge base: `bd4062a1e745b47d87ee8d0b24649529ed28b499`.
- Prototype with upstream's unmodified walker: `09ce010358d2f781ce3647e6ba235239fa287f0e`.
- Validated code snapshot: `7735277d1d35c3033a29beafd53414101d0d59a9`.

Local refs under `refs/merge-upstream/2026-11-16/` retain these snapshots.
The task-start working tree was clean. Existing fixups remain separate.
The runner repair changes no runtime source, so the task-start measurements
also cover the immediate pre-merge runtime state. The final merge differs
from the validated code snapshot only in its evidence files.

The [evidence file](performance/upstream-merge-2026-11-16.json) contains the
complete commit inventories, all 327 fork paths, source identities, controls,
binary hashes, raw samples, and derived results.

## Integration coverage

The upstream range contains 74 commits and changes 104 paths. The fork range
contains 137 commits and changes 327 paths. Thirteen paths overlap.

| Fork customization | Upstream interaction | Preservation check |
| --- | --- | --- |
| Batched screen writes and rectangle fills | Search drawing changes `setCharacter` and its callers | Keep the shared cell-buffer lock and uncovered-column fills |
| Binary search highlighting | Upstream adds a draw-local walker and fixes overlap priority | Seek past hidden matches per row, then walk visible matches |
| Incremental search, reusable results, KMP scanning | Paste adds query updates | Keep invalidation, query changes, result ownership, and current-match selection |
| Search while a renderer runs | Paste enters through `OnPromptContentChanged` | Keep loading, pending selection, cancellation, and replacement paths |
| Incremental editor parsing, rendering, and navigation | Paste adds bulk `TypeString` callers | Keep one render per key and explicit renders after bulk edits |
| Model-owned branch-marker ancestry | Detached HEAD adds subject data and status rendering | Keep model setters on the UI thread and configuration invalidation |
| File-tree construction, ordering, visibility, and selection | Upstream renames a test helper | Preserve implementations and both panels' tests |
| Graph allocation and worker scheduling | No production overlap | Preserve source identity and tests |
| ANSI writer, Unicode handling, and raw lines | No production overlap outside drawing | Preserve parser and writer code, metadata, and tests |
| Combined-diff copy guard | No production overlap | Preserve integration test and consuming action |
| Terminal-reader shutdown and tcell changes | No upstream vendor changes | Preserve source identity and race tests |
| Filtered ranges, counts, and filename escaping | No production overlap | Preserve source identity and tests |
| Performance reports, tooling, and README | Sponsor removal and added merge fixtures | Preserve historical reports and repair the benchmark registration |

Clean merges retain `SetCommits`, `SetSubCommits`, `SetBranches`, and
`SetMainBranchNames`. The status refresh adds the rebased branch name without
moving model writes off the UI thread. Search drawing keeps a fresh walker
per frame. Prefix reuse, offscreen swaps, and dirty-row tracking stay intact.

Upstream state changes replace shared mutable state with locked reads and
updates. Detached-HEAD display, merge menus, commit-message paste, copy and
clear actions, input blocking, sponsor removal, and test-driver paste support
are included. Generated files are regenerated. All upstream paths outside
the fork overlap match upstream exactly. Of the 327 existing fork paths,
313 retain their exact blobs. The remaining 14 are the 13 overlaps and the
comparison runner. Added tests and evidence are recorded separately.

## Correctness

The task-start build, unit tests, and lint pass. The final implementation
passes `just format`, `just generate`, `just build`, `just unit-test`,
`just lint`, and `just e2e`. It also passes `go test -race ./... -short` and
a Windows cross-build of the application.

The fork's overlap-color expectation changes from `yyc` to `ycc`, matching
upstream's intentional rule that the selected overlapping match shows in
full. Added differential drawing tests cover streamed append, changed
queries, missing matches, Unicode, wrapping, horizontal and vertical origins,
the inclusion gutter, selection changes, and offscreen replacement.

## Performance discovery

Five alternating pairs compare the task start with the unmodified-walker
prototype. Both binaries use identical benchmark fixtures, optimized builds,
one pinned CPU, `GOMAXPROCS=1`, `GOGC=100`, and `GOMEMLIMIT=off`. Each sample
runs for 100 ms. Setup and initial search are outside timing. Each iteration
draws the whole visible view, including the existing content-cell scan and
screen-buffer locking.

| Fixture | Fork median | Upstream walker median | Ratio |
| --- | ---: | ---: | ---: |
| 22 rows, 100,000 columns, left edge | 0.108 ms | 1.045 ms | 9.71 |
| 22 rows, 100,000 columns, middle | 3.452 ms | 4.740 ms | 1.37 |
| 22 rows, 100,000 columns, right edge | 6.958 ms | 8.160 ms | 1.17 |

The walker scans matches outside the visible columns. At the start of a new
row, it also scans the remaining matches in the previous row. The fork's
existing content-cell scan still costs time when horizontally scrolled.
These results include that cost. The user chose a binary seek to the first
visible cell of each row, followed by the walker for that row. The walker
remains local to a draw and uses the current result slice.

## Final performance

The matched comparison covers 183 cases in gocui, GUI editors, commit
viewports, and native model updates. Each side has five samples per case.
It includes 60 sparse, dense, overlapping, Unicode, and missing-match draw
cases, plus 12 long-line cases. Viewports cover the start, middle, and end,
with wrapping enabled and disabled.

Seven alternating pairs recheck all 12 long-line cases, six editor cases,
18 commit viewports, and three ancestry-update cases. The same pinned CPU
and GC controls apply. Rechecks use 100 ms per sample. Both broad and paired
results remain in the evidence, including slower rows.

| Long-line fixture, 22 rows × 100,000 columns | Fork median | Integrated median | Change |
| --- | ---: | ---: | ---: |
| Left edge | 0.105 ms | 0.072 ms | -31.4% |
| Middle | 3.693 ms | 3.711 ms | +0.5% |
| Right edge | 7.356 ms | 7.395 ms | +0.5% |

These are final paired measurements, separate from the prototype table.
Drawing adds no allocations. The fork's existing scan through content cells
still dominates far-horizontal scrolling. This merge does not optimize that
scan.

The broad run shows two editor cases about 5% slower. Paired measurements
reduce the `bytes=1024` input difference to 3.8%, or 26 microseconds for the
whole 1,020-byte input workload. The million-byte edit case becomes 5.3%
faster in the paired run. Allocation counts remain unchanged. These fixtures
do not measure individual keystroke latency.

The broad run also shows two commit viewports about 10% slower and the
100,000-commit model update 6% slower. Paired viewport ratios range from
0.978 to 1.048. The paired 100,000-commit model update is 7.897 ms before and
7.683 ms after. No history-sized allocation returns to viewport rendering.
Ancestry work remains on input replacement, with its cost measured separately.

No material regression remains confirmed on the measured fork performance
paths. Small timing differences remain visible in the data. Unchanged graph,
file-tree, filtered-list, writer, and terminal-reader implementations retain
their source identities and correctness coverage. Their historical benchmark
reports are preserved, not regenerated as new measurements.

## Shared-state correctness costs

Five alternating pairs measure complete state reads, unchanged updates, and
changed updates with empty and 1,000-entry shell histories. Reads also run
with eight goroutines. Version adapters select the old getter/save API or
the new read/update API outside timing. Configuration and fixture setup are
excluded. Calls include locking, parsing, serialization, and writes.

The initial storage is tmpfs. A second run places temporary configuration
under the repository on btrfs to include persistent-filesystem fsync cost.
Both runs use warm file caches.

| btrfs state operation | Fork median | Integrated median |
| --- | ---: | ---: |
| Read empty state | 1.5 nanoseconds | 14.4 microseconds |
| Change empty state | 18.2 microseconds | 938 microseconds |
| Read 1,000-entry history | 1.5 nanoseconds | 800 microseconds |
| Update without changes, 1,000 entries | 785 microseconds | 1,515 microseconds |
| Change state, 1,000 entries | 770 microseconds | 2,600 microseconds |

The old read returns an in-memory pointer. The new read locks and parses the
current file. A 1,000-entry read allocates about 408 kB, previously zero.
Changed updates allocate about 1.53 MB, previously 1.11 MB. Unchanged updates
avoid rewriting the file but still read and serialize it. Eight-reader
throughput is about 835 microseconds per operation on the pinned CPU.

These costs preserve upstream's cross-instance correctness and atomic writes.
State-access benchmarks exclude suggestion filtering and UI rendering, so
they do not establish full prompt latency. The integration tests exercise the
consuming UI paths; config tests cover eight concurrent updating instances.

## Gaps

The initial `just perf-compare` invocation failed with
`Missing description: BenchmarkViewWriteSGR`. Fixup `c468e9da0` registers the
benchmark introduced by `ca7ea3ef1`. Build, unit tests, lint, and fixture
preparation pass for the repair. The merge adds descriptions for its new draw
fixtures. Package-only comparisons leave the historical README report intact.

One broad comparison exceeded the tool's five-minute call limit after saving
the baseline. The complete candidate run resumed from that checkpoint. An
interrupted retry with incomplete color controls was discarded. Reported
runs remove `NO_COLOR` and set `FORCE_COLOR=1` and `CLICOLOR_FORCE=1`.

Native Windows execution, real-terminal interaction, cold physical-storage
latency, contended state writers, and network filesystems are not benchmarked.
The Windows build is compile coverage. State throughput excludes suggestion
filtering. The broad runner records source and fixture checkpoints; paired
runs additionally retain binary hashes. These limits preclude a claim that
every workload has zero regression.
