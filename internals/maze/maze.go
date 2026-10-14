package maze

import (
	"math/rand"
)

var neighborOffsets = []*Point{NewPoint(-1, 0), NewPoint(1, 0), NewPoint(0, -1), NewPoint(0, 1)}

type Cells [][]*Cell

type Maze struct {
	rand     *rand.Rand
	cells    Cells
	path     []*Cell
	explored []*Line
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

func (m *Maze) Solve() ([]*Cell, []*Line) {
	if m.path != nil && m.explored != nil {
		return m.path, m.explored
	}

	m.path = nil
	m.explored = nil
	m.resetVisited()
	m.solveR(m.cells[0][0], nil)
	return m.path, m.explored
}

func (m *Maze) solveR(cell *Cell, prevCell *Cell) bool {
	cell.visited = true

	if prevCell != nil {
		m.explored = append(m.explored, NewLine(cell.center, prevCell.center))
	}

	// base case: are you at the end?
	if cell.id.x == NumCols-1 && cell.id.y == NumRows-1 {
		m.path = append(m.path, cell)
		return true
	}

	// try any neighbor
	for _, offset := range neighborOffsets {
		x := cell.id.x + offset.x
		y := cell.id.y + offset.y

		neighbor := m.getCell(x, y)
		if neighbor != nil && !neighbor.visited && cell.CanGoTo(neighbor) {
			result := m.solveR(neighbor, cell)
			if result {
				m.path = append(m.path, cell)
				return true
			}
		}
	}

	return false
}

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
		for _, offset := range neighborOffsets {
			x := i + offset.x
			y := j + offset.y

			neighbor := m.getCell(x, y)
			if neighbor != nil && !neighbor.visited {
				neighbors = append(neighbors, neighbor)
			}
		}

		if len(neighbors) == 0 {
			return
		}

		neighbor := neighbors[m.rand.Intn(len(neighbors))]

		// knock down walls between current cell and neighbor
		if neighbor.id.x > i {
			// neighbor is to the right
			cell.hasRightWall = false
			neighbor.hasLeftWall = false
		} else if neighbor.id.x < i {
			// neighbor is to the left
			cell.hasLeftWall = false
			neighbor.hasRightWall = false
		} else if neighbor.id.y > j {
			// neighbor is to the bottom
			cell.hasBottomWall = false
			neighbor.hasTopWall = false
		} else if neighbor.id.y < j {
			// neighbor is to the top
			cell.hasTopWall = false
			neighbor.hasBottomWall = false
		}

		m.breakWalls(neighbor.id.x, neighbor.id.y)
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
			cells[i][j] = NewCell(
				// translate grid to the center
				// we add CellSize/2 because cells are drawn from the center, not the top left
				NewPoint(
					int32(i*CellSize+ScreenWidth/2-mazeWidth/2+CellSize/2),
					int32(j*CellSize+ScreenHeight/2-mazeHeight/2+CellSize/2),
				),
				NewPoint(int32(i), int32(j)),
			)
		}
	}

	// break entrance and exit
	cells[0][0].hasTopWall = false
	cells[NumCols-1][NumRows-1].hasBottomWall = false

	return cells
}
