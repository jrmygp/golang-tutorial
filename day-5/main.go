package main

import "fmt"

// Array - Slice - Map
func main() {
	// ===================== ARRAY =====================
	var languages [5]int
	// masukin value ke dalam array menggunakan index
	languages[0] = 12

	// mendapatkan length dari array
	// arrayLength := len(languages)

	// fmt.Println(languages, arrayLength)

	// cara ngedeclare array langsung tanpa var dan langsung diisi
	// pake ... supaya ga usah hardcode jumlah isinya
	// bahasa := [...]string{"Ruby", "Python", "JavaScript", "C++", "Java"}

	// fmt.Println(bahasa)

	// for i := 0; i < len(bahasa); i++ {
	// 	fmt.Println("index", i, "adalah", bahasa[i])
	// }

	// for index, lang := range bahasa {
	// 	fmt.Println("index", index, "adalah", lang)
	// }

	// ===================== SLICE =====================
	// Slice digunakan ketika kita butuh sebuah collection yang lengthnya bersifat dinamis
	// Cara declarenya mirip dengan array, tapi di dalem kurung siku ga usah dikasih apa"
	// Declare slice collection :
	var gamingConsoles []string

	// Cara menambahkan element ke dalam slice
	gamingConsoles = append(gamingConsoles, "Xbox 1")

	// fmt.Println(gamingConsoles)

	// ===================== MAP =====================
	// Ngedeclare new map
	var myMap map[string]int
	// Ngeinitialize mapnya (harus karena ketika first initialize map diatas default valuenya nil)
	myMap = map[string]int{}

	myMap["ruby"] = 9
	myMap["go"] = 9
	myMap["javascript"] = 9

	// fmt.Println((myMap["ruby"]))

	// Cara bikin map yang langsung diisi elementnya
	myLanguages := map[string]string{
		"ruby":       "is the worst language",
		"javascript": "is the best language",
		"go":         "is the hardest language",
	}

	// fmt.Println(myLanguages)

	// for key, value := range myLanguages {
	// 	fmt.Println(key, value)
	// }

	// delete element dari Map
	delete(myLanguages, "go")

	// for key, value := range myLanguages {
	// 	fmt.Println(key, value)
	// }

	// cara tau jika sebuah Map memiliki element tertentu
	// value, isAvailable := myMap["python"]
	// // isAvailable tidak mandatory dan hanya akan menunjukan true/false
	// fmt.Println(value)
	// fmt.Println(isAvailable)

	// ===================== SLICE + MAP =====================
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Abdul", "score": "B"},
		{"name": "Somad", "score": "C"},
		{"name": "Jihan", "score": "C"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}
}
