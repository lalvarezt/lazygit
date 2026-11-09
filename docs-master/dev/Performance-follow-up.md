# Performance follow-up, 4 November 2026

The [current comparison](../../README.md#performance) uses `8e3ac7699` as its
baseline, kept at `performance-baseline-crlf-20261104`. This separate baseline
includes the shared CRLF correction. The branch retains its original commit
hashes, with the CRLF demonstration and correction applied as tip fixups.

The full run measured `d9f407637`. Revision `fdbe79416` preserves the original
branch history and has the same Git tree; the data records this equivalence.
The focused reports below retain their original measured source revisions.
The remaining weakest cases are 100-file tree construction and ordinary text
searches, taking about 32 µs and 440–626 µs respectively in this run.

This investigation starts at `4268a97ed`, after the first six performance
passes. Those passes cover
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

The [measurement data](performance-follow-up.json) contains all 110 comparisons,
including individual samples, allocation counts, smaller workloads, ordinary
search text, file-tree refresh costs, and unshipped prototypes.

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
must not be added to their parent totals. This ranking led to the incremental
editor change below; graph work accounts for little CPU in the live navigation
workload.

The [live ranking data](performance-live-ranking.json) preserves the workload
descriptions and Git command totals. Raw profiles and terminal captures remain
locally under `.git/perf-investigation/`. The dedicated tmux profiling pane is
`%71` in the `lazygit-perf` session; the app has exited.

## Incremental editor results

[View the benchmark and live CPU charts](Performance-editor.html).

`d0f205ca2` preserves completed hard lines when text changes and reuses their
rendered rows. Changes to the final paragraph invalidate trailer lines when
needed. Each key still runs its editor handler. Tabs, control sequences, and
automatic hyperlinks keep the normal writer. Graphemes own their text, so retained
lines cannot retain whole previous message strings. Buffer writes, resets, swaps,
and relevant configuration changes invalidate the rendering cache.

| Typing workload | Before | After | Speedup |
| --- | ---: | ---: | ---: |
| Shared editor, 256 bytes | 2.271 ms | 0.260 ms | 8.7× |
| Shared editor, 1,024 bytes | 34.801 ms | 2.094 ms | 16.6× |
| Shared editor, 4,096 bytes | 592.204 ms | 24.256 ms | 24.4× |
| Commit description, 256 bytes | 2.575 ms | 0.268 ms | 9.6× |
| Commit description, 1,024 bytes | 39.842 ms | 2.066 ms | 19.3× |
| Commit description, 4,096 bytes | 662.229 ms | 24.157 ms | 27.4× |

These medians use the existing newline-separated typing benchmarks. At 4 KB,
cumulative allocation falls from 1.57 GB to 13.95 MB. Editing within one long hard
line still rebuilds that line. All six typing comparisons and their samples are
included in the measurement data.

The same 4,067-byte live paste workload used 0.15 app CPU seconds after the change,
down from 1.21 seconds before. Background garbage collection fell from 0.61 CPU
seconds to 0.05 seconds. The final marker appeared at the first observation,
205 ms after paste, compared with the 610 ms observation before. The 200 ms
observation interval cannot establish a precise latency ratio. Terminal captures
match apart from the version footer. This is one matched profile pair.

Differential tests compare layout with the full parser through 1,000 editor
actions and compare rendered content with the normal writer. They cover wrapping,
Unicode, trailer changes, skipped renders, external writes, copied snapshots,
terminal widths, and writer fallback. Build, unit tests, lint, full integration
tests, and the gocui race checks pass.

## Parser checkpoints for long paragraphs

`09d853084` expresses the full wrapping algorithm as a streaming parser.
`dc04d2880` resumes that parser at completed hard or soft line transitions. A
checkpoint retains the partial word, grapheme lookahead, row position, and
footnote and trailer state. Edits that join grapheme clusters resume before the
saved lookahead. Final-paragraph changes still invalidate earlier trailer lines.

| Workload | Before | After | Speedup |
| --- | ---: | ---: | ---: |
| 4 KB paragraph, before parser preparation | 713.479 ms | 22.385 ms | 31.9× |
| 4 KB paragraph, prepared parser baseline | 622.605 ms | 22.385 ms | 27.8× |
| 4 KB newline-separated input, prepared baseline | 26.211 ms | 15.781 ms | 1.7× |
| 100 KB bulk paste, prepared baseline | 3.162 ms | 2.823 ms | 1.1× |

The paragraph benchmark repeats `body text ` without explicit newlines and wraps
at 72 columns. The nominal 4,096-byte benchmark contains 4,090 bytes of repeated
text. The checkpoint and scratch arrays retain 9,408 bytes for that paragraph and
49,968 bytes for the newline-separated fixture. Clearing the text releases both
arrays. The rejected grapheme-cache prototype retained 451,008 bytes and
246,592 bytes respectively in its additional arrays.

A matched live paste uses 4,111 bytes in a single paragraph at 100 × 30 terminal
size. Against the prepared-parser baseline, app CPU falls from 1.44 seconds to
0.18 seconds over a 12-second profile. Terminal captures are identical. The final
marker is observed at 514 ms before and 105 ms after, with 100 ms observation
intervals. This is one profile pair and cannot establish a precise latency ratio.
The baseline binary uses the source from `09d853084` through a Go build overlay.

Tests cover 900 replacements inside paragraphs, in addition to editor-key
sequences and Unicode boundaries. A temporary independent oracle compares these
incremental layouts directly with the old parser. The streaming full parser also
matches the old parser on 1,000 generated inputs. Build, unit tests, lint, full
integration tests, and gocui race checks pass.

These initial checkpoints follow line transitions. The subsequent unbroken-text
measurements below cover the remaining case.

A separate ranking profile loads a 50,000-line diff, enters a search, and sends
120 up keys. It records 1.81 app CPU seconds over 14 seconds, including 0.56 seconds
in view writes and 0.10 seconds updating search positions. This is a ranking run,
not a before-and-after comparison.

## Unbroken input and editor row ownership

Long text without whitespace has no line transition at which to resume parsing.
It also makes the renderer allocate the same growing row on every key. Row reuse
keeps private cell storage between edits; `CopyContent` marks shared rows so an
edit copies their cells before writing. Existing snapshots retain both their
buffer contents and their wrapped view contents.

Periodic parser checkpoints then bound the graphemes parsed again within an
unbroken word. They are spaced at least 128 bytes apart where no breakable
whitespace remains. URL-style trailers can depend on text ahead of the parser,
so checkpoints are omitted while that lookahead-dependent match is active.

Median results for exactly 4,096 ASCII characters, with layout and view updates
on every key:

| Stage | Input time | Cumulative allocation |
| --- | ---: | ---: |
| Before row reuse | 370.9 ms | 629.7 MB |
| Reuse private rows | 226.5 ms | 11.3 MB |
| Add periodic checkpoints | 48.5 ms | 11.3 MB |

The combined speedup is 7.7×; cumulative allocation falls by 55.7×. Each stage
has three samples. The baseline sources come from `86869877d`, compiled through
a Go overlay; the row-reuse stage adds `40e373369`, and periodic checkpoints are
in `386cf4fd3`. These results concern the complete input sequence, not one key.
The final spaced-paragraph and newline-separated 4 KB benchmarks take 17.6 ms
and 14.7 ms respectively. The 100 KB bulk-paste benchmark remains near 3 ms.
The [unbroken-input chart](Performance-editor-unbroken.html) shows the staged
benchmark and matched live CPU results.

The 4,096-byte unbroken fixture retains 2,664 bytes of checkpoint descriptors and
204,768 bytes of pending-line scratch storage. The scratch storage already
exists in the checkpoint parser; periodic checkpoints add restart descriptors.
Clearing the text releases both arrays.

A matched live paste of 4,116 bytes in a 100 × 30 terminal reduces sampled app
CPU from 0.91 seconds to 0.24 seconds over a 12-second profile. Terminal captures
are identical. The final marker is observed at 410 ms before and 104 ms after,
with observations 100 ms apart. This single profile pair supports a CPU
reduction, not a precise completion-latency ratio.

Tests cover 3,200 replacements across ordinary and unbroken Unicode text,
footnotes, and trailers, plus targeted changes to trailer URL schemes. A
temporary independent oracle compares the incremental layouts with the original
parser. Build, unit tests, lint, integration tests, and editor race checks pass.
A broader race run exposed a baseline terminal suspend/resume race. The input
loop could exit while a terminal `Read` was still running, allowing resume to
replace state that the old reader could still access. A separate demonstration
and fix (`82dd68c27`, `fa9ab54a5`) make shutdown wait for outstanding reads after
stopping the terminal. The full gocui race suite and repeated suspend/resume
checks now pass. The Windows test target and Plan 9 tcell package compile; the
gocui Plan 9 target has pre-existing unsupported Unix syscall references.

## Screen cursor accounting

Replaying an unchanged editor prefix also updates the terminal's screen cursor.
The column loop now counts complete rows arithmetically while preserving the
pending wrap at the rightmost column. The change is in `15c65f224`.

| Columns written | Before | After |
| --- | ---: | ---: |
| 1 | 0.806 ns | 0.885 ns |
| 80 | 60.83 ns | 0.913 ns |
| 4,096 | 4,357 ns | 2.067 ns |
| 100,000 | 56,745 ns | 2.124 ns |

The 4,096-column operation improves by about 2,100×. This ratio applies to cursor
accounting: the matched full 4 KB unbroken-input benchmark improves from
50.3 ms to 44.3 ms, about 1.13×. Allocation is unchanged. All results have three
samples. The [unbroken-input chart](Performance-editor-unbroken.html) includes
this narrower comparison.

An independent temporary oracle compares 137,286 cursor states with the original
column loop. Focused tests cover pending wraps, multi-row writes, disabled
tracking, and empty writes. Build, unit tests, lint, integration tests, and the
full gocui race suite pass.

A separate reverse-byte-search prototype for trailer boundaries gives 2.8–4.9×
across four fixtures. It is not included in this change.

## One cell buffer for completed and pending text

`8913c16a6` keeps completed lines and the pending word in one cell buffer. A
checkpoint resumes directly in that buffer and retains coordinates for the
unchanged pending prefix. Soft wrapping inserts a break before the unfinished
word. The separate pending-word array is removed.

The edit benchmark starts with unbroken text already present. Each iteration
inserts one character, backspaces it, and updates the layout and view after
each key. It does not type the initial text one character at a time.

| Initial text | Before per pair | After per pair | Speedup |
| --- | ---: | ---: | ---: |
| 4,096 bytes | 40.67 µs | 17.64 µs | 2.31× |
| 100,000 bytes | 1.267 ms | 0.256 ms | 4.95× |
| 1,000,000 bytes | 23.66 ms | 2.779 ms | 8.51× |

Full 4 KB unbroken typing improves from 45.2 ms to 21.0 ms, about 2.15×. Spaced
paragraph input, newline-separated input, and 100 KB bulk paste remain near
their matched baselines. All measurements have three samples. The baseline
uses `b6e79411b` through a Go overlay.

For a 4,096-byte unbroken word, parser cell storage stays at 204,768 bytes and
checkpoint storage at 2,664 bytes; the additional 204,768-byte pending-word
array disappears. At 1,000,000 bytes, the removed pending array held 54,214,656
bytes. These figures count parser backing arrays, excluding strings and
rendered view cells. Clearing text releases the cells and checkpoints.

A matched live 4,116-byte paste records 0.23 app CPU seconds before and
0.19 seconds after over 12-second profiles. Both final-marker observations occur
at 104 ms, with observations 100 ms apart. Terminal captures are identical.
This small-input live profile does not measure the large-text edit workload.
The [single-buffer chart](Performance-editor-single-buffer.html) distinguishes
the edit and full-input comparisons.

The original-parser differential checks, unit tests, integration tests, build,
lint, and full gocui race suite pass. The user chose a separate performance
commit for this change, which spans the earlier parser and checkpoint storage.

## Remaining scans and allocation costs

A fresh profile at `209c9135f` measures an insert and backspace in 1 MB of
existing unbroken text. Trailer-boundary detection accounts for 64.7% of CPU
samples; the renderer's control-character check accounts for 13.8%. Both scan
the full text after each edit. The checkpoint parser itself uses little CPU.
These percentages include background GC and describe this benchmark only.

A prototype replaces the reverse substring search for two newlines with repeated
reverse byte searches. Across three samples per fixture, edit pairs improve by
1.58× at 4 KB, 2.60× at 100 KB, and 2.62× at 1 MB. The prototype is not shipped
because these whole-operation gains fall below the 5× target. Reusing the known
unchanged prefix for trailer and control-character detection led to the
incremental scan change below.

Independent profiles identify allocation costs in two other paths. Searching
2,000 lines of 200 `a` characters for `a` allocates about 76.3 MB per rebuild,
including temporary per-line results copied into the combined result list.
Writing 2,000 colored diff lines allocates about 11.7 MB, mostly while growing
cell arrays. These are repeated synthetic workloads, not live latency results.
The search allocation experiments below measure two ways to reduce that cost.

Building a tree of 10,000 commit files takes about 3.13 ms in one profiled
benchmark run. CPU cost is spread across path splitting and joining, sorting,
allocation, and GC. A separate directory-removal refresh takes about 3.73 ms;
its CPU profile also includes fixture setup while the benchmark timer is
stopped, so it cannot isolate refresh CPU cost.

The `remainingHotspotRanking` section of the
[profiling data](performance-live-ranking.json) records the commands, source
revision, profile summaries, and limits. Raw logs, overlays, and profiles remain
locally under `.git/perf-investigation/resume-*`.

## Incremental trailer and control scans

`40c04419c` separates blank-line detection from trailing-newline handling.
`9e2109e18` then uses the earliest edit to bound trailer searches and preserves
the first control character when it precedes that edit. Searches include
overlapping separators and trailing blank lines that become interior text.
The renderer uses the cached control position to select its writer.

| Edit pair in existing unbroken text | Before | After | Speedup |
| --- | ---: | ---: | ---: |
| 4,096 bytes | 17.88 µs | 7.97 µs | 2.24× |
| 100,000 bytes | 258.32 µs | 12.82 µs | 20.15× |
| 1,000,000 bytes | 2.540 ms | 92.85 µs | 27.35× |

Each pair inserts one character and backspaces it, updating layout and view
after each key. Initial text creation is outside the timer. Full 4 KB typing
improves by 3.40× for newline-separated input, 2.44× for a spaced paragraph,
and 1.97× for unbroken text. Bulk paste stays near its baseline. All fifteen
comparisons have three samples and use `722c56cef` through a Go overlay as
the baseline. The [scan charts](Performance-editor-scans.html) distinguish edit
pairs from full input sequences.

The cache adds one integer to the existing layout, with no additional arrays
or retained strings. Allocation for the larger edit pairs remains one growing
content string per pair. A fresh 1 MB edit profile is dominated by copying and
garbage collection; it is a synthetic benchmark profile, not a matched live
application measurement.

Independent full scans agree through 534,912 exhaustive replacements across
3,906 short strings and 6,000 generated batches of deferred edits. Tests compare
rendered content with the normal writer as controls are inserted and removed,
and after clears and width changes. The original-parser oracle, build, unit
tests, lint, full integration suite, and full gocui race suite pass. The user
chose to land this as a separate performance commit after the helper preparation.

## Search allocation experiments

Appending matches directly into the combined result list removes temporary
per-line arrays in both the model and view search paths. It improves the
400,000-match fixture by 1.35× and ordinary matching text by 1.08×. This prototype
preserves the existing result-buffer lifetime and is not shipped.

The user then chose to measure retaining that combined buffer while searching,
releasing it on `ClearSearch` or an empty query. With rendered layout and the
initial search warmed, repeated many-match queries improve from 13.00 ms to
2.96 ms, or 4.39×. Alternating `a`, `aa`, `aaa`, and `aaaa` queries improves
from 13.75 ms to 3.56 ms, or 3.86×. Clearing search on every iteration leaves
less capacity to reuse and improves from 13.26 ms to 8.70 ms, or 1.52×.

The fixture retains 11,345,904 bytes of result capacity for 400,000 matches.
Tests verify reuse across a query change, retention through a missing query,
and release on both clear paths. Existing search tests pass. Two scanner
variants, reading cells by address and replacing column-ring modulo operations
with a cyclic index, give no substantial additional improvement.

All timed comparisons have three samples. The baseline is `9e2109e18`, and the
prototype runs through a Go overlay. The allocation reduction is substantial,
but none of these timing results meets the 5× target, so these prototypes remain
outside the application. Their samples and limits are preserved in the
measurement and profiling data. No new live application speedup is claimed.

## Search result storage follow-up

The direct-append and reusable-result experiments now ship in `c7744f620`
and `1784ee790`. This supersedes the earlier decision to keep them outside
the application. Both model and rendered-view searches append into the same
result buffer. Query changes and content refreshes reuse its capacity;
`ClearSearch` and an empty query release it.

Fresh measurements against `643a94c3a` put the warmed 400,000-match rebuild at
15.92 ms before, 11.68 ms with direct appends, and 2.91 ms with reuse. Allocation
falls from 76,322,208 bytes and 18,030 allocations to 32 bytes and 3 allocations
per rebuild. Alternating queries improves from 16.40 ms to 3.41 ms. Clearing
search after each changing query improves from 16.81 ms to 12.15 ms.

The many-match fixture retains 11,345,904 bytes of result capacity during an
active search. Smaller or missing queries and content changes keep that
capacity until search is cleared. Initial result growth is outside the warmed
timings. These are synthetic benchmark results, not measured live latency.

The [search storage report](Performance-search-results.html) shows all search
cases. [Its data](performance-search-results.json) records five samples per
case, source revisions, commands, fixture details, and limits. The new query
lifetime benchmark runs against both implementations with the same source;
a Go overlay supplies the baseline application files.

Build, unit tests, lint, the gocui race suite, and selected search integration
tests pass. Unit tests run with `NO_COLOR` unset because the context-mode runner
sets it and presentation tests assert ANSI color output.

## Selection and graph allocation follow-up

Selection recovery spent most of its sampled CPU time splitting and rejoining
paths that nodes already stored. `e47b40cef` uses the stored path and parent
prefixes for compressed-directory aliases. The 5,000-file single-removal
fixture improves from 428.22 µs to 106.62 µs, about 4×, while allocation falls
from 320,160 bytes to 80,048 bytes. Recovery when many files disappear improves
from 2.03 ms to 0.71 ms. Rename handling and node ordering are unchanged.

Graph rendering now owns each row's cells in one slice in `d48c4dc0a`, and
`1ac1a7ef2` replaces three column maps with per-column flags. The existing
50-commit fixture improves from 87.10 µs to 56.58 µs. It allocates 109,576 bytes
in 585 allocations, down from 131,296 bytes in 1,139 allocations. New fixtures
cover both symbol sets and a real selected commit at 50 and 500 commits. Their
combined gains range from 1.42× to 1.99×. The original fixture also improves
with four CPUs, from 84.12 µs to 53.68 µs.

The column array uses memory proportional to the incoming graph's width and
the number of new parents. A sparse, wide input can use more flag storage than
the maps. Sparse-column and many-parent tests pass against both implementations.

The [selection and graph report](Performance-selection-graph.html) and
[its data](performance-selection-graph.json) preserve all five samples per
case, source revisions, and commands. These are synthetic measurements;
selection timings exclude tree reconstruction. Build, unit tests, lint,
file-tree and graph race tests, rename and range-selection integration tests,
and the commit-graph integration test pass.

The full fork-base comparison was rerun at `2769d840b`. All 90 cases have five
matched samples per revision and faster medians than the fork base. The
remaining weakest cases are 100-file tree construction at about 10% faster,
and writing metadata and colored output at about 13–17% faster. Writing takes
milliseconds per fixture, while the small tree build takes about 32 µs, so the
writer remains the larger cost to investigate. The refreshed
[summary](../../README.md#performance) and [data](performance-comparison.json)
include the new query-lifetime and graph variants.

## Printable-ASCII writer batches

`a9abf0957` shares destination-cell reservation between cell copying and direct
producers. `309851b54` batches printable ASCII into those reserved cells when
the escape interpreter is idle. Escapes, controls, and Unicode retain the
grapheme decoder, including the last ASCII byte before Unicode that could
extend its grapheme. Short initial reservations leave room for following cells
on mixed-text lines.

Against `afa2c8215`, the 2,000-line fixtures improve by 3.70× for plain text,
1.67× for colored text, 1.54× for mixed Unicode, and 1.59× for metadata. All four
allocate fewer bytes. The searched-view write improves by 2.87× and the long
line by 5.18×. These are fresh-view synthetic benchmarks with five samples per
case, not measured live latency.

Cell capacity stays the same or falls for all five measured line fixtures.
A temporary differential harness compares 2,400 snapshots from 300 deterministic
mixed-input streams against the pre-batch writer. Cells, styles, hyperlinks,
metadata, text as written, and cursor state match. Tests also compare normal
writes one grapheme at a time, including Unicode, ANSI, metadata, carriage
returns, tabs, cursor motion, and soft wrapping. The shared CRLF correction at
the foundation also makes the stored text consistent at line-write boundaries.
All 2,400 differential snapshots match after replaying the batching change on
that corrected foundation.

An editor smoke check showed a small paragraph-input timing difference. A
longer seven-sample recheck shows overlapping timings and no material slowdown;
no editor speedup is claimed. Build, unit tests, lint, the gocui race suite, and
the file, diff, main-view, and search integration groups pass.

The [writer batch report](Performance-writer-batches.html) and
[its data](performance-writer-batches.json) preserve the samples, source
revisions, commands, retained-capacity measurements, and limits. The initial
exact-capacity prototype increased mixed-Unicode allocation; it is not shipped.

## Writer capacity probe

A benchmark-only probe pre-sizes all 2,000 destination lines and their cell
arrays to the exact size of the repeated fixture. A sample line provides that
size before timing; destination allocation remains inside the timed loop.
The production decoder is unchanged.

The probe improves plain text by 1.89×, colored text by 1.75×, Unicode by 2.06×,
and metadata-bearing text by 1.68×. Each comparison has three samples against
`22eb6ebd1`. Even with exact fixture knowledge, avoiding array growth alone
does not meet the 5× target. No preallocation strategy is added to the writer.
The samples are included in the measurement data, with the fixture and limits
under `writerCapacityProbe` in the profiling data.

Reproduce the committed benchmarks with:

```sh
go test ./pkg/gui/context -run='^$' -bench=BenchmarkFilteredListNavigation -benchmem -benchtime=200ms -count=3
go test ./pkg/gui/filetree -run='^$' -bench='BenchmarkFileTree(RangeSelection|Navigation)' -benchmem -benchtime=200ms -count=3
go test ./pkg/gocui -run='^$' -bench='Benchmark(SearchRepeatedPrefix|SearchOrdinaryText|NearestSearchPosition|ScrollSearchPosition)' -benchmem -benchtime=150ms -count=3
go test ./pkg/gocui ./pkg/gui -run='^$' -bench='Benchmark(EditorInput|CommitDescriptionInput)' -benchmem -benchtime=300ms -count=3
go test ./pkg/gocui -run='^$' -bench='Benchmark(EditorParagraphInput|EditorUnbrokenInput|PasteText)' -benchmem -benchtime=150ms -count=3
go test ./pkg/gocui -run='^$' -bench=BenchmarkNotifyCellsWritten -benchmem -benchtime=150ms -count=3
go test ./pkg/gocui -run='^$' -bench=BenchmarkEditorEditUnbrokenText -benchmem -benchtime=150ms -count=3
```

## README comparison after the independent follow-ups

The initial follow-up comparison measured `3c989658a2e5` against the fixed fork base
`8e3ac76991f5`. It retains all 103 existing cases and adds 66 cases for editor
navigation, small tree selections, flat construction and refresh, graph
columns, cached graph viewports, and streamed search. The
[raw comparison](performance-comparison.json) contains five samples per side
and case, measured for 200 ms per sample with the pinned settings.

Cached graph viewports also run with GOMAXPROCS=4 and 16. Those rows state
their CPU setting; the other rows use GOMAXPROCS=1. That run retained slower medians, including navigation near the beginning of
a long editor buffer and some linear graph viewports. These operation timings do not measure
end-to-end UI latency.

The runner extracts benchmark-only fixtures from files whose names do not
contain `bench`. It copies identical fixture text into both measurement
checkouts and records source hashes. Public construction entry points bridge
the ordering-interface change between revisions. The invalidated-visibility
case toggles a directory twice before selecting a row; those operations are
included in its time. Streamed search reads status after each batch. Large
editor fixtures seed their content outside the timed operation.

Reproduce this comparison with:

```sh
just perf-compare 8e3ac76991f5 3c989658a2e5
```

Check fixture compilation without measuring with `--prepare-only`. For a
bounded run, measure one package and then collect the complete comparison:

```sh
just perf-compare 8e3ac76991f5 3c989658a2e5 --prepare-only
just perf-compare 8e3ac76991f5 3c989658a2e5 --resume --package=./pkg/gui/filetree
just perf-compare 8e3ac76991f5 3c989658a2e5 --resume
```

A package-only run saves samples without changing the README. Resume reuses
only checkpoints whose revisions, Go sources, toolchain, CPU, sampling
settings, command arguments, and output hashes match. Missing or invalid
checkpoints are measured again. The final collection verifies matching case
sets and complete repetitions before writing the README and timing images.

The [independent implementation results](Independent-performance-follow-ups.md)
retain the immediate before-and-after measurements for those changes. Their
baselines and some fixtures differ from the broader README comparison, so
the two sets of timings should be interpreted separately.

## Audit against the previous README snapshot

The [snapshot regression audit](Snapshot-regression-audit.md) compares the
previous snapshot `818a17b1d` with the corrected checkpoint `b920143e7`. The
README now includes 18 full commit-viewport cases, bringing its fixed-base
comparison to 187 cases. These include branch-marker work outside graph
rendering. Early returns remove unused ancestry work; the visible-marker,
unmerged-main case still rebuilds its history map and is awaiting the
model-lifetime design decision. Its slower measurements remain visible.

The graph and editor fixes, full samples, review scope, remaining search
lookup cost, and validation results are recorded in that audit.
