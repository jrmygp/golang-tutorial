package main

import "fmt"

// Variables
func main() {
	// membuat variable yang diisinya bisa nanti saja
	var name string
	name = "Jeremy"

	// membuat variable yang langsung ada isinya seperti diatas bisa langsung kayak gini :
	kalimat := "Hello World"

	// untuk mengganti value dari variable yang udah di declare sebelumnya
	kalimat = "Halo Dunia"

	fmt.Println(name, kalimat)
}
