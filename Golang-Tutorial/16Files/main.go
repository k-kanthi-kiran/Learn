package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	file, error := os.Create("kanthi.txt")
	if error != nil {
		panic(error)
	}
	content := "First file in GoLang"

	len, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Lenght of the file:", len)
	file.Close()
	readFile("kanthi.txt")
}

func readFile(fileName string) {
	bytes, error := ioutil.ReadFile(fileName)
	if error != nil {
		panic(error)
	}
	fmt.Println("Content is :\n", string(bytes))
}
