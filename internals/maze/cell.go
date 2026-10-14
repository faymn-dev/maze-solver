package maze

import rl "github.com/gen2brain/raylib-go/raylib"

type Cell struct {
	hasLeftWall   bool
	hasRightWall  bool
	hasTopWall    bool
	hasBottomWall bool
	center        *Point
	visited       bool
}

func NewCell(center *Point) *Cell {
	return &Cell{
		hasLeftWall:   true,
		hasRightWall:  true,
		hasTopWall:    true,
		hasBottomWall: true,
		center:        center,
	}
}

func (c *Cell) Draw() {
	x1 := c.center.x - CellSize/2
	x2 := c.center.x + CellSize/2
	y1 := c.center.y - CellSize/2
	y2 := c.center.y + CellSize/2

	if c.hasLeftWall {
		NewLine(NewPoint(x1, y1), NewPoint(x1, y2)).Draw(rl.Black)
	}
	if c.hasRightWall {
		NewLine(NewPoint(x2, y1), NewPoint(x2, y2)).Draw(rl.Black)
	}
	if c.hasTopWall {
		NewLine(NewPoint(x1, y1), NewPoint(x2, y1)).Draw(rl.Black)
	}
	if c.hasBottomWall {
		NewLine(NewPoint(x1, y2), NewPoint(x2, y2)).Draw(rl.Black)
	}
}

func (c *Cell) DrawMove(toCell *Cell, undo bool) {
	color := rl.Red
	if undo {
		color = rl.LightGray
	}
	NewLine(c.center, toCell.center).Draw(color)
}
