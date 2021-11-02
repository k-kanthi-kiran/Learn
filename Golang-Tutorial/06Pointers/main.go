package main

import "fmt"

func main() {

	fmt.Println("Hi for Pointers")
	var number int = 23
	var ptr = &number
	fmt.Println("Address of  number ", ptr)
	fmt.Println("Value of  number ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Updated number value:", number)
}
