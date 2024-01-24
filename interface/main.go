package main

import "fmt"

// interface seperti kontrak yang harus dipenuhi / aturan main
// seperti interface di typescript

// cara buat sebuah interface
type BangunDatar interface {
	// supaya memenuhi kontrak dia harus punya sebuah method bernama HitungLuas
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

// persegi sudah memenuhi kontrak karena persegi sudah punya sebuah method bernama HitungLuas
func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

// persegi panjang juga sama
func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 10, Lebar: 8}
	persegi := Persegi{Sisi: 4}

	luasPersegiPanjang := SeberapaLuas(persegiPanjang)
	fmt.Println(luasPersegiPanjang)

	luasPersegi := SeberapaLuas(persegi)
	fmt.Println(luasPersegi)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
