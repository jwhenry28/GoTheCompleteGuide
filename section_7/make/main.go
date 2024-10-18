package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}


func main() {
	userNames := make([]string, 2, 4)

	userNames[0] = "Alice"
	userNames[1] = "Bob"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "John")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 2)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}