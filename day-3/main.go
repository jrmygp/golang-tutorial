package main

import "fmt"

// If chain
func main() {
	age := 70

	if age < 30 {
		fmt.Println("Umur anda kurang dari 30")
	} else {
		fmt.Println("Umur anda sudah pas")
	}

	var keterangan string

	if age <= 30 {
		keterangan = "kamu masih mudah"
	} else if age > 30 && age <= 50 {
		keterangan = "kamu sudah agak tua"
	} else {
		keterangan = "kamu sudah tua"
	}

	fmt.Println(keterangan)
}
