package main

import "fmt"

func main() {
	bulan := []string{
		"Jan", "Feb", "Mar", "Apr",
		"Mei", "Jun", "Jul", "Agu",
		"Sep", "Okt", "Nov", "Des",
	}

	slice1 := bulan[4:8]
	fmt.Println(slice1)

	len := len(slice1)
	cap := cap(slice1)
	fmt.Println("Len : ", len)
	fmt.Println("Cap : ", cap)

	slice2 := bulan[2:]
	fmt.Println(slice2)

	slice3 := bulan[:8]
	fmt.Println(slice3)

	slice4 := bulan[:]
	fmt.Println(slice4)
}
