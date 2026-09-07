package main

import "fmt"

func zadanie4() {
	var fahrenheit float64
	fmt.Print("Введите температуру в градусах Фаренгейта ")
	fmt.Scan(&fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, celsius)
}
