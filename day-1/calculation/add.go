package calculation

// function dengan huruf depan kapital bisa dipanggil di package yang berbeda
func Add(number int, numberTwo int) int {
	return add(number, numberTwo)
}

// function dengan huruf depan kecil cuma bisa dipanggil di package yang sama
func add(number int, numberTwo int) int {
	return number + numberTwo
}
