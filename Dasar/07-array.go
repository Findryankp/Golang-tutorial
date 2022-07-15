package main

import "fmt"

func main() {
	//Manual
	var arrayNames [3]string

	arrayNames[0] = "Findryan"
	arrayNames[1] = "Kurnia"
	arrayNames[2] = "Pradana"

	fmt.Println(arrayNames[0])
	fmt.Println(arrayNames[1])
	fmt.Println(arrayNames[2])

	//Langsung
	var arrayNilai = [...]int{
		98,
		90,
		92,
	}

	fmt.Println(arrayNilai[0])
	fmt.Println(arrayNilai[1])
	fmt.Println(arrayNilai[2])

	fmt.Println(len(arrayNames))
	fmt.Println(len(arrayNilai))

	// var arrayNilai [...]string
}
