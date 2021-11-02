package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()
	fmt.Println("The present time is :", presentTime)
	fmt.Println("Formatting the time as per our requirement:", presentTime.Format("02/01/2006 15:04:05 Monday"))

	createdDate := time.Date(2021, time.October, 26, 07, 43, 00, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
