# Delta ANSI ingestion follow-up

Delta's upstream integration is complete. Its
[integration report](../../../delta/etc/performance/upstream-integration-2026-11-12.md)
records the accepted costs of corrected output. This note records the lazygit
follow-up and fresh measurements from 2026-11-12.

The boundary prototype passes the correctness checks below and improves Delta
ingestion. It also adds a reproducible cost to the combining-mark fallback.
The parser change remains uncommitted pending a decision about that tradeoff.
The benchmark coverage and investigation evidence are committed separately.

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

## Boundary prototype

Focus on `escapeInterpreter.parseSGR` in `pkg/gocui/escape.go`. Its fast path
rejects a completed SGR escape whenever a non-ASCII byte follows the final `m`.
Ordinary gutter box characters therefore send complete RGB escapes through
the streaming grapheme path one character at a time.

That guard protects a real invariant: a following combining rune can extend
the grapheme containing `m`. Preserve that behavior while recognizing complete
escapes before ordinary Unicode characters.

The existing boundary-only prototype checks that the next rune is complete,
then applies `uniseg.FirstGraphemeCluster` to just `m` and that rune. It accepts
the fast path when `m` is a separate grapheme and keeps the streaming fallback
for incomplete UTF-8 or an extending grapheme. It introduces no cache.

The [reviewable prototype patch](../../../delta/etc/performance/data/lazygit-sgr-unicode-boundary-prototype.patch)
is in the sibling delta repository. The working lazygit candidate applies that
boundary check, corrects its function comment, and adds direct test coverage.
Commit `44945495c` introduced the fast path and is still the last commit to
change `pkg/gocui/escape.go`. Its refinement belongs in a separate reviewable
fixup for that commit.

Historical prototype checks passed the gocui view tests and selected snapshots
captured from the original consumer at chunk boundaries. On that prototype,
the larger old and corrected streams measured 387.69 and 387.33 ms. The smaller
case still measured 30.76 and 32.21 ms. These are exploratory observations with
incomplete coverage, not proof that the regression is resolved or a direct
paired comparison against the original implementation.

## Fresh paired measurements

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

| Input stream | Original parser | Boundary prototype |
| --- | ---: | ---: |
| 1,000-line Go, previous Delta output | 35.90 | 32.39 |
| 1,000-line Go, corrected Delta output | 46.56 | 32.83 |
| 12,000-line Go, previous Delta output | 474.29 | 427.65 |
| 12,000-line Go, corrected Delta output | 601.02 | 446.10 |

On the large corrected stream, allocated bytes per operation fall from
713,661,256 to 674,253,856, and allocations fall from 5,023,819 to 2,967,712.
Seven alternating memory probes show median sampled process peaks of 567,348
and 535,600 KiB. These sample Linux `VmHWM` every 2 ms and can miss a final peak.
Allocated bytes are cumulative allocation volume, not peak resident memory.

CPU profiles confirm less time in `parseOne`. SGR parameter splitting remains
visible after the change, but the profile does not establish that another
optimization is warranted. The corrected stream still takes more work than
the previous stream with the candidate parser.

## Confirmed fallback tradeoff

`BenchmarkViewWriteSGR` covers repeated and distinct RGB sequences followed by
ASCII, ordinary Unicode, or a combining mark, with one-line and 2,000-line
updates. Existing `BenchmarkViewWrite` cases cover plain text, colored text,
Unicode without ANSI, and metadata.

An initial 11-pair matrix found a possible combining-mark slowdown. A separate
31-pair recheck uses three warmup pairs and 200 ms benchmark samples:

| Fallback input | Original median | Prototype median | Paired median change |
| --- | ---: | ---: | ---: |
| One line | 2,771 ns | 2,819 ns | +1.90% |
| 2,000 lines | 5.126 ms | 5.281 ms | +3.22% |

The deterministic paired bootstrap intervals exclude zero for both fallback
cases. ASCII rechecks do not show a consistent slowdown. Allocation counts
remain unchanged on the fallback. This stress case puts a combining mark
immediately after the SGR terminator, where the candidate checks the boundary
before falling back to streaming interpretation.

The decision is whether to reduce this extra boundary-check cost before
adoption or explicitly accept it in exchange for the measured ingestion gains.
The current recommendation is to tune the boundary check first.

## Correctness checks completed

Direct tests verify fast-path acceptance and complete interpreter state against
the streaming interpreter in normal, 256-color, and true-color modes. They
cover box characters, wide characters, emoji, flags, prepend characters,
combining marks, variation selectors, ZWJ sequences, spacing marks, malformed
UTF-8, incomplete UTF-8, resets, and extended colors.

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

The [fresh raw measurements](data/delta-ansi-ingestion-2026-11-12.json) retain all
384 initial benchmark samples, 248 recheck samples, memory samples, controls,
fixture and executable hashes, commands, and correctness results. They also
record the identity of the local bundle at
`.git/delta-ansi-ingestion-2026-11-12.tar.gz`, relative to the repository root.
That bundle contains preserved test executables, source overlays, fixtures,
snapshot comparisons, probes, profiles, and logs.

The working extraction is `.git/delta-ansi-follow-up/`. Its `measure.cjs`,
`recheck.cjs`, and `memory.cjs` replay the measurements with the preserved test
executables. The overlay JSON files reproduce the test harness builds against
the pinned source. The synthetic snapshot test writes its output to
`DELTA_ANSI_SNAPSHOTS`; the consumer probe reads `DELTA_ANSI_FIXTURE`.

The sibling delta repository holds the [historical regression report](../../../delta/etc/performance/upstream-benchmarks-2026-11-11.md)
and [raw evidence with exploratory scripts](../../../delta/etc/performance/data/upstream-benchmarks-2026-11-11.json).
The local raw bundle is at
`../delta/.git/performance/delta-upstream-2026-11-11.tar.gz`, relative to the
lazygit repository root. It includes frozen fixtures, executable copies,
exploratory sources, and logs. Its identity is recorded in the evidence JSON.
