/*
	Struct : sebuah template yang digunakan untuk menggabungkan berbagai tipe data dalam satu kesatuan
*/

package main

import "fmt"

type Pegawai struct {
	Name, Address string
	Age           int
}

func main() {
	var findryan Pegawai
	findryan.Name = "Findryan Kurnia Pradana"
	findryan.Address = "Madiun"
	findryan.Age = 26

	dian := Pegawai{
		Name:    "Dian",
		Address: "Madiun",
		Age:     25,
	}

	alfian := Pegawai{"Alfian", "Madiun", 22}

	fmt.Println(findryan)
	fmt.Println(dian)
	fmt.Println(alfian)
}
