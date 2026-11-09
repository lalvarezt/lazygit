# Avoid unused directory structure in flat file lists

Investigate this follow-up after the pending upstream integration, following
the [shared implementation guidance](README.md). This is a source-backed
candidate with no measured alternative yet.

## Establish the flat-mode contract

Inspect [`build_tree.go`](../../../pkg/gui/filetree/build_tree.go), node
sorting, both file-tree modules, and their selection-recovery callers.

`BuildFlatTreeFromFiles` and `BuildFlatTreeFromCommitFiles` currently build,
sort, and compress a directory tree, collect its leaves, and retain only the
file nodes. The indexed builder improves child lookup during construction,
but flat-mode output does not display the intermediate directories.

Record the current output for every supported sort order and case setting.
Preserve hierarchical leaf ordering, root-prefix rules, and file-node
identities as observed by callers. Working-tree flat mode also applies a
stable conflict and tracked-status ordering.

Do not assume that sorting full path strings reproduces the current leaf
order. Do not assume input arrives sorted. Files-first and folders-first
hierarchical ordering can affect where leaves appear even when the final
view displays no directories.

## Measure before choosing a replacement

Add flat-mode construction cases for working-tree and commit files. Vary
file count, path depth, common prefixes, sibling counts, input order, and
status groups. Include small ordinary lists as well as large lists.

Compare cold construction and complete refresh, including visibility-index
construction and selection recovery where the caller performs them. Keep
fixture setup outside construction timing.

Measure allocated bytes, allocations, retained memory, and elapsed time.
The existing tree-construction benchmark is not sufficient evidence for
flat mode. Treat a direct builder as an experiment until those measurements
show that intermediate directory construction materially matters.

## Compare a direct leaf builder with the current tree path

Prototype constructing only the file nodes required by flat mode. Derive
ordering from the existing contract rather than duplicating an approximate
comparator.

Compare ways to share ordering semantics with tree mode. A path-component
comparison or prepared ordering keys may be appropriate, but preprocessing
and temporary storage must be included in the cost. Choose the simplest
compatible approach after measurement.

Keep typed wrappers and the visible-row index compatible with both panels.
Preserve internal paths, display paths, underlying file pointers, rename
information, and any root-prefix convention used by selection recovery.
Handle empty input without constructing unnecessary auxiliary state.

If exact ordering requires a substantial redesign shared with tree mode,
present that choice to the user before implementation. Avoid maintaining
two subtly different definitions of file order.

## Prove parity at construction and caller interfaces

Compare ordered leaf paths and underlying files with the current builder
across generated inputs. Include:

- Mixed, files-first, and folders-first sorting with both case settings.
- Top-level files, deep paths, shared prefixes, and compressed directories.
- Root-display settings, the single-file special case, and empty input.
- Working-tree conflict and tracked-status groups with stable ordering.
- Rename cases, text and status filters, and tree/flat mode switches.
- Selection preservation after refresh, file removal, and directory changes.

Run existing tree and selection tests for both panels. Add relevant
integration coverage if construction-level parity does not exercise a
changed caller contract.

Keep the replacement only if full flat-mode workloads support it without
changing ordering or imposing a material small-list regression. Record
sorting and preprocessing costs, not just the final node-allocation loop.

Inspect the shared builder, indexed-construction, and selection-related
commits before choosing history placement. Discuss it if the work cuts
across several existing logical units.
