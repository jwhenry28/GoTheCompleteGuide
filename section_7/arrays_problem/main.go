package main

import (
	"fmt"
)

type Product struct {
	title string
	id int
	price float64
}

func main() {
	// 1.
	hobbies := [3]string{"reading", "lifting", "coding"}
	fmt.Println("Hobbies:", hobbies)

	// 2. 
	fmt.Println("First element:", hobbies[0])
	fmt.Println("Second and third elements:", hobbies[1:])

	// 3.
	hobbiesSlice1 := hobbies[0:2]
	hobbiesSlice2 := hobbies[:2]
	fmt.Println("Slice 1:", hobbiesSlice1)
	fmt.Println("Slice 2:", hobbiesSlice2)

	// 4.
	hobbiesSlice3 := hobbiesSlice2[1:3]
	fmt.Println("Slice 3:", hobbiesSlice3)

	// 5. 
	goals := []string{"learn Go", "become a coder"}

	// 6. 
	goals[1] = "become a Go programmer"
	goals = append(goals, "build Chariot")
	fmt.Println("Goals:", goals)

	// 7. 
	products := []Product{
		{"book", 1, 9.99}, 
		{"bicycle", 2, 119.99},
	}
	fmt.Println("Products:", products)
	products = append(products, Product{"laptop", 3, 999.99})
	fmt.Println("Products:", products)
}