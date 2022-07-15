package main

import "fmt"

func main() {
	//type declaration is declaration for make alias name from a typeData
	type aliasString string
	type aliasInt int

	var nama aliasString
	nama = "Findryan Kurnia Pradana"
	fmt.Println("Alias String : ", nama)

	var noHp aliasInt
	noHp = 81000000
	fmt.Println("Alias INT : ", noHp)
}
