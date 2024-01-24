package main

import "fmt"

func main() {
	// & = referencing
	// * = dereferencing
	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 20

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	number := 5
	// mengubah number menjadi 100
	fmt.Println("alamat memory:", &number)
	fmt.Println("awal:", number)

	change(&number, 100)

	fmt.Println("alamat memory:", &number)
	fmt.Println("akhir:", number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("alamat memory:", old)
	fmt.Println("Di dalam :", *old)
}
