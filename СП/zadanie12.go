package main

import "fmt"

func zadanie12() {
	const (
		Online       = "онлайн"
		Offline      = "офлайн"
		Away         = "отошёл"
		DoNotDisturb = "не беспокоить"
	)

	users := map[string]string{
		"student1": Online,
		"student2": Offline,
		"student3": Away,
		"student4": DoNotDisturb,
	}

	fmt.Println("Текущие статусы")
	for name, status := range users {
		fmt.Printf("%s: %s\n", name, status)
	}

	users["student2"] = Offline
	users["student1"] = Online
	users["student4"] = Online

	fmt.Println("\n=== После обновления ===")
	for name, status := range users {
		fmt.Printf("%s: %s\n", name, status)
	}

	fmt.Println("\n=== Пользователи онлайн ===")
	for name, status := range users {
		if status == Online {
			fmt.Printf("%s\n", name)
		}
	}
}
