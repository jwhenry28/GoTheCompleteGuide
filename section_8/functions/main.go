package main

import "fmt"

type transform func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	dNumbers := transformNumbers(&numbers, double)
	tNumbers := transformNumbers(&numbers, triple)

	transformed4 := transformNumbers(&numbers, buildTransformer(4))
	transformed5 := transformNumbers(&numbers, buildTransformer(5))
	transformed6 := transformNumbers(&numbers, func(number int) int { return number * 6 })

	fmt.Println(dNumbers)
	fmt.Println(tNumbers)
	fmt.Println(transformed4)
	fmt.Println(transformed5)
	fmt.Println(transformed6)
}

func buildTransformer(factor int) transform {
	return func(number int) int {
		return number * factor
	}
}

func transformNumbers(numbers *[]int, transformer transform) []int {
	tNumbers := []int{}
	
	for _, num := range *numbers {
		tNumbers = append(tNumbers, transformer(num))
	}

	return tNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}