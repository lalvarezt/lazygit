# Performance follow-up, 4 November 2026

This investigation starts at `4268a97ed`, after the six passes in
[the earlier performance report](Performance-summary.html). Those passes cover
view drawing, diff streaming, search highlighting, paste, editors, filtered lists,
file-tree construction, selection recovery, rectangle fills, and batched screen
writes.

The results below are staged microbenchmarks on Linux amd64 with an Intel
Core i7-1370P. Each figure is the median of three runs with optimized Go
compilation. Each change has its own baseline. The ratios describe the measured
operation and cannot be combined into an application speedup.

| Operation and workload | Before | After | Speedup |
| --- | ---: | ---: | ---: |
| Filtered navigation, 10,000 matching branches | 24.525 µs | 30.44 ns | 806× |
| Select a file-tree range, 10,000 files | 157.221 ms | 176.778 µs | 889× |
| File-tree navigation, 10,000 files | 106.553 µs | 20.46 ns | 5,208× |
| Search 2,000 repeated-prefix lines, 32-grapheme query | 64.700 ms | 4.505 ms | 14× |
| Search 2,000 repeated-prefix lines, 128-grapheme query | 172.172 ms | 4.032 ms | 43× |
| Nearest result near the end of 100,000 matches | 39.271 µs | 10.25 ns | 3,831× |
| Distant scroll down across 100,000 matches | 44.045 µs | 24.38 ns | 1,807× |
| Distant scroll up across 100,000 matches | 35.958 µs | 25.45 ns | 1,413× |

The [measurement data](performance-follow-up.json) contains all 30 comparisons,
including the individual samples, allocation counts, smaller workloads, ordinary
search text, and file-tree refresh costs.

Filtered selections now copy only their requested rows. At 10,000 matches,
navigation allocates 8 bytes instead of 81,920 bytes. Multi-selection uses the
same range accessor.

Both file panels share a bounded traversal for selecting ranges and a visible-row
index for navigation. The index rebuilds when the tree or collapsed paths change.
The 10,000-file fixture retains 188,416 bytes in the node and depth arrays. A
5,000-file directory-removal refresh measured 4.367 ms before indexing and
3.892 ms after. Navigation benchmarks repeatedly attempt to advance near the end
of the list; they include cursor clamping and selected-item retrieval.

Search rebuilding uses a prefix table over graphemes. Its circular column buffer
preserves positions for wide and combining characters. The repeated-prefix fixture
has 2,000 lines of 200 `a` characters; its longer queries end in `b` and have no
matches. Rendered-line layout is warmed before timing. Ordinary matching text
stays near 0.67 ms. The query adds two integer arrays proportional to its length.

The search-navigation benchmarks isolate lookup costs. The scroll benchmarks
reset the selection far outside the viewport on every iteration. Ordinary small
scrolls usually cross fewer matches and have smaller absolute savings.

Each implementation passed `just format`, `just build`, `just unit-test`,
`just lint`, and the full `just e2e` suite. Unit tests need `NO_COLOR` unset in
this execution environment because existing color assertions expect ANSI output.
Tests cover filtered ranges, collapsed paths, replaced trees, Unicode search,
overlapping matches, duplicate result rows, gaps, and status callbacks.

Two prototypes did not meet the 5× threshold and were not added to the application.
Reusing text-layout storage reduced the 4 KB editor-input benchmark from roughly
642 ms to 433 ms. A single-pass flattening prototype improved its 10,000-file
benchmark from roughly 306 µs to 122 µs. Editor CPU profiling still shows full
text layout, view reconstruction, and garbage collection as substantial costs.

## Live profiling to rank the next change

An optimized binary at `5b094f878` ran against an isolated shared clone of this
repository. Automatic fetching and update checks were disabled. Go CPU profiles
covered the app process; Git Trace2 recorded subprocess elapsed time separately.
Each workload was profiled once. These are ranking observations, not matched
before-and-after comparisons or repeated latency measurements.

| Workload | Input and terminal | Profile duration | App CPU |
| --- | --- | ---: | ---: |
| Commit navigation | 400 down keys at roughly 40 ms intervals, 160 × 48 | 18 s | 3.55 s |
| File navigation | 400 down keys, 10,000 untracked files in 100 directories, 160 × 48 | 18 s | 1.17 s |
| Commit-description paste | 4,067 bytes with wide and combining characters, 100 × 30 | 12 s | 1.21 s |

Commit navigation spent about 0.04 CPU seconds rendering graph rows. Its 397
`git show` commands totalled 1.69 seconds of subprocess elapsed time. File
navigation issued 400 diffs totalling 0.39 seconds. Git elapsed totals can overlap
and do not measure Git CPU time.

The editor's final marker was visible by the 610 ms observation, with observations
200 ms apart. Its profile recorded 0.61 CPU seconds in background garbage
collection and 0.40 seconds in the editor handler. Within that handler, layout
used 0.24 seconds and view reconstruction used 0.16 seconds. Inclusive samples
must not be added to their parent totals. The editor remains the next candidate
for incremental layout and rendering; graph work accounts for little CPU in the
live navigation workload.

The [live ranking data](performance-live-ranking.json) preserves the workload
descriptions and Git command totals. Raw profiles and terminal captures remain
locally under `.git/perf-investigation/`. The dedicated tmux profiling pane is
`%71` in the `lazygit-perf` session; the app has exited.

Reproduce the committed benchmarks with:

```sh
go test ./pkg/gui/context -run='^$' -bench=BenchmarkFilteredListNavigation -benchmem -benchtime=200ms -count=3
go test ./pkg/gui/filetree -run='^$' -bench='BenchmarkFileTree(RangeSelection|Navigation)' -benchmem -benchtime=200ms -count=3
go test ./pkg/gocui -run='^$' -bench='Benchmark(SearchRepeatedPrefix|SearchOrdinaryText|NearestSearchPosition|ScrollSearchPosition)' -benchmem -benchtime=150ms -count=3
```
