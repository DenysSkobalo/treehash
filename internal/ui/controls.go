package ui

import (
	"strconv"
	"treehash/internal/tree"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Controls(w fyne.Window, root **tree.TreeNode, redraw func()) *fyne.Container {
	status := widget.NewLabel("Ready")

	addBtn := widget.NewButton("Add Node", func() {
		valEntry := widget.NewEntry()
		valEntry.SetPlaceHolder("Enter value (int)")

		sideSelect := widget.NewSelect([]string{"Left", "Right"}, func(s string) {})
		sideSelect.PlaceHolder = "Choose side"

		form := &widget.Form{
			Items: []*widget.FormItem{
				{Text: "Value", Widget: valEntry},
				{Text: "Side", Widget: sideSelect},
			},
			OnSubmit: func() {
				val, err := strconv.Atoi(valEntry.Text)
				if err != nil || sideSelect.Selected == "" {
					status.SetText("Invalid input")
					return
				}
				side := 0
				if sideSelect.Selected == "Right" {
					side = 1
				}
				*root = tree.Insert(val, side, *root)
				status.SetText("Inserted: " + valEntry.Text + " (" + sideSelect.Selected + ")")
				redraw()
			},
		}

		dialog.ShowCustom("Add Node", "Close", form, w)
	})

	resetBtn := widget.NewButton("Reset", func() {
		*root = nil
		status.SetText("Reset")
		redraw()
	})

	return container.NewHBox(addBtn, resetBtn, widget.NewSeparator(), status)
}
