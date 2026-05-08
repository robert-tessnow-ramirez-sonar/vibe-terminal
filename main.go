package main

import (
	"vibe-terminal/ui"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	commandList := tview.NewList()

	terminalOut := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWrap(true)

	categoryDropdown := ui.CreateDropdown(commandList, app, terminalOut)

	commandList.SetBorder(true).SetTitle(" [ Commands ] ")
	categoryDropdown.SetBorder(true).SetTitle(" [ Categories ] ")
	terminalOut.SetBorder(true).SetTitle(" [ Terminal Output ] ")

	// --- RESIZE STATE ---
	listWidth := 30 // Starting width of the command list

	// Setup layout now catches both the main layout and the body container
	layout, bodyFlex := ui.SetupLayout(categoryDropdown, commandList, terminalOut, listWidth)

	// --- FOCUS LOGIC ---
	focusElements := []tview.Primitive{categoryDropdown, commandList, terminalOut}
	currentFocus := 0

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// 1. Check for RESIZE commands (Shift + Arrows)
		if event.Modifiers() == tcell.ModShift {
			switch event.Key() {
			case tcell.KeyLeft:
				if listWidth > 15 { // Prevent crushing the list too small
					listWidth -= 3 // Jump 3 columns for faster resizing
					bodyFlex.ResizeItem(commandList, listWidth, 0)
				}
				return nil // We handled the key, stop bubbling

			case tcell.KeyRight:
				if listWidth < 80 { // Prevent crushing the terminal output
					listWidth += 3
					bodyFlex.ResizeItem(commandList, listWidth, 0)
				}
				return nil
			}
		}

		// 2. Check for FOCUS commands (Standard Arrows / Tab)
		// We only hit this if Shift was NOT held down
		switch event.Key() {
		case tcell.KeyTab, tcell.KeyRight:
			currentFocus = (currentFocus + 1) % len(focusElements)
			app.SetFocus(focusElements[currentFocus])
			return nil

		case tcell.KeyLeft:
			currentFocus = (currentFocus - 1 + len(focusElements)) % len(focusElements)
			app.SetFocus(focusElements[currentFocus])
			return nil
		}

		return event
	})

	app.EnableMouse(true)

	if err := app.SetRoot(layout, true).SetFocus(categoryDropdown).Run(); err != nil {
		panic(err)
	}
}
