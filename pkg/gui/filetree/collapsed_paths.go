package filetree

import "github.com/jesseduffield/generics/set"

type CollapsedPaths struct {
	collapsedPaths *set.Set[string]
	revision       uint64
}

func NewCollapsedPaths() *CollapsedPaths {
	return &CollapsedPaths{
		collapsedPaths: set.New[string](),
	}
}

func (self *CollapsedPaths) ExpandToPath(path string) {
	// need every directory along the way
	splitPath := split(path)
	for i := range splitPath {
		dir := join(splitPath[0 : i+1])
		if self.collapsedPaths.Includes(dir) {
			self.collapsedPaths.Remove(dir)
			self.revision++
		}
	}
}

func (self *CollapsedPaths) IsCollapsed(path string) bool {
	return self.collapsedPaths.Includes(path)
}

func (self *CollapsedPaths) Collapse(path string) {
	if !self.collapsedPaths.Includes(path) {
		self.collapsedPaths.Add(path)
		self.revision++
	}
}

func (self *CollapsedPaths) ToggleCollapsed(path string) {
	self.revision++
	if self.collapsedPaths.Includes(path) {
		self.collapsedPaths.Remove(path)
	} else {
		self.collapsedPaths.Add(path)
	}
}

func (self *CollapsedPaths) ExpandAll() {
	// Could be cleaner if Set had a Clear() method...
	paths := self.collapsedPaths.ToSlice()
	if len(paths) > 0 {
		self.collapsedPaths.RemoveSlice(paths)
		self.revision++
	}
}
