package main

import "fmt"

func main() {
	counter := 1
	for counter < 10 {
		counter++

		if counter%2 == 0 {
			continue
		}

		fmt.Println(counter)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var arrayNilai = [...]int{
		98, 1,
		90, 2,
		92, 3,
	}

	for i, value := range arrayNilai {
		fmt.Println("index ke:", i, "=", value)
	}

	for _, value := range arrayNilai {
		fmt.Println(value)
	}
}
