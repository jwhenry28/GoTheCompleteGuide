package main

import "fmt"

type logString string

func (text logString) log() {
	fmt.Println(text)
}

func main() {
	var name logString = "Joseph"
	name.log()
}