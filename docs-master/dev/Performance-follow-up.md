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

Reproduce the committed benchmarks with:

```sh
go test ./pkg/gui/context -run='^$' -bench=BenchmarkFilteredListNavigation -benchmem -benchtime=200ms -count=3
go test ./pkg/gui/filetree -run='^$' -bench='BenchmarkFileTree(RangeSelection|Navigation)' -benchmem -benchtime=200ms -count=3
go test ./pkg/gocui -run='^$' -bench='Benchmark(SearchRepeatedPrefix|SearchOrdinaryText|NearestSearchPosition|ScrollSearchPosition)' -benchmem -benchtime=150ms -count=3
```
