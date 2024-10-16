package main

import "fmt"

func main() {
	age := 32 // regulat variable

	var agePointer *int
	agePointer = &age

	fmt.Println("Age Pointer:", agePointer)
	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Age:", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}