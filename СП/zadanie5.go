package main

import (
	"fmt"
	"math"
)

func zadanie5() {
	var radius float64
	fmt.Print("введите радиус ")
	fmt.Scan(&radius)
	L := 2 * math.Pi * radius
	S := math.Pi * radius * radius
	fmt.Printf("длина : %.2f\n", L)
	fmt.Printf("площадь : %.2f", S)
}
