package main

import "fmt"

func main() {
	// && --> AND
	// || --> OR

	a := 10
	b := 20

	c := a > b
	d := a < b

	fmt.Println("AND :", c && d)
	fmt.Println("OR :", c || d)
}
