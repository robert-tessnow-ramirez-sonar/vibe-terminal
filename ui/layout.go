package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetupLayout(catBar *tview.Table, cmdList *tview.List, terminalOut *tview.TextView, initialWidth int) (*tview.Flex, *tview.Flex) {

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tview.NewTextView().SetText(" Categories: ").SetTextColor(tcell.ColorYellow), 13, 0, false).
		AddItem(catBar, 0, 1, true)

	bodyFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(cmdList, initialWidth, 0, false).
		AddItem(terminalOut, 0, 1, false)

	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText(" ⚡ VIBE CODE ⚡ ").SetTextAlign(tview.AlignCenter), 1, 1, false).
		AddItem(header, 1, 1, true).
		AddItem(bodyFlex, 0, 1, false).
		// NEW: Updated Footer Instructions
		AddItem(tview.NewTextView().
			SetText(" [Enter/Dbl-Click] Run | [PgUp/PgDn] Scroll Output | [Ctrl+Y] Copy | [Tab] Focus ").
			SetTextAlign(tview.AlignCenter), 1, 1, false)

	return mainLayout, bodyFlex
}
