package maze

type Point struct {
	x int32
	y int32
}

func NewPoint(x int32, y int32) *Point {
	return &Point{x, y}
}
