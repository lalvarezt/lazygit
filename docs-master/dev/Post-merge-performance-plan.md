# Implement performance follow-ups after the upstream merge

Cached tree-range selection is implemented independently, as recorded in the
[implementation results](Independent-performance-follow-ups.md).
Prepared-query reuse remains deferred until
[upstream PR #6201](https://github.com/jesseduffield/lazygit/pull/6201) is merged
and integrated. Do not cherry-pick the open PR or replace our drawing
implementation early. The sections below retain the original plans.

The source inspection used fork revision `96e977c85`. Recheck the named
functions after integration because upstream may change their interfaces.
Keep the fork's cell-buffer batching, trailing-fill behavior, incremental
search, and editor improvements during integration.

## Reuse visible tree rows for range selection

### Establish the baseline

Inspect these modules and their callers:

- [`treeView.getNodes`](../../pkg/gui/filetree/tree_view.go) caches visible
  nodes and depths. It invalidates the cache when the root, collapsed-path
  object, or collapsed-path revision changes.
- [`selectedNodes`](../../pkg/gui/filetree/selected_nodes.go) starts a new
  depth-first traversal from the hidden root for each selected range.
- Both file-tree view models call that helper through `GetSelectedItems`.
  They hold typed tree interfaces, so they cannot directly access the private
  cache in `treeView`.
- [`BenchmarkFileTreeRangeSelection`](../../pkg/gui/filetree/range_bench_test.go)
  selects the entire tree. It does not isolate the cost of reaching a small
  selection near the bottom.

Add benchmark cases for both working-tree files and commit files. Vary the
total file count, selection size, and position independently. Include single
rows and short ranges at the beginning, middle, and end, plus full selection.
Keep fixture construction outside timed selection loops.

Measure a warm visible-row cache separately from the first selection after a
tree replacement or collapse change. Include flat and tree display modes,
compressed directories, collapsed subtrees, and the optional displayed root.
Report time, allocated bytes, and allocations. Do not infer an interactive
latency improvement from an isolated selection benchmark.

### Implement the range operation at the tree interface

1. Add a range operation that obtains the current visible nodes through
   `treeView.getNodes` and wraps only the requested slice.
2. Expose that operation through the typed file-tree interfaces. Keep cache
   ownership and invalidation inside `treeView`; do not expose its mutable
   backing slice to controllers or view models.
3. Route both `GetSelectedItems` implementations through the same range
   implementation. Preserve the selected start and end indices returned to
   callers.
4. Remove the redundant traversal once both panels use the new operation.

Prefer one shared range implementation with typed wrappers for `FileNode` and
`CommitFileNode`. Resolve the exact interface shape after checking upstream's
merged tree code. Do not implement the range by repeatedly calling `Get`, or
by wrapping every visible node before slicing.

Preserve inclusive cursor-range semantics while defining slice bounds at one
place. Keep the hidden root excluded, collapsed descendants excluded, and
visible directory ordering unchanged. Preserve the empty-selection result
and the existing selection-clamping contract.

With a warm cache, process only the selected rows. A cold cache still needs
to enumerate visible rows once. Avoid a second cache or path index for this
change.

### Verify behavior and record the result

Extend the existing tree visibility and selection tests through the typed
tree interfaces. Cover both panels, empty trees, a shrinking tree, collapsed
directories, compression changes, and root-display modes. Compare selected
node identities and returned range bounds with the existing behavior.

Run the expanded benchmarks against identical fixtures before and after the
change. Include the cost of rebuilding the cache in the cold cases. Record
retained cache memory separately from allocations per selection.

Inspect the commits that introduced the shared selection helper and the
visible-row cache before choosing fixup targets. This proposal joins changes
introduced by several commits. Discuss the history placement with the user
if it does not belong to one existing logical unit. Do not rewrite or squash
the branch to make it fit.

## Prepare a query once while its content changes

### Establish the baseline

Inspect these modules and their callers:

- [`View.updateSearchPositions`](../../pkg/gocui/view.go) creates a
  `newLineSearcher` each time stale search positions are refreshed.
- [`newLineSearcher`](../../pkg/gocui/search.go) normalizes the query and
  allocates its graphemes, prefix table, and match-column scratch storage.
- Incremental search preserves unchanged result rows, but the prepared query
  currently lasts for only one refresh.
- [`BenchmarkIncrementalSearchedDiff`](../../pkg/gocui/search_stream_test.go)
  refreshes a fixed query after each batch. Its current fixture uses a short
  query and batches of 128 lines.

Extend the streaming benchmark with small and large batches, short and long
queries, matching and missing queries, and wrapped and unwrapped views.
Keep the total content fixed when comparing batch sizes. Separate initial
query preparation, subsequent refreshes, and query changes where practical.
Retain a changing-query case so reuse cannot conceal a regression when users
type a new query.

### Give prepared query state the query's lifetime

1. Keep the prepared `lineSearcher` with the owning `searcher`, rather than
   rebuilding it on every content refresh.
2. Prepare it when a query becomes active, or lazily at its first refresh.
   Choose one policy after measuring initial-search and refresh costs.
3. Replace it when the query changes. Text writes, wrapping changes, and
   buffer replacement may invalidate positions without invalidating the
   prepared query.
4. Release its query-derived slices when search is cleared. Verify the
   existing query-change path does not retain obsolete query storage.

Keep this state under the same synchronization and ownership as the view's
current search state. The `starts` slice is mutable scratch storage. Do not
share it between views or concurrent scans.

Keep the KMP match length and display-column position local to each scanned
line. Reuse query preparation across lines and refreshes, without treating
separate lines as a continuous match. Preserve overlapping matches, Unicode
graphemes, display widths, and smart-case behavior.

Keep model-provided search results on their existing result-generation
path. Query reuse does not authorize changes to model-result semantics or
incremental invalidation.

Keep upstream's draw walker local to a draw. Its cursor depends on the
current result slice and viewport, so query-lifetime storage is not a reason
to retain the draw walker across frames.

### Verify behavior and record the result

Extend the existing streaming differential test, which compares each content
update with a fresh full search. Cover query changes between batches,
clearing and restarting search, case changes, overlapping prefixes, Unicode,
rewrapping, overwritten rows, replacement buffers, and model-based results.
Compare both match positions and current-match selection where applicable.

Check retained memory after clearing a long query. Run the relevant race
tests if ownership or synchronization changes. Keep a performance result
only if repeated-refresh measurements support it without a material
regression in initial searches or changing-query workloads.

Inspect both the query-preparation and incremental-search commits before
choosing fixup targets. Discuss history placement if the change spans their
logical units. Leave any resulting fixups visible for user review.

## Finish each implementation independently

Run `just format`, `just build`, `just unit-test`, and `just lint` before each
implementation commit. Run relevant integration tests for the changed
behavior. Use `just generate` if integration tests are added or renamed.
Record benchmark commands, fixture definitions, pinned revisions, raw
samples, and measurement limits with each result.

Preserve upstream's selected-overlap highlighting behavior when integrating
#6201. Keep that integration separate from these follow-ups so its correctness
and performance can be checked on their own.
