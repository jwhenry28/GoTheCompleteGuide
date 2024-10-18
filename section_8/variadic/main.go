package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15)
	sum2 := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(sum2)
}

func sumup(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}