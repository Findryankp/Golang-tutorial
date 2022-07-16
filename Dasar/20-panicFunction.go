/*
	Panic : Fungsi yang bisa digunakan untuk menghentikan program
	1. panic fungtion dipanggil, program berhenti, defer function tetap akan diekseskusi
*/

package main

import "fmt"

func panicFunc() {
	fmt.Println("Program berhenti")
}

func runApp(error bool) {
	defer panicFunc()
	if error {
		panic("Aplikasi error")
	}
}

func main() {
	runApp(true)
}
