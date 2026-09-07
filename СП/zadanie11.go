package main

import "fmt"

type UserProfile struct {
	Username     string
	Age          int
	FriendsCount int
	IsVerified   bool
	Rating       float64
}

func zadanie11() {
	user1 := UserProfile{
		Username:     "Egor",
		Age:          17,
		FriendsCount: 150,
		IsVerified:   true,
		Rating:       4.2,
	}

	user2 := UserProfile{
		Username:     "Ivan",
		Age:          18,
		FriendsCount: 80,
		IsVerified:   false,
		Rating:       4.3,
	}

	user3 := UserProfile{
		Username:     "Gregor",
		Age:          22,
		FriendsCount: 200,
		IsVerified:   true,
		Rating:       4.9,
	}
	fmt.Println("Профили пользователей")
	fmt.Printf("Имя: %s, Возраст: %d, Друзей: %d, Верифицирован: %t, Рейтинг: %.1f\n",
		user1.Username, user1.Age, user1.FriendsCount, user1.IsVerified, user1.Rating)
	fmt.Printf("Имя: %s, Возраст: %d, Друзей: %d, Верифицирован: %t, Рейтинг: %.1f\n",
		user2.Username, user2.Age, user2.FriendsCount, user2.IsVerified, user2.Rating)
	fmt.Printf("Имя: %s, Возраст: %d, Друзей: %d, Верифицирован: %t, Рейтинг: %.1f\n",
		user3.Username, user3.Age, user3.FriendsCount, user3.IsVerified, user3.Rating)
}
