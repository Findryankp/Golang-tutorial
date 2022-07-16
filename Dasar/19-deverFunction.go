/*
	Dever : Fungsi yang bisa dijadwalkan untuk berjalan setelah sebuah fungsi selesai di eksekusi (*walaupun eror tetap dieksekusi)
*/

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("RUN Application")

	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApplication(0)
}
