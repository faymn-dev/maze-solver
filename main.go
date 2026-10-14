package main

import (
	"github.com/faymn-dev/maze-solver/internals/maze"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(maze.ScreenWidth, maze.ScreenHeight, "maze-solver")
	defer rl.CloseWindow()

	maze := maze.NewMaze(42)

	path, explored := maze.Solve()

	var timer float32 = 0
	currentLine := 0

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		timer += rl.GetFrameTime()
		if timer > 0.005 {
			currentLine = min(currentLine+1, len(explored))
			timer = 0
		}

		maze.Draw()

		if currentLine == len(explored) {
			for i := range currentLine {
				explored[i].Draw(rl.LightGray)
			}

			// if "finished" exploring, then draw path
			for i := range len(path) - 1 {
				cell := path[i]
				nextCell := path[i+1]
				cell.DrawMove(nextCell)
			}
		} else {
			for i := range currentLine {
				explored[i].Draw(rl.Gray)
			}
		}

		rl.EndDrawing()
	}
}
