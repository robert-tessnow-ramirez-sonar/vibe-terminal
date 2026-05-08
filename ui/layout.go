package ui

import "github.com/rivo/tview"

// Notice we now return TWO *tview.Flex objects and accept an initial width
func SetupLayout(dropdown *tview.DropDown, list *tview.List, terminalOut *tview.TextView, initialListWidth int) (*tview.Flex, *tview.Flex) {

	// Create the side-by-side layout
	bodyFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(list, initialListWidth, 0, false). // FIXED width
		AddItem(terminalOut, 0, 1, false)          // Dynamic width (fills the rest)

	// Main vertical layout
	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText(" ⚡ VIBE CODE TERMINAL ⚡ ").SetTextAlign(tview.AlignCenter), 1, 1, false).
		AddItem(dropdown, 3, 1, true).
		AddItem(bodyFlex, 0, 1, false).
		AddItem(tview.NewTextView().SetText(" [Shift+Left/Right] Resize | [Tab/Arrows] Focus | [Enter] Run | [Ctrl+C] Quit"), 1, 1, false)

	return mainFlex, bodyFlex
}
