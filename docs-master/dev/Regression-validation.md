# Validate lazygit changes

For local edits and upstream integration, use the shared protocol at
`~/.agents/skills/merge-upstream/references/regression-validation.md`.

## Caller coverage

- Commits and graphs: fixed viewports over growing histories; branch markers,
  main-tip states, graph widths, symbols, selection, and CPU counts.
- Editors: full key handling near buffer start and end, wrapping, Unicode,
  boundaries, edits, and resize.
- File trees: both panels and modes, small and large lists, ordering edge cases,
  construction, refresh, and cached and cold selection.
- Search: changing and repeated queries, misses, streamed append and
  replacement, wrapping, navigation, loading, cancellation, and current-match
  preservation.
- Derived state: commit, subcommit, branch, configuration, dimension, and
  repository changes; verify the state published on the UI thread.
- Diff integrations: configured producer-to-view behavior, first output,
  completion, cancellation, row metadata, hyperlinks, and consuming actions.

See the [snapshot audit](Snapshot-regression-audit.md) for the caller work
missed by graph-only benchmarks.

## Checks and measurements

Follow `AGENTS.md` and the `justfile`; include relevant integration and race
checks. For color-sensitive checks, remove `NO_COLOR` and set `FORCE_COLOR=1`
and `CLICOLOR_FORCE=1`.

Compare pinned local snapshots with the affected packages:

```sh
just perf-compare <baseline> <candidate> --prepare-only
just perf-compare <baseline> <candidate> --resume --package=./pkg/gui/presentation
```

Package-only runs save samples and metadata in `.git/perf-comparison/`.
Copy them before switching references: output filenames are reused. A full
run rewrites the README's historical fork-base report; publish that deliberately.

`BenchmarkCommitViewport` excludes ancestry preparation. Measure model updates
separately when that lifetime changes:

```sh
go test ./pkg/gui/types -run='^$' -bench='^BenchmarkModelBranchMarkerUpdate$' -benchmem -count=5 -benchtime=200ms -cpu=1
```

Match GC and memory controls with the runner. Keep durable evidence under
`docs-master/dev/` and exploratory output under `.git/`.
