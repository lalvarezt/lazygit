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

The [measurement data](performance-follow-up.json) contains all 106 comparisons,
including individual samples, allocation counts, smaller workloads, ordinary
search text, file-tree refresh costs, and unshipped search prototypes.

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
