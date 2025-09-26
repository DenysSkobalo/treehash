package main

import (
	"fmt"
	"image/color"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- Your custom Insert with side flag ---
func Insert(val int, side int, node *TreeNode) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	if side == 0 {
		node.Left = Insert(val, side, node.Left)
	} else {
		node.Right = Insert(val, side, node.Right)
	}
	return node
}

// --- Assign in-order indices and depths ---
func indexAndLevel(root *TreeNode) (map[*TreeNode][2]int, int, int) {
	res := make(map[*TreeNode][2]int)
	idx := 0
	maxLev := 0

	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, lev int) {
		if n == nil {
			return
		}
		if lev > maxLev {
			maxLev = lev
		}
		dfs(n.Left, lev+1)
		idx++
		res[n] = [2]int{idx, lev}
		dfs(n.Right, lev+1)
	}
	dfs(root, 0)
	return res, idx, maxLev
}

type coord struct {
	x, y float32
}

// --- Draw the tree ---
func buildScene(root *TreeNode, w fyne.Window, width float32) *fyne.Container {
	c := container.NewWithoutLayout()
	if root == nil {
		return c
	}

	indexLevel, total, _ := indexAndLevel(root)

	margin := float32(40)
	radius := float32(18)
	xGap := (width - 2*margin) / float32(total+1)
	levelGap := float32(100)

	coords := make(map[*TreeNode]coord)
	for n, il := range indexLevel {
		idx := il[0]
		lev := il[1]
		x := margin + float32(idx)*xGap
		y := margin + float32(lev)*levelGap
		coords[n] = coord{x: x, y: y}
	}

	var nodes []*TreeNode
	for n := range coords {
		nodes = append(nodes, n)
	}
	sort.Slice(nodes, func(i, j int) bool { return nodes[i].Val < nodes[j].Val })

	// edges
	for _, n := range nodes {
		p := coords[n]
		if n.Left != nil {
			ch := coords[n.Left]
			line := canvas.NewLine(color.NRGBA{R: 120, G: 120, B: 120, A: 255})
			line.StrokeWidth = 2
			line.Position1 = fyne.NewPos(p.x, p.y)
			line.Position2 = fyne.NewPos(ch.x, ch.y)
			c.Add(line)
		}
		if n.Right != nil {
			ch := coords[n.Right]
			line := canvas.NewLine(color.NRGBA{R: 120, G: 120, B: 120, A: 255})
			line.StrokeWidth = 2
			line.Position1 = fyne.NewPos(p.x, p.y)
			line.Position2 = fyne.NewPos(ch.x, ch.y)
			c.Add(line)
		}
	}

	// nodes
	for _, n := range nodes {
		p := coords[n]
		circ := canvas.NewCircle(color.NRGBA{R: 245, G: 245, B: 245, A: 255})
		circ.StrokeColor = color.NRGBA{R: 30, G: 30, B: 30, A: 255}
		circ.StrokeWidth = 2
		circ.Resize(fyne.NewSize(radius*2, radius*2))
		circ.Move(fyne.NewPos(p.x-radius, p.y-radius))
		c.Add(circ)

		txt := canvas.NewText(fmt.Sprintf("%d", n.Val), color.NRGBA{R: 20, G: 20, B: 20, A: 255})
		txt.TextSize = 14
		size := txt.MinSize()
		txt.Move(fyne.NewPos(p.x-size.Width/2, p.y-size.Height/2))
		c.Add(txt)
	}

	return c
}

func main() {
	a := app.New()
	w := a.NewWindow("Custom Hash Tree – Step by Step")
	w.Resize(fyne.NewSize(900, 900))
	w.SetFixedSize(true)
	w.CenterOnScreen()

	// Hash input: side, value
	hash := map[int][]int{
		0: {0, 10},
		1: {1, 20},
		2: {1, 30},
		3: {0, 40},
	}

	keys := []int{0, 1, 2, 3} // order to insert
	step := 0
	var root *TreeNode

	status := widget.NewLabel("Ready")
	scene := container.NewWithoutLayout()
	canvasWrap := container.NewMax(scene)

	redraw := func() {
		scene.Objects = nil
		scene.Add(buildScene(root, w, 900))
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
		root = Insert(val, side, root)
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
