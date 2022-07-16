package main

import "fmt"

type Blacklist func(string) bool

func registeredUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//function tanpa nama
	blacklist := func(name string) bool {
		return name == "Findryan"
	}

	registeredUser("Findryan", blacklist)
	registeredUser("Kurnia", blacklist)

	registeredUser("Pradana", func(name string) bool {
		return name == "Findryan"
	})
}
