package ui

import (
	"fmt"
	"image/color"
	"sort"
	"treehash/internal/tree"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type coord struct {
	x, y float32
}

// --- Draw the tree ---
func BuildScene(root *tree.TreeNode, w fyne.Window, width float32) *fyne.Container {
	c := container.NewWithoutLayout()
	if root == nil {
		return c
	}

	indexLevel, total, _ := tree.IndexAndLevel(root)

	margin := float32(40)
	radius := float32(18)
	xGap := (width - 2*margin) / float32(total+1)
	levelGap := float32(100)

	coords := make(map[*tree.TreeNode]coord)
	for n, il := range indexLevel {
		idx := il[0]
		lev := il[1]
		x := margin + float32(idx)*xGap
		y := margin + float32(lev)*levelGap
		coords[n] = coord{x: x, y: y}
	}

	var nodes []*tree.TreeNode
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
