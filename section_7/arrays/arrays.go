package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0} 
	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[2])

	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	featuredPrices := prices[0:1]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[0:]
	fmt.Println(featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)

	featuredPrices[0] = 100.0
	fmt.Println(prices)

	featuredPrices = prices[0:2]
	fmt.Println(len(prices), cap(prices))
	fmt.Println(len(featuredPrices), cap(featuredPrices))

	dynamicPrices := []float64{10.99, 8.99}
	updatedPrices := append(dynamicPrices, 5.99)
	dynamicPrices[0] = 0.99

	fmt.Println(dynamicPrices)
	fmt.Println(updatedPrices)

	discountPrices := []float64{101.99, 90.99, 20.59}
	updatedPrices = append(dynamicPrices, discountPrices...)
	fmt.Println(updatedPrices)
}