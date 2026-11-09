# Implement the algorithm-context follow-ups

These plans expand the findings in the
[algorithm-context audit](../Algorithm-context-audit.md). Wait for upstream
PR #6201 to merge and for its merged code to be integrated into the fork
before implementing them. Recheck the source after integration.

Each plan describes a separate investigation and possible implementation:

| Plan | Current evidence | Next step |
| --- | --- | --- |
| [Graph occupancy](Graph-occupancy.md) | Repeated column checks counted in the real algorithm | Compare independent forward-only allocation cursors |
| [Editor navigation](Editor-navigation.md) | Navigation cost grows with the preceding text | Compare coordinate-based lookup with local cursor traversal |
| [Graph render scheduling](Graph-render-scheduling.md) | Serial and parallel paths win different viewport cases | Measure cached viewport work before choosing a policy |
| [Flat file lists](Flat-file-lists.md) | Intermediate directory structures are discarded | Establish ordering parity and measure flat-mode construction |

The earlier [post-merge plan](../Post-merge-performance-plan.md) covers
cached tree-range selection and prepared-query reuse. Keep those changes
separate from these four follow-ups.

## Start with the recorded evidence

Use the [audit evidence](../performance/algorithm-context-audit.json) to
recover the pinned revision, environment, raw samples, exact probe sources,
and commands. Its temporary alternatives are measurement tools. Re-evaluate
their implementation choices against the integrated code.

Agree on unresolved design choices with the user before production changes.
Treat operation counts and isolated benchmark results as evidence for where
to investigate. Measure complete caller operations before claiming an
interactive benefit.

## Keep each change independently reviewable

Find the commits that introduced the code before choosing a commit layout.
Use the repository's fixup rules for refinements to existing branch work.
Discuss placement when a change spans several logical units. Preserve
separate behavior-preserving preparation where the history rules require it.
Leave fixups visible for review.

Run `just format`, `just build`, `just unit-test`, and `just lint` before each
implementation commit. Run relevant integration tests for changed behavior.
Use `just generate` when adding or renaming integration tests.

Record matched fixtures, pinned revisions, raw benchmark samples, allocated
bytes, allocation counts, and retained-memory changes. State which setup,
rendering, and callback costs each measurement includes.
