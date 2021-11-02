package main

import "fmt"

func main() {

	//creating the map with key and value as string
	languages := make(map[string]string)

	//loading the values in the map
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)

	fmt.Println("JS short for :", languages["JS"])

	//Deleting the content from map

	delete(languages, "RB")
	fmt.Println("Content Ruby after deleting :", languages)

	//looping through the map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
