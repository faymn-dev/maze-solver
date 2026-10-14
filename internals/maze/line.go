package maze

import rl "github.com/gen2brain/raylib-go/raylib"

type Line struct {
	p1 *Point
	p2 *Point
}

func NewLine(p1 *Point, p2 *Point) *Line {
	return &Line{p1, p2}
}

func (l *Line) Draw(color rl.Color) {
	// https://pkg.go.dev/github.com/gen2brain/raylib-go/raylib#DrawLine
	rl.DrawLine(l.p1.x, l.p1.y, l.p2.x, l.p2.y, color)
}
