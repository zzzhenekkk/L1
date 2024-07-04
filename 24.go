package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func DistanceBetweenPoints(p1, p2 *Point) float64 {
	x := p2.x - p1.x
	y := p2.y - p1.y
	return math.Sqrt(x*x + y*y)
}

func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	a := NewPoint(1.25, 1)
	b := NewPoint(3.25, 1)
	dist := DistanceBetweenPoints(a, b)
	dist2 := b.Distance(a)

	fmt.Println("1) Distance between points: ", dist)
	fmt.Println("2) Distance between points: ", dist2)
}
