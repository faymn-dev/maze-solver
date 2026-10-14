package maze

import rl "github.com/gen2brain/raylib-go/raylib"

type Cell struct {
	hasLeftWall   bool
	hasRightWall  bool
	hasTopWall    bool
	hasBottomWall bool
	center        *Point
	visited       bool
	id            *Point
}

func NewCell(center *Point, id *Point) *Cell {
	return &Cell{
		hasLeftWall:   true,
		hasRightWall:  true,
		hasTopWall:    true,
		hasBottomWall: true,
		center:        center,
		id:            id,
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

func (c *Cell) DrawMove(neighbor *Cell, undo bool) {
	color := rl.Red
	if undo {
		color = rl.LightGray
	}
	NewLine(c.center, neighbor.center).Draw(color)
}

func (c *Cell) CanGoTo(neighbor *Cell) bool {
	if neighbor.id.x > c.id.x {
		// neighbor is to the right
		return !c.hasRightWall && !neighbor.hasLeftWall
	} else if neighbor.id.x < c.id.x {
		// neighbor is to the left
		return !c.hasLeftWall && !neighbor.hasRightWall
	} else if neighbor.id.y > c.id.y {
		// neighbor is to the bottom
		return !c.hasBottomWall && !neighbor.hasTopWall
	} else if neighbor.id.y < c.id.y {
		// neighbor is to the top
		return !c.hasTopWall && !neighbor.hasBottomWall
	}

	return false
}
