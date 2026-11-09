package filetree

import (
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/common"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/samber/lo"
)

type ICommitFileTree interface {
	ITree[models.CommitFile]

	Get(index int) *CommitFileNode
	GetRange(start, end int) []*CommitFileNode
	GetFile(path string) *models.CommitFile
	GetAllItems() []*CommitFileNode
	GetAllFiles() []*models.CommitFile
	GetRoot() *CommitFileNode
	SetTextFilter(filter string, useFuzzySearch bool)
	GetTextFilter() string
}

type CommitFileTree struct {
	treeView[models.CommitFile]
	getFiles       func() []*models.CommitFile
	showTree       bool
	common         *common.Common
	textFilter     string
	useFuzzySearch bool
}

func (self *CommitFileTree) CollapseAll() {
	dirPaths := lo.FilterMap(self.GetAllItems(), func(file *CommitFileNode, index int) (string, bool) {
		return file.path, !file.IsFile()
	})

	for _, path := range dirPaths {
		self.collapsedPaths.Collapse(path)
	}
}

var _ ICommitFileTree = &CommitFileTree{}

func NewCommitFileTree(getFiles func() []*models.CommitFile, common *common.Common, showTree bool) *CommitFileTree {
	return &CommitFileTree{
		getFiles: getFiles,
		common:   common,
		showTree: showTree,
		treeView: newTreeView[models.CommitFile](),
	}
}

func (self *CommitFileTree) ToggleShowTree() {
	self.showTree = !self.showTree
	self.SetTree()
}

func (self *CommitFileTree) Get(index int) *CommitFileNode {
	// need to traverse the three depth first until we get to the index.
	return NewCommitFileNode(self.getNode(index)) // ignoring root
}

func (self *CommitFileTree) GetRange(start, end int) []*CommitFileNode {
	return lo.Map(self.getNodes()[start:end], func(node *Node[models.CommitFile], _ int) *CommitFileNode {
		return NewCommitFileNode(node)
	})
}

func (self *CommitFileTree) GetAllItems() []*CommitFileNode {
	if self.tree == nil {
		return nil
	}

	// ignoring root
	return lo.Map(self.getNodes(), func(node *Node[models.CommitFile], _ int) *CommitFileNode {
		return NewCommitFileNode(node)
	})
}

func (self *CommitFileTree) Len() int {
	return self.length()
}

func (self *CommitFileTree) GetItem(index int) types.HasUrn {
	// Unimplemented because we don't yet need to show inlines statuses in commit file views
	return nil
}

func (self *CommitFileTree) GetAllFiles() []*models.CommitFile {
	return self.getFiles()
}

func (self *CommitFileTree) getFilesForDisplay() []*models.CommitFile {
	files := self.getFiles()
	if self.textFilter != "" {
		files = filterCommitFilesByText(files, self.textFilter, self.useFuzzySearch)
	}
	return files
}

func (self *CommitFileTree) SetTree() {
	filesForDisplay := self.getFilesForDisplay()
	guiConfig := self.common.UserConfig().Gui
	showRootItem := guiConfig.ShowRootItemInFileTree
	cmp := NodeSortComparator[models.CommitFile](guiConfig.FileTreeSortOrder, guiConfig.FileTreeSortCaseSensitive)
	if self.showTree {
		self.tree = BuildTreeFromCommitFiles(filesForDisplay, showRootItem, cmp)
	} else {
		self.tree = BuildFlatTreeFromCommitFiles(filesForDisplay, showRootItem, cmp)
	}
}

func (self *CommitFileTree) SetTextFilter(filter string, useFuzzySearch bool) {
	self.textFilter = filter
	self.useFuzzySearch = useFuzzySearch
	self.SetTree()
}

func (self *CommitFileTree) GetTextFilter() string {
	return self.textFilter
}

func (self *CommitFileTree) GetRoot() *CommitFileNode {
	return NewCommitFileNode(self.tree)
}

func (self *CommitFileTree) GetFile(path string) *models.CommitFile {
	for _, file := range self.getFiles() {
		if file.Path == path {
			return file
		}
	}

	return nil
}

func (self *CommitFileTree) InTreeMode() bool {
	return self.showTree
}
