package main

import "fmt"

func main() {

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Guava"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Length is :", len(fruitList))

	var vegList = [5]string{"brinjal", "beans", "potato"}
	fmt.Println("vegList list is :", vegList)
}
