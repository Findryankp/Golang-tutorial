package main

import "fmt"

func sumInt(number ...int) int {
	total := 0

	for _, v := range number {
		total = total + v
	}

	return total
}

func main() {
	fmt.Println(sumInt(2, 3, 4, 5, 6, 7, 8))

	numbers := []int{2, 3, 4, 5, 6, 7, 8}
	total := sumInt(numbers...)
	fmt.Println(total)

}
