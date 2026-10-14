package maze

import (
	"math/rand"
)

type Cells [][]*Cell

type Maze struct {
	rand  *rand.Rand
	cells Cells
}

func NewMaze(seed int64) *Maze {
	maze := &Maze{
		rand:  rand.New(rand.NewSource(seed)),
		cells: createCells(),
	}

	maze.breakWalls(0, 0)
	maze.resetVisited()

	return maze
}

func (m *Maze) Draw() {
	for _, col := range m.cells {
		for _, cell := range col {
			cell.Draw()
		}
	}
}

var neighborOffsets = []*Point{NewPoint(-1, 0), NewPoint(1, 0), NewPoint(0, -1), NewPoint(0, 1)}

func (m *Maze) resetVisited() {
	for _, col := range m.cells {
		for _, cell := range col {
			cell.visited = false
		}
	}
}

func (m *Maze) breakWalls(i int32, j int32) {
	cell := m.getCell(i, j)
	if cell == nil {
		return
	}

	cell.visited = true

	for {
		neighbors := []*Cell{}
		neighborIDs := []*Point{}
		for _, offset := range neighborOffsets {
			x := i + offset.x
			y := j + offset.y

			neighbor := m.getCell(x, y)
			if neighbor != nil && !neighbor.visited {
				neighbors = append(neighbors, neighbor)
				neighborIDs = append(neighborIDs, NewPoint(x, y))
			}
		}

		if len(neighbors) == 0 {
			return
		}

		neighborIdx := m.rand.Intn(len(neighbors))
		neighbor := neighbors[neighborIdx]
		neighborID := neighborIDs[neighborIdx]

		// knock down walls between current cell and neighbor
		if neighborID.x > i {
			// neighbor is to the right
			cell.hasRightWall = false
			neighbor.hasLeftWall = false
		} else if neighborID.x < i {
			// neighbor is to the left
			cell.hasLeftWall = false
			neighbor.hasRightWall = false
		} else if neighborID.y > j {
			// neighbor is to the bottom
			cell.hasBottomWall = false
			neighbor.hasTopWall = false
		} else if neighborID.y < j {
			// neighbor is to the top
			cell.hasTopWall = false
			neighbor.hasBottomWall = false
		}

		m.breakWalls(neighborID.x, neighborID.y)
	}
}

func (m *Maze) getCell(i int32, j int32) *Cell {
	if i >= 0 && i < NumCols && j >= 0 && j < NumRows {
		return m.cells[i][j]
	}
	return nil
}

func createCells() Cells {
	cells := make([][]*Cell, NumCols)

	mazeWidth := (CellSize * NumCols)
	mazeHeight := (CellSize * NumRows)

	for i := range NumCols {
		cells[i] = make([]*Cell, NumRows)

		for j := range NumRows {
			cells[i][j] = NewCell(NewPoint(
				// translate grid to the center
				// we add CellSize/2 because cells are drawn from the center, not the top left
				int32(i*CellSize+ScreenWidth/2-mazeWidth/2+CellSize/2),
				int32(j*CellSize+ScreenHeight/2-mazeHeight/2+CellSize/2),
			))
		}
	}

	// break entrance and exit
	cells[0][0].hasTopWall = false
	cells[NumCols-1][NumRows-1].hasBottomWall = false

	return cells
}
