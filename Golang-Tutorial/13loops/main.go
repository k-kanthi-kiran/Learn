package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	//fmt.Print(days)

	//syntax 1
	for d := 0; d < len(days); d++ {
		//	fmt.Println(days[d])
	}

	//syntax 2
	//for i := range days {
	//fmt.Println(days[i])
	//}

	//Syntax 3
	for _, day := range days {
		fmt.Printf("index is  and values is %v\n", day)
	}
}
