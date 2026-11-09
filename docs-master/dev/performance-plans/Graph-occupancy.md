# Avoid repeated graph-column prefix scans

Implement this follow-up after the pending upstream integration, following
the [shared implementation guidance](README.md). This plan preserves graph
layout and changes how available columns are found within one row.

## Establish the current behavior

Inspect [`getNextPipes`](../../../pkg/gui/presentation/graph/graph.go), its
callers, and the existing layout tests. The fork's `f2ab2c610` replaced hashed
occupancy with a dense `columnState` array. Two closures still begin at zero
for every allocation:

- `getNextAvailablePosForContinuingPipe` rejects traversed columns.
- `getNextAvailablePosForNewPipe` rejects taken columns and columns occupied
  by continuing pipes.

The occupied sets only grow or remain fixed during one `getNextPipes` call.
A column rejected by either predicate cannot later become available to that
same predicate within the row.

The [audit evidence](../performance/algorithm-context-audit.json) records
2,016 checks when a 64-column fixture assigns its 63 continuing pipes. An
octopus merge with 64 parents makes the same number of new-parent checks.
These counts demonstrate repeated work, not its share of real application
CPU.

## Measure graph width independently

Add fixtures that vary active columns and parent count independently of
commit count. Include ordinary narrow graphs, wide unrelated histories,
sparse incoming columns, merges, octopus merges, and root commits.

Measure cold `GetPipeSets` and full graph construction separately from
cached viewport rendering. Layout improvements affect cache misses; the UI
does not rebuild layout on every selection change.

Count occupancy probes separately from timing. Keep instrumentation out of
timed runs. Record maximum active width for any representative repository
profile so the synthetic wide cases have context.

## Compare forward-only allocation cursors

Give each allocation closure its own next-candidate position, local to one
`getNextPipes` call. Advance it over rejected columns. After the caller marks
the returned position occupied, the next request may continue there or
advance past it, according to that predicate's consumption rules.

Keep the cursors separate because the predicates reject different sets.
Discard both cursors after the row. Preserve the order of incoming pipes,
parent allocation, traversal marking, and final pipe ordering.

Verify the monotonicity assumption against every write to `taken`,
`traversed`, and `continuing` in the integrated code. If a state transition
can free a rejected column, discuss a compatible algorithm before changing
allocation order.

Keep the occupancy array and existing graph representation for this
experiment. A bitset, interval structure, or broader graph rewrite needs
separate evidence. Avoid introducing a general allocator interface for
these two private closures.

## Prove layout parity and measure complete work

Compare complete `Pipe` results with the current algorithm across generated
histories and the targeted fixtures. Preserve pooled hash identities in
fixtures because graph equality relies on the existing hash representation.

Run existing layout and rendering tests for both symbol sets. Cover
selected and unselected commits, sparse columns, unrelated histories,
multiple terminating pipes, and multiple parents. Check empty inputs and
root behavior through their existing public callers.

Confirm that the occupancy searches no longer repeatedly visit a rejected
prefix. A linear bound for those searches does not imply that all graph
layout work is linear; traversal marking and sorting remain separate costs.

Compare narrow and wide construction times, allocation counts, and cached
redraw behavior. Keep the change only if complete workloads support it
without a material narrow-graph regression.

Before committing, inspect `f2ab2c610` and the earlier allocation-loop history
to decide where this refinement belongs. Discuss placement if the change
does not fit one existing logical unit.
