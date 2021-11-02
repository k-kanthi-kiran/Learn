package main

import ("fmt"
	    "bufio"
	    "os"
	    )

func main(){

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the rating");


	input, err := reader.ReadString('\n');
	fmt.Println("Thanks for rating :",input);
	fmt.Printf("Type of reader : %T \n",reader);
	fmt.Printf("Type of input : %T",input);

}