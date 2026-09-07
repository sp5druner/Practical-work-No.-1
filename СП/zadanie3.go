package main

import "fmt"

func zadanie3() {
	disk := 5000.0
	file := 256.0
	otvet1 := int(disk / file)
	otvet2 := disk - float64(otvet1)*file
	fmt.Printf("можно поместить %d файлов\n", otvet1)
	fmt.Printf("останется %.2f гб", otvet2)
}
