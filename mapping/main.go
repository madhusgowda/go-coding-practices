package main

import "fmt"

var data = map[string]string{
	"foo":   "bar",
	"Hello": "world",
}

func main() {
	//mapping
	if val, isExists := data["Hello"]; isExists {
		fmt.Println("Found the value:", val)
	} else {
		fmt.Println("Uh oh couldn't find the value :'( for data")
	}
}
