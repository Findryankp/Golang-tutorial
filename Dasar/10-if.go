package main

import "fmt"

func main() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("A > B")
	} else if a < b {
		fmt.Println("A < B")
	} else {
		fmt.Println("Tidak Keduanya")
	}

}
