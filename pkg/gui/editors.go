package gui

import (
	"github.com/jesseduffield/lazygit/pkg/gocui"
)

func (gui *Gui) handleEditorKeypress(v *gocui.View, key gocui.Key, allowMultiline bool) bool {
	if key.Equals(gocui.NewKeyName(gocui.KeyEnter)) && allowMultiline {
		v.TextArea.TypeCharacter("\n")
		v.RenderTextArea()
		return true
	}

	return gocui.DefaultEditor.Edit(v, key)
}

func (gui *Gui) commitMessageEditor(v *gocui.View, key gocui.Key) bool {
	matched := gui.handleEditorKeypress(v, key, false)
	gui.c.Contexts().CommitMessage.RenderSubtitle()
	return matched
}

func (gui *Gui) commitDescriptionEditor(v *gocui.View, key gocui.Key) bool {
	return gui.handleEditorKeypress(v, key, true)
}

func (gui *Gui) promptEditor(v *gocui.View, key gocui.Key) bool {
	matched := gui.handleEditorKeypress(v, key, false)

	suggestionsContext := gui.State.Contexts.Suggestions
	// Capture the suggestions function and the input here, on the UI thread; the
	// main thread rewrites State.FindSuggestions when it (re)creates a prompt
	// panel, so reading it from the worker below would race that write.
	if findSuggestions := suggestionsContext.State.FindSuggestions; findSuggestions != nil {
		input := v.TextArea.GetContent()
		suggestionsContext.State.AsyncHandler.Do(func() func() {
			suggestions := findSuggestions(input)
			return func() { suggestionsContext.SetSuggestions(suggestions) }
		})
	}

	return matched
}

func (gui *Gui) menuFilterEditor(v *gocui.View, key gocui.Key) bool {
	contentBefore := v.TextArea.GetContent()

	matched := gui.handleEditorKeypress(v, key, false)
	if !matched {
		// Give the global keybindings a chance at the key, e.g. so that ctrl-c
		// still quits while a menu is open.
		return false
	}

	content := v.TextArea.GetContent()
	if content == contentBefore {
		// The key just moved the cursor around within the filter; refiltering would
		// throw away the menu's selection for nothing.
		return true
	}

	menuContext := gui.State.Contexts.Menu
	menuContext.SetFilterStarted(true)
	gui.helpers.Search.ApplyFilter(menuContext, content)

	return true
}

func (gui *Gui) searchEditor(v *gocui.View, key gocui.Key) bool {
	matched := gui.handleEditorKeypress(v, key, false)

	searchString := v.TextArea.GetContent()

	gui.helpers.Search.OnPromptContentChanged(searchString)

	return matched
}
