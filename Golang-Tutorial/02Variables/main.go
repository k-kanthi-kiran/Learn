package main
import "fmt"

const LoginToken string = "aslfjlkjasf0" //public 

func main(){

	var username string = "kanthi kiran"
	fmt.Println(username)
	fmt.Printf("Variable type is : %T \n", username)

	var isFlag bool = true
	fmt.Println(isFlag)
	fmt.Printf("Variable type is : %T \n", isFlag)

	var smallFloat float32 = 255.4555555
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is : %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is : %T \n", anotherVariable)

	//Implicit Type
	var website = "learncodeline.in"
	fmt.Println(website)
	
	//no var style
	numberOfUsers := 30000.343
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is : %T \n", LoginToken)



}