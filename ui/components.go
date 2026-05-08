package ui

import (
	"fmt"
	"os/exec"
	"vibe-terminal/commands"

	"github.com/rivo/tview"
)

func CreateDropdown(list *tview.List, app *tview.Application, terminalOut *tview.TextView) *tview.DropDown {
	dropdown := tview.NewDropDown().SetLabel("Select Category (Enter): ")

	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if secondaryText != "" {
			terminalOut.Clear()
			terminalOut.SetText("Running: " + secondaryText + "...\n")

			// Run the command in a goroutine so it doesn't freeze the UI
			go func() {
				cmd := exec.Command("bash", "-c", secondaryText)
				out, err := cmd.CombinedOutput() // Captures both stdout and stderr

				// Safely update the UI from the background thread
				app.QueueUpdateDraw(func() {
					if err != nil {
						terminalOut.SetText(fmt.Sprintf("Error: %v\n\n%s", err, string(out)))
					} else {
						terminalOut.SetText(string(out))
					}
				})
			}()
		}
	})

	onSelect := func(text string, index int) {
		list.Clear()
		for _, cmd := range commands.Categories[text] {
			list.AddItem(cmd.Name, cmd.Execute, 0, nil)
		}
	}

	categoryNames := commands.GetCategoryNames()
	dropdown.SetOptions(categoryNames, onSelect)

	for i, name := range categoryNames {
		if name == "Networking" {
			dropdown.SetCurrentOption(i)
			onSelect(name, i)
			break
		}
	}

	return dropdown
}
