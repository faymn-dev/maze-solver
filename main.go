package main

import (
	"github.com/faymn-dev/maze-solver/internals/maze"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(maze.ScreenWidth, maze.ScreenHeight, "maze-solver")
	defer rl.CloseWindow()

	maze := maze.NewMaze(42)

	path := maze.Solve()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		maze.Draw()
		for i := range len(path) - 1 {
			cell := path[i]
			nextCell := path[i+1]
			cell.DrawMove(nextCell, false)
		}

		rl.EndDrawing()
	}
}
