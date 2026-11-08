package filter_and_search

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var SearchWhileADiffLoads = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Typing and cancelling search work while a long diff is still loading",
	Width:        120,
	Height:       40,
	ExtraEnvVars: map[string]string{"LAZYGIT_SLOW_RENDER": "1"},
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.CreateFileAndAdd("file1", "")
		shell.Commit("one")
		shell.UpdateFile("file1", longFileWithThreeMatches())
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().IsFocused().Press(keys.Universal.FocusMainView)
		t.Views().Main().IsFocused().PressRapidly(
			keys.Universal.StartSearch,
			config.Keybinding{"N"}, config.Keybinding{"E"}, config.Keybinding{"E"},
			config.Keybinding{"D"}, config.Keybinding{"L"}, config.Keybinding{"E"},
		)
		t.Views().Search().IsFocused().Content(Equals("NEEDLE"))
		t.ExpectSearch().Confirm()
		t.Views().Search().Content(Contains("matches for 'NEEDLE' (1 of 3)"))

		t.Views().Main().Press(keys.Universal.Return).
			Press(keys.Universal.IncreaseContextInDiffView).
			Tap(func() { t.ExpectToast(Equals("Changed diff context size to 4")) }).
			PressRapidly(
				keys.Universal.StartSearch,
				config.Keybinding{"N"}, config.Keybinding{"E"}, config.Keybinding{"E"},
				config.Keybinding{"D"}, config.Keybinding{"L"}, config.Keybinding{"E"},
				config.Keybinding{"<enter>"},
			).
			SelectedLines(Contains("+NEEDLE first")).
			Press(keys.Universal.Return).
			Press(keys.Universal.IncreaseContextInDiffView).
			Tap(func() { t.ExpectToast(Equals("Changed diff context size to 5")) }).
			PressRapidly(keys.Universal.StartSearch, keys.Universal.Return).
			IsFocused()
	},
})
