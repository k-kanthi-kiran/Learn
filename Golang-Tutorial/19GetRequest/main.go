package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	//performGetRequest()
	//performPostJSONRequest()
	performPostFormJsonRequest()

}

func performGetRequest() {
	const url = "http://localhost:8000/get"

	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	output, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Output is :", string(output))
}

func performPostJSONRequest() {
	const url = "http://localhost:8000/post"

	//Json payload
	requestBody := strings.NewReader(`
		{
				"Animal" : "Tiger",
				"Currency" : 20.50,
				"country" : "India"
		}
	`)
	response, error := http.Post(url, "application/json", requestBody)
	defer response.Body.Close()
	if error != nil {
		panic(error)
	}
	output, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Message received from server :", string(output))
}

func performPostFormJsonRequest() {

	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Kanthi")
	data.Add("lastname", "Kiran")
	data.Add("email", "kk@gmail.com")

	response, error := http.PostForm(myurl, data)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	output, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Output is :", string(output))
}
