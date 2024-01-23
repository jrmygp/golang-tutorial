package main

import "fmt"

// Loop
func main() {
	// max := 10
	// for i := 0; i <= max; i++ {
	// 	fmt.Println(i)
	// }

	title := "Golang the best language"

	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Printf(string(letter))
	// 	}
	// }

	for _, letter := range title {
		if string(letter) == "a" || string(letter) == "i" || string(letter) == "u" || string(letter) == "e" || string(letter) == "o" {
			fmt.Printf(string(letter))
		}
	}
}
