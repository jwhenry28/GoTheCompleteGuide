package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := retrieveValue("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := retrieveValue("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := retrieveValue("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	printValues(ebt, profit, ratio)
	writeValues(ebt, profit, ratio)
}

func writeValues(ebt, profit, ratio float64) error {
	profitText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	err := os.WriteFile("results.txt", []byte(profitText), 0644)
	return err
}

func retrieveValue(text string) (float64, error) {
	var input float64

	fmt.Print(text)
	_, err := fmt.Scan(&input)
	if err != nil {
		return input, err
	} else if input <= 0 {
		return input, errors.New("value must be positive")
	}

	return input, nil
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func printValues(ebt, profit, ratio float64) {
	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}