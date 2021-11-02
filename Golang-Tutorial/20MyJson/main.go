package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string `json:"coursename"`
	Price     int
	Platform  string   `json:"website"`
	Passwords string   `json:"-"`
	Tags      []string `json:"tags,omitempty"`
}

func main() {
	//jsonMarshal()
	jsonUnMarshal()
}

func jsonMarshal() {

	lcoCourse := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnLine.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnLine.in", "bcd123", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 199, "LearnCodeOnLine.in", "kiran123", nil},
	}

	//The below line of Marshal just marshall's the data where the output of the json is not formatted.
	//	finalJson, error := json.Marshal(lcoCourse)

	//In order to format the json
	finalJson, error := json.MarshalIndent(lcoCourse, "", "\t")
	if error != nil {
		panic(error)
	}

	fmt.Printf("%s\n", finalJson)
}

func jsonUnMarshal() {
	jsonWebData := []byte(`{
	"coursename": "ReactJS Bootcamp",
	"Price": 299,
	"website": "LearnCodeOnLine.in", 
	"tags": ["web-dev",	"js"]
}`)

	var lcoCourse course

	checkValid := json.Valid(jsonWebData)
	if checkValid {

		json.Unmarshal(jsonWebData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Printf("Invalid JSON")
	}
} //end of Json
