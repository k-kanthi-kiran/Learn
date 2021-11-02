package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    string
	Status string
}

func (u User) GetStatus() {
	fmt.Println("The status of user is :", u.Status)
}

func (u User) setEmail() {
	u.Email = "kk@kk.com"
	fmt.Println("Maid Id is updated to :", u.Email)
	fmt.Printf("%+v\n", u)
}

func main() {

	//Approach 1
	rob := User{"rob", "rob@xyz.com", "21", "Active"}
	//more information is printed on the console
	fmt.Printf("%+v\n", rob)
	rob.GetStatus()
	fmt.Printf("%+v\n", rob)
	rob.setEmail()
	fmt.Printf("%+v\n", rob)
}
