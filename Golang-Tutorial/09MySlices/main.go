package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println("Welcome to slices")
	var fruitList = []string{"Apple", "Banana", "Gauava"}
	//fmt.Println("Type of fruitList : %T", fruitList)

	fruitList = append(fruitList, "Mango", "Custard")
	//	fmt.Println(fruitList)

	//fruitList = append(fruitList[1:])
	//fmt.Println(fruitList)

	fruitList = append(fruitList[0:6])
	//fmt.Println(fruitList)

	//Approach 2
	highScores := make([]int, 4)

	highScores[0] = 353
	highScores[1] = 454
	highScores[2] = 234
	highScores[3] = 443
	//fmt.Println(highScores)

	highScores = append(highScores, 565, 333, 878)
	//fmt.Println(highScores)

	//fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	//fmt.Println(highScores)
	//fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove from slice based index.
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	//deleting using the append method
	fmt.Println("Deleting the swift in the slice")
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
