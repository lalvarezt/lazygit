# Delta ANSI ingestion follow-up

Delta's upstream integration is complete. Its
[integration report](../../../delta/etc/performance/upstream-integration-2026-11-12.md)
records the accepted costs of corrected output. This note records the lazygit
follow-up and fresh measurements from 2026-11-12.

The implementation is committed as `011f797f8`, a reviewable fixup for the
original SGR fast path. It passes the correctness checks below and improves
Delta ingestion. The initial prototype's confirmed combining-mark slowdown
does not reproduce after tuning. The remaining measurement uncertainty does
not establish zero overhead.

## Why the delta update exposes more work

Delta commit `876a162` preserves feature-defined line-number formats when
`--side-by-side` is supplied on the command line. The personal Catppuccin
configuration consequently emits its intended gutter styles, including longer
RGB color escapes adjacent to Unicode box characters.

With lazygit pinned to `eb0320973ffe795970ad95c047ad147a15c53e81`, fresh true-color
view population plus diff identity extraction measured:

| Input | Previous delta output | Corrected delta output |
| --- | ---: | ---: |
| 1,000-line Go change | 33.43 ms | 43.14 ms |
| 12,000-line Go change | 441.14 ms | 578.70 ms |

These are consumer component timings, excluding delta execution, physical
redraw, and user actions. They do not establish complete interactive latency.
For the larger input, output grows from 15,016,960 to 19,125,954 bytes and from
743,586 to 839,144 SGR sequences. Some extra work is expected, but the current
parser also takes an avoidable slow path.

The delta deleted-file syntax-highlighting cost is a separate renderer issue.
Changing lazygit's ANSI handling will not remove that parsing cost.

## Unicode boundary check

Focus on `escapeInterpreter.parseSGR` in `pkg/gocui/escape.go`. Its fast path
rejects a completed SGR escape whenever a non-ASCII byte follows the final `m`.
Ordinary gutter box characters therefore send complete RGB escapes through
the streaming grapheme path one character at a time.

That guard protects a real invariant: a following combining rune can extend
the grapheme containing `m`. Preserve that behavior while recognizing complete
escapes before ordinary Unicode characters.

The boundary check decodes the following rune once. It rejects incomplete
UTF-8, Unicode marks, and ZWJ before calling the grapheme engine. Common
combining diacritics use a range check that avoids the Unicode property lookup.
Other runes go through `uniseg.FirstGraphemeCluster` with just `m` and that rune.
The fast path accepts the escape when `m` is a separate grapheme. The change
introduces no cache.

The [reviewable prototype patch](../../../delta/etc/performance/data/lazygit-sgr-unicode-boundary-prototype.patch)
is in the sibling delta repository. It records the initial check before the
fallback tuning. Commit `44945495c` introduced the fast path and was the last
commit to change `pkg/gocui/escape.go` at task start. Fixup `011f797f8` refines
that implementation and adds direct test coverage.

Historical prototype checks passed the gocui view tests and selected snapshots
captured from the original consumer at chunk boundaries. On that prototype,
the larger old and corrected streams measured 387.69 and 387.33 ms. The smaller
case still measured 30.76 and 32.21 ms. These are exploratory observations with
incomplete coverage, not proof that the regression is resolved or a direct
paired comparison against the original implementation.

## Final paired measurements

The baseline is the task-start revision,
`eb0320973ffe795970ad95c047ad147a15c53e81`. Baseline and candidate test
executables use Go 1.27.4 with normal optimization. Each sample includes a fresh
true-color view, `SetContent`, `DiffLineContents`, and `parseDiffLineIdentities`,
at identical view dimensions. It excludes Delta execution, physical redraw,
and user actions.

Both executables consume the same frozen stream in each comparison. Runs use
CPU 2, `GOMAXPROCS=1`, `GOGC=100`, `GOMEMLIMIT=off`, and forced color, with
`NO_COLOR` unset. Builds and tests are idle during timing. Two warmup pairs
precede 15 serial pairs with alternating execution order. The governor is
`powersave`; CPU frequency is not locked.

These values are medians in milliseconds:

| Input stream | Original parser | Tuned parser |
| --- | ---: | ---: |
| 1,000-line Go, previous Delta output | 35.44 | 31.88 |
| 1,000-line Go, corrected Delta output | 47.35 | 34.39 |
| 12,000-line Go, previous Delta output | 487.93 | 442.35 |
| 12,000-line Go, corrected Delta output | 599.17 | 458.02 |

On the large corrected stream, allocated bytes per operation fall from
713,661,256 to 674,253,856, and allocations fall from 5,023,819 to 2,967,712.
Sampled process peaks vary between runs. A subsequent 21-pair check uses Linux
`wait4` resource usage to capture exact process peaks, with medians of 551,200
and 557,708 KiB. The paired median change is -1.07%, with a bootstrap interval
of -2.74% to +1.82%. This does not establish a consistent peak-memory benefit.
Allocated bytes are cumulative allocation volume, not peak resident memory.

Initial prototype profiles show less time in `parseOne`. SGR parameter
splitting remains visible, but the profile does not establish that another
optimization is warranted. The corrected stream still takes more work than
the previous stream with the tuned parser.

## Fallback tuning

`BenchmarkViewWriteSGR` covers repeated and distinct RGB sequences followed by
ASCII, ordinary Unicode, or a combining mark, with one-line and 2,000-line
updates. Existing `BenchmarkViewWrite` cases cover plain text, colored text,
Unicode without ANSI, and metadata.

An initial 11-pair matrix found a possible combining-mark slowdown. A separate
31-pair recheck confirmed it with three warmup pairs and 200 ms samples:

| Fallback input | Original median | Initial prototype median | Paired median change |
| --- | ---: | ---: | ---: |
| One line | 2,771 ns | 2,819 ns | +1.90% |
| 2,000 lines | 5.126 ms | 5.281 ms | +3.22% |

The paired bootstrap intervals exclude zero for both initial fallback cases.
This stress case puts a combining mark immediately after the SGR terminator.
The initial prototype checks the grapheme boundary before falling back to
streaming interpretation.

The user chose further tuning. Decoding once and rejecting marks early reduces
the cost. The range check for common combining diacritics avoids the property
lookup on that path. The final fallback comparisons are:

| Fallback input | Original median | Tuned median | Paired median change | Bootstrap interval |
| --- | ---: | ---: | ---: | ---: |
| One line, 31 pairs at 200 ms | 2,776 ns | 2,787 ns | +0.40% | -0.22% to +0.65% |
| 2,000 lines, 31 pairs at 500 ms | 5.159 ms | 5.193 ms | +1.23% | -0.40% to +2.05% |

The original confirmed loss no longer reproduces, but these intervals do not
prove zero overhead. Fallback allocation counts remain unchanged. Other
fallback classes have correctness coverage without dedicated performance
rechecks.

The final short matrix also flags plain text and Unicode without ANSI. Longer
31-pair, 500 ms rechecks do not reproduce those slowdowns. Their paired median
changes are -0.78% and -0.11%, and both intervals include zero. One-line ASCII
SGR updates show a sub-1% paired difference in the 31-pair check, about 4 ns
between medians. No material common-path loss was established. All original
samples remain in the evidence.

## Correctness checks completed

Direct tests verify fast-path acceptance and complete interpreter state against
the streaming interpreter in normal, 256-color, and true-color modes. They
cover box characters, wide characters, emoji, flags, prepend characters,
combining marks, variation selectors, ZWJ sequences, spacing marks, malformed
UTF-8, incomplete UTF-8, resets, and extended colors.

An exhaustive probe checks all 1,112,064 valid Unicode scalars against the
existing grapheme engine. It finds no unsafe fast-path acceptance. There are
124 conservative mark rejections where the engine would allow the fast path.
These keep the original streaming behavior.

The preserved original executable and candidate match in 33,009 synthetic
chunk cases. Each comparison includes cells, attributes, links, row metadata,
cursor position, parser state, and continuation output. Cases split escape
sequences and UTF-8 at every byte boundary, including OSC 8 and OSC 1717.

Seven real renderer streams also match in text and extracted row identities
across all three color modes, complete writes, 7-byte, 4,096-byte, and
65,536-byte chunks, and content replacement. These include previous and
corrected Go output and the opt-in hunk layout.

The candidate passes `just format`, `just build`, `just unit-test`, `just lint`,
and race tests for `pkg/gocui` and `pkg/gui/controllers/helpers`. Selected
integration tests cover metadata, copying, and staging. An additional overlay
test runs the actual corrected Delta renderer through copying original source
text, editing its source line, and staging. Both parsers pass that test with
Delta using the pane's `{{width}}`.

An initial unit run without forced color failed color assertions. The required
color controls resolve those failures. An initial action probe used a fixed
renderer width larger than the pane, causing wrapping that invalidated the
test's navigation. Matching the pane width resolves that probe for both
parsers. Neither failure establishes a parser regression.

Windows, physical terminal redraw, complete key-to-pixel latency, and
cancellation through the actual renderer were not remeasured. The unit suite
covers existing offscreen view behavior. No executable was installed or pushed.

## Further work if profiling still justifies it

Measure allocations and CPU in SGR parameter splitting, numeric parsing, and
attribute application after the boundary fix. A scanner that handles common
numeric parameters without temporary strings may help. Keep one interpretation
of reset, foreground, background, and attribute semantics across callers.

Treat SGR caching as a later option. Earlier overlays caching complete incoming
styles, style operations, and validated sequences gave partial improvements
with additional state. None was adopted. Any cache needs a firm size bound and
must preserve incoming attribute state and terminal color-mode behavior. A
sequence-to-final-style cache without those dependencies is unsafe.

## Validation before adopting a change

Follow [the repository validation profile](Regression-validation.md). Preserve
an executable or reproducible snapshot of the original implementation first.

- Compare original and candidate consumers on both the old and corrected delta
  streams. Use serial alternating pairs with builds and tests idle, true color
  forced, and identical view dimensions. Retain raw samples and output hashes.
- Exercise the complete view population and identity-extraction caller. Also
  measure ordinary ASCII, Unicode without ANSI, many unique SGR sequences,
  and short updates so an optimization cannot hide a new common-path cost.
- Split escape sequences and UTF-8 at every relevant byte boundary. Cover
  combining marks, variation selectors, ZWJ sequences, malformed UTF-8,
  incomplete CSI, resets, extended colors, OSC 8, and OSC 1717 records. Compare
  cells, attributes, links, row identities, and final parser state against the
  original. Account for any new cache bookkeeping separately.
- Cover append, replacement, cancellation, and terminal color-mode changes.
  Run relevant unit, race, and integration checks. Verify copying original
  source text, editing its source line, and staging through the deployed stack.
- Measure allocations and memory as well as completion and first-visible-body
  timings. Avoid adding component medians from unrelated runs.

Retain the corrected colors and metadata. Disabling those features would change
the behavior this optimization is intended to preserve.

## Evidence and reproduction

The [final raw measurements](data/delta-ansi-ingestion-tuned-2026-11-12.json)
retain all 384 final matrix samples, tuning experiments and rechecks, memory
samples, controls, source and executable hashes, commands, and correctness
results. They record the implementation commit and the identity of the local
bundle at `.git/delta-ansi-ingestion-tuned-2026-11-12.tar.gz`, relative to the
repository root. That bundle contains preserved test executables, source
overlays, fixtures, snapshot comparisons, probes, and logs.

The working extraction is `.git/delta-ansi-tuning/`. The `compare.cjs` script
replays tuning comparisons. Its arguments select the output label, pair count,
sample duration, benchmark pattern, and executable names. The final directory
contains `measure.cjs` and `memory.cjs`, preserved executables, and reproduction
overlays. `rssprobe.go` captures exact child-process resource usage. Synthetic
snapshot tests write `DELTA_ANSI_SNAPSHOTS`; consumer probes read
`DELTA_ANSI_FIXTURE`.

The [initial prototype evidence](data/delta-ansi-ingestion-2026-11-12.json)
retains its 384 matrix samples, 248 fallback recheck samples, profiles, and
bundle identity. Its working extraction is `.git/delta-ansi-follow-up/`.

The sibling delta repository holds the [historical regression report](../../../delta/etc/performance/upstream-benchmarks-2026-11-11.md)
and [raw evidence with exploratory scripts](../../../delta/etc/performance/data/upstream-benchmarks-2026-11-11.json).
The local raw bundle is at
`../delta/.git/performance/delta-upstream-2026-11-11.tar.gz`, relative to the
lazygit repository root. It includes frozen fixtures, executable copies,
exploratory sources, and logs. Its identity is recorded in the evidence JSON.
