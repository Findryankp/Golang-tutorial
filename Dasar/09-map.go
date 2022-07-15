package main

import "fmt"

func main() {
	//map[tipe key nya][tipe valuenya]
	person := map[string]string{
		"nama": "findryan",
		"nohp": "819400000",
	}

	fmt.Println(person)

	book := make(map[string]string)
	book["judul"] = "Tutorial Golang"
	book["author"] = "Findryan"
	fmt.Println(book)
	fmt.Println(book["judul"])

	delete(book, "author")
	fmt.Println(book)
}
