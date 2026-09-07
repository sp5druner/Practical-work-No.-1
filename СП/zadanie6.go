package main

import (
	"fmt"
	"math"
)

func zadanie6() {
	var defoltsum, stavka, years float64
	fmt.Print("введите начальную сумму : ")
	fmt.Scan(&defoltsum)
	fmt.Print("введите ставку : ")
	fmt.Scan(&stavka)
	fmt.Print("введите кол во лет : ")
	fmt.Scan(&years)
	result := defoltsum * math.Pow(1+stavka/100, years)
	fmt.Printf("Итоговая сумма: %.2f\n", result)
}
