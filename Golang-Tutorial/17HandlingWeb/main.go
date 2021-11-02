package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {

	response, error := http.Get(URL)
	if error != nil {
		panic(error)
	}

	content, error := ioutil.ReadAll(response.Body)

	fmt.Println("Output is :", string(content))

	defer response.Body.Close()

}
