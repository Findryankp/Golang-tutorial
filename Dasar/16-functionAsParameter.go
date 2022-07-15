package main

import "fmt"

func hitungUsia(tahunLahir int, kriteria func(int) string) {
	usia := 2022 - tahunLahir
	fmt.Println("Usia : ", usia)

	varKriteria := kriteria(usia)
	fmt.Println("Hello", varKriteria)
}

func cekKriteria(usia int) string {
	if usia < 40 {
		return "Muda"
	} else {
		return "Tua"
	}
}

func main() {
	hitungUsia(1996, cekKriteria)
}
