package filetree

import (
	"slices"
	"sort"
	"strings"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
)

func BuildTreeFromFiles(
	files []*models.File,
	showRootItem bool,
	order NodeSortOrder,
) *Node[models.File] {
	return buildTree(files, func(file *models.File) string { return file.Path }, showRootItem, order)
}

func buildTree[T any](
	files []*T,
	getPath func(*T) string,
	showRootItem bool,
	order NodeSortOrder,
) *Node[T] {
	root := &Node[T]{}

	childrenMapsByNode := make(map[*Node[T]]map[string]*Node[T])

	var curr *Node[T]
	for _, file := range files {
		splitPath := SplitFileTreePath(getPath(file), showRootItem)
		curr = root
	outer:
		for i := range splitPath {
			var setFile *T
			isFile := i == len(splitPath)-1
			if isFile {
				setFile = file
			}

			path := join(splitPath[:i+1])

			var currNodeChildrenMap map[string]*Node[T]
			var isCurrNodeMapped bool

			if currNodeChildrenMap, isCurrNodeMapped = childrenMapsByNode[curr]; !isCurrNodeMapped {
				currNodeChildrenMap = make(map[string]*Node[T])
				childrenMapsByNode[curr] = currNodeChildrenMap
			}

			child, doesCurrNodeHaveChildAlready := currNodeChildrenMap[path]
			if doesCurrNodeHaveChildAlready {
				curr = child
				continue outer
			}

			if i == 0 && len(files) == 1 && len(splitPath) == 2 {
				// skip the root item when there's only one file at top level; we don't need it in that case
				continue outer
			}

			newChild := &Node[T]{
				path: path,
				File: setFile,
			}
			curr.Children = append(curr.Children, newChild)

			currNodeChildrenMap[path] = newChild

			curr = newChild
		}
	}

	root.Sort(order)
	root.Compress()

	return root
}

func BuildFlatTreeFromCommitFiles(
	files []*models.CommitFile,
	showRootItem bool,
	order NodeSortOrder,
) *Node[models.CommitFile] {
	sortedFiles := flatTreeLeaves(files, func(file *models.CommitFile) string { return file.Path }, showRootItem, order)
	return &Node[models.CommitFile]{Children: sortedFiles}
}

func flatTreeLeaves[T any](files []*T, getPath func(*T) string, showRootItem bool, order NodeSortOrder) []*Node[T] {
	type flatFile struct {
		file     *T
		path     string
		sortPath string
	}
	type pathKind struct {
		original string
		dir      bool
	}
	paths := make(map[string]pathKind, len(files))
	entries := make([]flatFile, 0, len(files))
	for _, file := range files {
		path := InternalTreePathForFilePath(getPath(file), showRootItem)
		sortPath := order.sortPath(path)
		rawRest, sortRest := path, sortPath
		rawEnd, sortEnd := 0, 0
		duplicate := false
		for {
			rawName, rawTail, dir := strings.Cut(rawRest, "/")
			sortName, sortTail, _ := strings.Cut(sortRest, "/")
			rawEnd += len(rawName)
			sortEnd += len(sortName)
			original := path[:rawEnd]
			key := sortPath[:sortEnd]
			if previous, found := paths[key]; found {
				if previous.original != original || previous.dir != dir {
					// Case aliases and file/directory collisions depend on the
					// hierarchy's insertion and sibling-sort order.
					return buildTree(files, getPath, showRootItem, order).GetLeaves()
				}
				duplicate = !dir
			} else {
				paths[key] = pathKind{original: original, dir: dir}
			}
			if !dir {
				break
			}
			rawEnd++
			sortEnd++
			rawRest, sortRest = rawTail, sortTail
		}
		if !duplicate {
			entries = append(entries, flatFile{file: file, path: path, sortPath: sortPath})
		}
	}
	// Each sort key has already been normalized for the requested case mode.
	order.caseSensitive = true
	slices.SortFunc(entries, func(a, b flatFile) int {
		left, right := a.sortPath, b.sortPath
		for {
			aName, aTail, aDir := strings.Cut(left, "/")
			bName, bTail, bDir := strings.Cut(right, "/")
			if aName != bName || aDir != bDir {
				return order.compare(aName, bName, aDir, bDir)
			}
			if !aDir {
				return 0
			}
			left, right = aTail, bTail
		}
	})
	leaves := make([]*Node[T], len(entries))
	for i, entry := range entries {
		leaves[i] = &Node[T]{File: entry.file, path: entry.path}
	}
	return leaves
}

func BuildTreeFromCommitFiles(
	files []*models.CommitFile,
	showRootItem bool,
	order NodeSortOrder,
) *Node[models.CommitFile] {
	return buildTree(files, func(file *models.CommitFile) string { return file.Path }, showRootItem, order)
}

func BuildFlatTreeFromFiles(
	files []*models.File,
	showRootItem bool,
	order NodeSortOrder,
) *Node[models.File] {
	sortedFiles := flatTreeLeaves(files, func(file *models.File) string { return file.Path }, showRootItem, order)

	// from top down we have merge conflict files, then tracked file, then untracked
	// files. This is the one way in which sorting differs between flat mode and
	// tree mode
	sort.SliceStable(sortedFiles, func(i, j int) bool {
		iFile := sortedFiles[i].File
		jFile := sortedFiles[j].File

		// never going to happen but just to be safe
		if iFile == nil || jFile == nil {
			return false
		}

		if iFile.HasMergeConflicts && !jFile.HasMergeConflicts {
			return true
		}

		if jFile.HasMergeConflicts && !iFile.HasMergeConflicts {
			return false
		}

		if iFile.Tracked && !jFile.Tracked {
			return true
		}

		if jFile.Tracked && !iFile.Tracked {
			return false
		}

		return false
	})

	return &Node[models.File]{Children: sortedFiles}
}

func split(str string) []string {
	return strings.Split(str, "/")
}

func join(strs []string) string {
	return strings.Join(strs, "/")
}

func SplitFileTreePath(path string, showRootItem bool) []string {
	return split(InternalTreePathForFilePath(path, showRootItem))
}

func InternalTreePathForFilePath(path string, showRootItem bool) string {
	if showRootItem {
		return "./" + path
	}

	return path
}
