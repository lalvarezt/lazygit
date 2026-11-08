package main_view

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var CopyCombinedDiffWithoutLineMetadata = NewIntegrationTest(NewIntegrationTestArgs{
	Description: "Combined merge diff rows without line metadata cannot be copied as file lines",
	SetupConfig: func(cfg *config.AppConfig) {
		cfg.GetUserConfig().Gui.UseHunkModeInDiffView = false
		cfg.GetUserConfig().OS.CopyToClipboardCmd = "printf '%s' {{text}} > clipboard"
		cfg.GetUserConfig().Git.DiffRenderers = []config.DiffRendererConfig{
			{
				// Delta describes combined diff file headers but not their content rows.
				Command: `printf '\033]1717;1\007'; ` +
					`awk '/^diff --cc / { printf "\033]1717;1;f;;;file1\007" } { print }'`,
				ColorArg: "never",
			},
		}
	},
	SetupRepo: func(shell *Shell) {
		shell.CreateFileAndAdd("file1", "original\n")
		shell.Commit("base")
		shell.NewBranch("left")
		shell.UpdateFileAndAdd("file1", "left\n")
		shell.Commit("left")
		shell.NewBranchFrom("right", "HEAD~1")
		shell.UpdateFileAndAdd("file1", "right\n")
		shell.Commit("right")
		shell.Checkout("left")
		shell.RunCommandExpectError([]string{"git", "merge", "right"})
		shell.UpdateFileAndAdd("file1", "resolved\n")
		shell.Commit("merge")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Commits().
			Focus().
			Press(keys.Universal.FocusMainView)

		t.Views().Main().
			IsFocused().
			Content(Contains("diff --cc file1")).
			Content(Contains("++resolved")).
			Press(keys.Universal.CopyToClipboard)

		t.ExpectToast(Equals("Disabled: There is nothing to select here"))
		t.FileSystem().PathNotPresent("clipboard")
	},
})
