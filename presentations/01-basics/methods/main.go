package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y int }

func (p Point) Distance() float64 { // ресивер-значение: работает с копией
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func main() {
	p := Point{X: 3, Y: 4}
	fmt.Println(p.Distance()) // 5
}
