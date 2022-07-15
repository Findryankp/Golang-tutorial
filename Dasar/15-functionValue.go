package main

import "fmt"

func printHello() string {
	return "Hello"
}

func main() {
	varFunc := printHello

	fmt.Println(varFunc())
}
