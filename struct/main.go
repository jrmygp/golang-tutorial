package main

import (
	"fmt"
	"struct/management"
)

func main() {
	// Cara buat user dari sebuah struct
	user := management.User{
		ID:        1,
		FirstName: "Jeremy",
		LastName:  "Puglisi",
		Email:     "jeremy@kolabora.com",
		IsActive:  true,
	}

	// Cara lain buat user dari sebuah struct
	user2 := management.User{2, "Gipen", "Surepen", "pepen@surepen.com", true}

	// fmt.Println(displayUser(user))

	// // Buat group baru
	gamers := []management.User{user, user2}
	group := management.Group{"Gamer", user, gamers, true}

	// displayGroup(group)

	// memanggil sebuah method
	result := user.Display()
	fmt.Println(result)
	fmt.Println(user2.Display())

	group.DisplayGroup()
}

// func displayUser(user User) string {
// 	// Kalo sprintf itu dia nyimpen hasil ke dalem variable
// 	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)

// }

// func displayGroup(group Group) {
// 	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count: %d", len(group.Users))
// 	fmt.Println("")
// 	fmt.Println("Users:")
// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}
// }
