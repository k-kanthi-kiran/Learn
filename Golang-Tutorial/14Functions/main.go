package main

import "fmt"

func main() {
	fmt.Println("Sum is :", adder(2, 3))
	fmt.Println("Slice adding is :", multipleAdder(1, 2, 3, 4, 5, 6, 7, 8, 9))

	total, response := multipleAdderResponse(11, 22, 33, 44, 55, 66, 77, 88, 99)
	fmt.Println("Total is :", total)
	fmt.Println("response is :", response)
}

func adder(value1 int, value2 int) int {
	return value1 + value2
}

func multipleAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func multipleAdderResponse(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hello shambu"
}
