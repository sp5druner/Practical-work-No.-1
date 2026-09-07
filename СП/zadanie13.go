package main

import "fmt"

type Lesson struct {
	Subject string
	Room    int
	Teacher string
}

func zadanie13() {
	const (
		Monday    = "Понедельник"
		Wednesday = "Среда"
	)

	schedule := map[string][]Lesson{
		Monday: {
			{Subject: "ТВиМС", Room: 41, Teacher: "Сапожникова"},
			{Subject: "ФГ", Room: 33, Teacher: "Непоспехова"},
			{Subject: "СП", Room: 21, Teacher: "Лобанова"},
		},
		Wednesday: {
			{Subject: "ТВиМС", Room: 41, Teacher: "Сапожникова"},
			{Subject: "ТВиМС", Room: 41, Teacher: "Сапожникова"},
			{Subject: "МПД", Room: 33, Teacher: "Непоспехова"},
		},
	}

	fmt.Println("Расписание на Понедельник")
	for _, lesson := range schedule[Monday] {
		fmt.Printf("  Предмет: %s, Кабинет: %d, Преподаватель: %s\n",
			lesson.Subject, lesson.Room, lesson.Teacher)
	}

	fmt.Println("\n Расписание на Среду")
	for _, lesson := range schedule[Wednesday] {
		fmt.Printf("  Предмет: %s, Кабинет: %d, Преподаватель: %s\n",
			lesson.Subject, lesson.Room, lesson.Teacher)
	}
}
