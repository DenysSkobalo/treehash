package main

import (
	"fmt"
	"treehash/internal/tree"
	"treehash/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("ArborHash – Step by Step")
	w.Resize(fyne.NewSize(900, 900))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	hash := map[int][]int{
		0: {0, 10},
		1: {1, 20},
		2: {1, 30},
		3: {0, 40},
	}
	keys := []int{0, 1, 2, 3}

	step := 0
	var root *tree.TreeNode

	status := widget.NewLabel("Ready")
	scene := container.NewWithoutLayout()
	canvasWrap := container.NewMax(scene)

	redraw := func() {
		scene.Objects = nil
		scene.Add(ui.BuildScene(root, w, 900))
		scene.Refresh()
	}

	nextBtn := widget.NewButton("Next", func() {
		if step >= len(keys) {
			status.SetText("All values inserted.")
			return
		}
		k := keys[step]
		arr := hash[k]
		side, val := arr[0], arr[1]
		step++
		root = tree.Insert(val, side, root)
		status.SetText(fmt.Sprintf("Inserted: %d side=%d (%d/%d)", val, side, step, len(keys)))
		redraw()
	})

	resetBtn := widget.NewButton("Reset", func() {
		root = nil
		step = 0
		status.SetText("Reset")
		redraw()
	})

	btns := container.NewHBox(nextBtn, resetBtn, widget.NewSeparator(), status)
	content := container.NewBorder(btns, nil, nil, nil, canvasWrap)

	w.SetContent(content)
	w.ShowAndRun()
}
