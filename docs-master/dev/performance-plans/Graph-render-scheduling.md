# Schedule graph rendering for the visible workload

Implemented with shared output storage and a pipe-count worker budget. See
the [implementation record](../Independent-performance-follow-ups.md) for
commits and validation. The plan below selects scheduling from caller-shaped
measurements while preserving rendered lines and order.

## Separate cold layout from viewport redraws

Inspect [`RenderAux`](../../../pkg/gui/presentation/graph/graph.go) and its
callers in [`presentation/commits.go`](../../../pkg/gui/presentation/commits.go).
The UI obtains cached pipe sets, then passes a visible slice to the renderer.
The principal graph benchmark calls `RenderCommitGraph`, which also builds
layout and renders the whole fixture.

`RenderAux` currently creates `GOMAXPROCS` workers regardless of row count.
Each worker builds an output slice; the caller then flattens the slices.
Some workers receive empty ranges when rows are fewer than workers.

The [audit evidence](../performance/algorithm-context-audit.json) includes a
sequential probe with checked output parity. Serial rendering wins some
small cases; parallel rendering wins some larger cases. Those short samples
do not define a production cutoff.

## Establish representative rendering measurements

Create separate cases for cold layout, cached viewport rendering, and
redrawing after selected-commit changes. Vary:

- Visible row count, including empty and very short slices.
- Active graph width and merge density, independently of history length.
- Box and branch symbols, with and without a selected commit.
- Available CPUs and the number of workers actually used.
- Normal and divergence views, including a viewport spanning both sections.

Keep fixture creation and layout outside the cached-render timing. Control
color settings consistently. Measure allocations, bytes, and elapsed time.
Use representative terminal heights and graph shapes before tuning policy.

## Compare scheduling and output allocation independently

Compare a sequential small-work path with capped worker counts. Consider row
width or another cheap work estimate only if row count alone proves
insufficient. Avoid making the work estimator as costly as the rendering
it schedules.

Evaluate one output slice whose disjoint ranges are filled by workers.
Measure that allocation change separately from policy changes, so its
benefit and concurrency behavior remain attributable.

Preserve useful parallelism for larger workloads. Do not replace the
renderer with an always-sequential implementation based on tiny viewports.
Do not start workers with empty ranges merely to match CPU count.

Present the measured policy options and their tradeoffs to the user before
selecting cutoffs or adding a work estimator. Keep the chosen policy private
to graph rendering unless another real caller needs to control it.

## Verify output and concurrency

Compare complete output slices with the existing renderer for all workload
variants. Preserve ordering, styles, selected-commit treatment, and previous
commit handling at viewport boundaries. Preserve empty-input semantics.

Run graph tests and relevant commit-view integration tests. Check the race
detector when changing shared output writes or scheduling. Ensure every
worker's output range is disjoint and complete, including uneven divisions.

Profile a cached viewport redraw in the application before claiming an
interactive improvement. Include CPU and allocation costs for creating,
joining, and combining worker results in the measured operation.

This scheduling policy predates the fork's graph changes. Inspect its
history and the fork's cell-allocation commits before splitting preparation,
allocation changes, and behavior-preserving scheduling changes into commits.
