package main

import (
	"time"
	"vibe-terminal/ui"

	"github.com/atotto/clipboard"
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

	// NEW: Auto-scroll to the bottom as commands output text
	terminalOut.SetChangedFunc(func() {
		terminalOut.ScrollToEnd()
		app.Draw()
	})

	catBar := ui.CreateCategoryBar(commandList, app, terminalOut)

	commandList.SetBorder(true).SetTitle(" [ Commands ] ")
	terminalOut.SetBorder(true).SetTitle(" [ Terminal Output ] ")

	listWidth := 30
	layout, bodyFlex := ui.SetupLayout(catBar, commandList, terminalOut, listWidth)

	focusElements := []tview.Primitive{catBar, commandList, terminalOut}
	currentFocus := 0

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// 1. COPY LOGIC (Ctrl+Y)
		if event.Key() == tcell.KeyCtrlY {
			clipboard.WriteAll(terminalOut.GetText(false))
			terminalOut.SetTitle(" [ COPIED TO CLIPBOARD! ] ")
			go func() {
				time.Sleep(2 * time.Second)
				app.QueueUpdateDraw(func() { terminalOut.SetTitle(" [ Terminal Output ] ") })
			}()
			return nil
		}

		// 2. GLOBAL SCROLLING (Scroll terminal without switching context!)
		if event.Key() == tcell.KeyPgUp {
			row, col := terminalOut.GetScrollOffset()
			terminalOut.ScrollTo(row-3, col)
			return nil
		}
		if event.Key() == tcell.KeyPgDn {
			row, col := terminalOut.GetScrollOffset()
			terminalOut.ScrollTo(row+3, col)
			return nil
		}

		// 3. Category Bar Navigation
		if app.GetFocus() == catBar {
			if event.Key() == tcell.KeyLeft || event.Key() == tcell.KeyRight {
				return event
			}
		}

		// 4. Tab Navigation
		if event.Key() == tcell.KeyTab {
			currentFocus = (currentFocus + 1) % len(focusElements)
			app.SetFocus(focusElements[currentFocus])
			return nil
		}

		// 5. Resize Logic
		if event.Modifiers() == tcell.ModShift {
			if event.Key() == tcell.KeyLeft && listWidth > 15 {
				listWidth -= 2
				bodyFlex.ResizeItem(commandList, listWidth, 0)
				return nil
			}
			if event.Key() == tcell.KeyRight && listWidth < 80 {
				listWidth += 2
				bodyFlex.ResizeItem(commandList, listWidth, 0)
				return nil
			}
		}

		// 6. Panel Switching
		switch event.Key() {
		case tcell.KeyRight:
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

	if err := app.SetRoot(layout, true).SetFocus(catBar).Run(); err != nil {
		panic(err)
	}
}
