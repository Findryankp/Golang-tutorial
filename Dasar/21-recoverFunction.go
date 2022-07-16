/*
	Recover : Fungsi yang bisa digunakan untuk menangkap data panic
	1. recover akan menghentikan proses panic sehingga program tetap berjalan
*/

package main

import "fmt"

func panicFunc() {
	message := recover()
	fmt.Println(message)
	fmt.Println("Program berhenti")
}

func runApp(error bool) {
	defer panicFunc()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
