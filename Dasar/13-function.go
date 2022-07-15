package main

import "fmt"

//fungsi biasa
func sayHello() {
	fmt.Println("Hello Findryan")
}

//function with parameter
func sayTo(nama string, kalimat string) {
	fmt.Println(kalimat, nama)
}

//func with return value
func kalkulator(bil1 int, bil2 int, operator string) int {
	var result int
	if operator == "+" {
		result = bil1 + bil2
	} else if operator == "-" {
		result = bil1 - bil2
	} else if operator == "*" {
		result = bil1 * bil2
	} else if operator == "/" {
		result = bil1 / bil2
	}

	return result
}

//function return multiple values
func getFullName() (string, string, string) {
	return "Findryan", "Kurnia", "Pradana"
}

//function return multiple values
func getFullAddress() (city string, district string, number string) {
	city, district, number = "Madiun", "Tulungrejo", "29"
	return
}

func main() {
	sayHello()
	sayTo("Findryan", "Semangat")

	fmt.Println(kalkulator(56, 8, "/"))

	firstName, midName, lastName := getFullName()
	fmt.Println(firstName, midName, lastName)

	namaAwal, _, _ := getFullName()
	fmt.Println(namaAwal)

	city, _, _ := getFullAddress()
	_, district, _ := getFullAddress()
	_, _, nomor := getFullAddress()
	fmt.Println("Kota", city)
	fmt.Println("District", district)
	fmt.Println("Nomor", nomor)
}
