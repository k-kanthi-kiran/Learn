package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   string
}

func main() {

	//Approach 1
	rob := User{"rob", "rob@xyz.com", "21"}
	//less information is printed on the console.
	fmt.Printf("%v\n", rob)
	//more information is printed on the console
	fmt.Printf("%+v\n", rob)

	//Approach 2
	var sam = new(User)
	sam.Email = "Sam@xyz.com"
	sam.Name = "Sam"
	sam.Age = "33"
	//Reference will be printed in this scenario
	fmt.Printf("%+v\n", sam)

	//Approach 3
	//We are creating the structure itself as reference type.
	var tobby = &User{"tobby", "tobby@xyz.com", "43"}

	fmt.Printf("%+v\n", tobby)

}
