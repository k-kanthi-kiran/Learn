package main

import "fmt"

func main() {

	var loginCount int
	var result string

	loginCount = 10

	if loginCount < 10 {
		result = "login count less than 10"
	} else if loginCount > 10 {
		result = " login count greater than 10"
	} else {
		result = " login count exactly 10"
	}

	fmt.Println(result)

	if numb := 13; numb < 10 {
		result = "Number is less than 10"
	} else {
		result = "Number is more than 10"
	}

	fmt.Println(result)
}
