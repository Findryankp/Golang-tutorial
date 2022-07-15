package main

import (
	"fmt"
)

func main() {
	//1. Number
	var noInduk int
	noInduk = 2
	fmt.Println("Tipe data number : ", noInduk)

	//2. Bool
	var isMale bool
	isMale = true
	fmt.Println("Tipe data bool : ", isMale)

	//3. String
	var nama string
	nama = "Findryan Kurnia Pradana"
	fmt.Println("Tipe data string : ", nama)

	//4. Float
	var ipk float32
	ipk = 2.4
	fmt.Println("Tipe data string : ", ipk)

	//5. Shortcut variabel
	alamatSatu := "Jalan Raya No. "
	alamatDua := 34
	fmt.Println("Alamat : ", alamatSatu, alamatDua)

	//6. Constant
	const noKtp = "357702000000000"
	fmt.Println("No KTP (const) : ", noKtp)

}
