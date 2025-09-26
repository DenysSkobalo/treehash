package main

import (
	"treehash/internal/tree"
	"treehash/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("ArborHash – Step by Step")
	w.Resize(fyne.NewSize(900, 900))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	var root *tree.TreeNode
	scene := container.NewWithoutLayout()
	canvasWrap := container.NewMax(scene)

	redraw := func() {
		scene.Objects = nil
		scene.Add(ui.BuildScene(root, w, 900))
		scene.Refresh()
	}

	controls := ui.Controls(w, &root, redraw)
	content := container.NewBorder(controls, nil, nil, nil, canvasWrap)

	w.SetContent(content)
	w.ShowAndRun()
}
