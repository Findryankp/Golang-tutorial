package main

import "fmt"

func main() {
	nilai := 70

	switch {
	case nilai >= 100:
		fmt.Println("nilai excelent")
	case nilai < 100 && nilai > 90:
		fmt.Println("Hebat")
	default:
		fmt.Println("Cukup")
	}
}
