package main

import "fmt"

func zadanie9() {
	var sum float64
	fmt.Print("Введите сумму покупки: ")
	fmt.Scan(&sum)
	count := sum * 0.20
	total := sum - count
	fmt.Printf("Скидка: %.2f руб\n", count)
	fmt.Printf("Итоговая сумма: %.2f руб\n", total)
}
