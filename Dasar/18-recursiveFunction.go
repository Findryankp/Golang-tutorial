/*
	Fungsi yang memanggil dirinya sendiri
*/

package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func fibonaci(n int) int {
	if n < 2 {
		return n
	}

	return fibonaci(n-1) + fibonaci(n-2)
}

func main() {
	fmt.Println(faktorial(7))
	fmt.Println(fibonaci(7))
}
