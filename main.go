package main

import (
	"github.com/faymn-dev/maze-solver/internals/maze"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(maze.ScreenWidth, maze.ScreenHeight, "maze-solver")
	defer rl.CloseWindow()

	maze := maze.NewMaze(42)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		maze.Draw()

		rl.EndDrawing()
	}
}
