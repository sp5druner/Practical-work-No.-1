package main

import "fmt"

func zadanie2() {
	notebook := 55480.0
	monitor := 21830.0
	mouse := 890.0
	keyboard := 1560.0
	vsego := notebook*6 + monitor*3 + mouse*11 + keyboard*5
	fmt.Printf("Нужно выделить : %.2f руб", vsego)
}
