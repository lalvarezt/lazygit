# Why search opens before a renderer finishes

The 2026-11-08 Neovim UI profile measured the first large-diff body at about
35 ms, but its search prompt at 1.23 seconds. `MainViewController.openSearch`
waited for `ReadToEnd` before showing the prompt. The user chose immediate
typing, incremental results, and a visible loading state.

The prompt now opens synchronously. Full reads publish progress on a 33 ms
timer, including while the producer is stalled between lines. Progress cannot
overwrite the text field, EOF removes the loading label, and cancellation
cannot reopen the prompt. Enter before any match arrives records a pending
selection, so the first arriving match is selected. Clearing or cancelling
search clears that pending selection.

Progress makes prefix reuse necessary. Wrapping and searching the entire
existing buffer on each frame repeats work as the buffer grows. Gocui retains
wrapped rows and search results before the first changed row. Query changes,
width changes, overwritten rows, and replacement buffers invalidate the needed
prefix. Model-based search retains its full model path. Differential tests
compare streamed, overwritten, resized, and replaced content to a fresh search.

On 8,192 lines in 64 chunks, five-sample streaming-search medians improved from
85.11 to 13.35 ms unwrapped, and 89.67 to 16.74 ms wrapped. Allocated bytes were
essentially unchanged. Complete SGR sequences also use the existing color
interpreter without per-byte grapheme segmentation. Existing colored-Git and
colored-large-diff benchmarks improved by about 15% and 7% respectively.
Incomplete and unusual sequences keep the streaming parser's behavior.

The installed 120×40 Neovim UI, with the local optimized delta fork, opens the
large-diff search prompt in a five-sample median 1.96 ms. A match near old line
10,000 becomes visible 343.65 ms after typing starts; all output finishes at
1,065.35 ms on that clock. Prompt timing starts at slash, while query timing
starts after an intentional 80 ms interval. Do not combine those boundaries.
Normal medium/large first-body medians are 21.42/30.94 ms. These are redraw-flush
observations, not physical display timings or a universal renderer speedup.

The tests cover a producer that stalls before EOF, fast typing during slow
rendering, early confirmation, cancellation, full results, refresh, and selected
diff operations. The stalled-producer test fails on the old read loop. Unit,
lint, formatting, full `gocui` and `tasks` race suites, and the `filter_and_search`,
`main_view`, and `diff` integration groups pass. In an environment with
`NO_COLOR`, unset it and force color for the existing color-sensitive tests.

[The measurements](performance/incremental-search.json) retain the benchmark
samples and UI summaries. Full cross-layer evidence and Markdown validation
live in the chezmoi repository's `docs/nvim-stack-performance.md`. Delta's build
optimization and output-parity evidence live in that fork under
`etc/performance/`. Renderer completion remains real work; the prompt is usable
while it happens.
