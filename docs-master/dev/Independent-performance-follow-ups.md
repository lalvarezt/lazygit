# Independent performance follow-ups

The five follow-ups that do not touch upstream PR #6201 are implemented.
Prepared-query reuse and integration of that PR remain deferred. The changed
source files have no overlap with the three files in the inspected PR.

## Implementation

| Follow-up | Result | Commits |
| --- | --- | --- |
| Graph occupancy | Independent forward-only column cursors for each graph row | `07e4bf9fe`, a fixup for `f2ab2c610` |
| Editor navigation | Coordinate-based row lookup and a backward walk to the nearest preceding break | `e7fe1a597`, `43d454348` |
| Graph rendering | One output slice and a pipe-count budget for worker selection | `2f9bb7e38`, `b41c6edca` |
| Cached tree selections | Both panels wrap only the requested visible-row slice | `a1402046e`, with coverage restored in `240a2cfe1` |
| Flat file lists | Shared ordering descriptor and direct leaf construction | `89a8dfabf`, `8b4df7052`, `62b93fab9`, `0ab8449ad` |

The flat builder preserves case-alias and file/directory-collision behavior
through the existing tree builder. Its ordinary path normalizes sort keys
once, compares path components with the shared ordering rules, and constructs
only file nodes. Duplicate paths retain the first file, as before.

The graph renderer budgets one worker per 128 pipes, capped by available
CPUs and visible rows. Small workloads run directly. This is a private
heuristic supported by the recorded fixtures, not an optimal threshold for
every graph shape or machine.

## Matched measurements

Each row describes a separate operation. Times are microseconds per operation;
bytes are allocated bytes per operation. These synthetic measurements do not
establish end-to-end UI latency improvements.

| Operation | Before, microseconds | After, microseconds | Before bytes | After bytes |
| --- | ---: | ---: | ---: | ---: |
| Graph columns | 24.576 | 14.157 | 19840 | 19840 |
| Editor navigation | 469.156 | 0.160 | 0 | 0 |
| Editor full key path | 459.817 | 0.281 | 0 | 0 |
| Line-start navigation | 152.790 | 0.068 | 0 | 0 |
| Cached ranges | 36.403 | 0.032 | 24 | 16 |
| Cold ranges | 67.778 | 30.583 | 24 | 16 |
| Flat refresh | 10472.138 | 4547.765 | 11572736 | 6623192 |
| Graph narrow viewport | 25.126 | 7.948 | 13328 | 10080 |
| Graph branching viewport | 258.303 | 181.014 | 519989 | 517868 |

Graph columns use a 256-column row. Editor navigation and complete key
handling each measure an Up/Down pair near the end of 10,000 lines. Line-start
navigation uses the same document size. Tree selections request one row near
the end of a 10,000-file tree. Cold selection explicitly invalidates the
visible-row index, retaining its allocated storage.

Flat refresh rebuilds a 10,000-file commit-file list and includes its normal
view-model work. Graph viewports use 40 narrow rows with GOMAXPROCS=16 and
120 branching rows with GOMAXPROCS=4. Other sizes, positions, and CPU settings
remain in the evidence. Editor operations near the beginning can gain less
or pay a small lookup cost; they are not omitted from the samples.

The [evidence JSON](performance/independent-follow-ups.json) records all
samples, benchmark argument arrays, baseline overlays, implementation
revisions, environment settings, and limits. The
[comparison chart](performance/independent-follow-ups.html) shows the selected
operations above. Final graph viewport comparisons use five 100 ms samples;
other final comparisons use three 100 ms samples. Disposable worker-policy
exploration used shorter samples and is labeled separately.

Baseline overlays substitute only the named production file at the recorded
revision. The benchmark fixtures are identical for each matched comparison.
Run argument arrays directly, or quote regular expressions containing shell
metacharacters when translating them to a shell command.

## Validation

Formatting, build, unit tests, and lint pass. Color-sensitive checks use the
fork's documented environment: remove NO_COLOR and set FORCE_COLOR=1 and
CLICOLOR_FORCE=1.

The final race checks pass for gocui, filetree, and graph packages. All tests
in the file, commit, and UI integration groups pass through `just e2e file
commit ui`. Existing selection tests remain alongside the new differential
range tests. Graph output tests cover both symbol sets, selected and
unselected commits, empty ranges, and multiple CPU counts. Flat-order tests
cover all sort modes, case sensitivity, root settings, generated layouts,
Unicode case mappings, duplicates, case aliases, and file/directory collisions.

Fixups and the editor's message-only amend commit remain separate for review.
No branch history was rewritten, and no upstream changes were merged.
