package main

import (
	"fmt"
	"math"
)

func zadanie10() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	result := float64(a) / float64(b)
	rounded := math.Round(result)
	floored := math.Floor(result)
	fmt.Printf("Результат: %.2f\n", result)
	fmt.Printf("Округление до ближайшего: %.0f\n", rounded)
	fmt.Printf("Округление вниз: %.0f\n", floored)
}
