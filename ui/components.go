package ui

import (
	"fmt"
	"io"
	"os/exec"
	"vibe-terminal/commands"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateCategoryBar(list *tview.List, app *tview.Application, terminalOut *tview.TextView) *tview.Table {
	catBar := tview.NewTable().
		SetBorders(false).
		SetSelectable(true, true)

	names := commands.GetCategoryNames()
	for col, name := range names {
		cell := tview.NewTableCell(" " + name + " ").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter)
		catBar.SetCell(0, col, cell)
	}

	catBar.SetSelectionChangedFunc(func(row, col int) {
		if col >= 0 && col < len(names) {
			categoryName := names[col]
			list.Clear()
			for _, cmd := range commands.Categories[categoryName] {
				list.AddItem(cmd.Name, cmd.Execute, 0, nil)
			}
		}
	})

	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if secondaryText != "" {
			terminalOut.Clear()
			ansiWriter := tview.ANSIWriter(terminalOut)
			fmt.Fprintf(ansiWriter, "Running: \x1b[32m%s\x1b[0m...\n\n", secondaryText)

			go func() {
				cmd := exec.Command("bash", "-c", secondaryText)
				stdout, _ := cmd.StdoutPipe()
				stderr, _ := cmd.StderrPipe()
				cmd.Start()
				go io.Copy(ansiWriter, stdout)
				go io.Copy(ansiWriter, stderr)
				cmd.Wait()
			}()

			// NEW: Explicitly force the app to keep your cursor on the list!
			app.SetFocus(list)
		}
	})

	return catBar
}
