package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func zadanie8() {
	point1 := Point{X: 1.0, Y: 2.0}
	point2 := Point{X: 4.0, Y: 6.0}
	point3 := Point{X: -1.0, Y: -3.0}
	point4 := Point{X: 5.0, Y: 1.0}
	distance12 := math.Sqrt(math.Pow(point2.X-point1.X, 2) + math.Pow(point2.Y-point1.Y, 2))
	distance34 := math.Sqrt(math.Pow(point4.X-point3.X, 2) + math.Pow(point4.Y-point3.Y, 2))
	distance13 := math.Sqrt(math.Pow(point3.X-point1.X, 2) + math.Pow(point3.Y-point1.Y, 2))
	fmt.Printf("Расстояние между точкой 1 (%.2f, %.2f) и точкой 2 (%.2f, %.2f) = %.2f\n",
		point1.X, point1.Y, point2.X, point2.Y, distance12)
	fmt.Printf("Расстояние между точкой 3 (%.2f, %.2f) и точкой 4 (%.2f, %.2f) = %.2f\n",
		point3.X, point3.Y, point4.X, point4.Y, distance34)
	fmt.Printf("Расстояние между точкой 1 (%.2f, %.2f) и точкой 3 (%.2f, %.2f) = %.2f\n",
		point1.X, point1.Y, point3.X, point3.Y, distance13)
}
