package main

import "fmt"

// calcultae returns the sum and product of two integers using named return values
func calcultae(a int, b int) (sum int, product int) {
	// Calculate sum
	sum = a + b
	// Calculate product
	product = a * b
	return // Named return values are returned automatically
}

func main() {
	// Initialize an array with numbers 1 to 10
	poolOfNumbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Pool of Numbers:", poolOfNumbers)
	// Print sum and product of 2 and 5
	fmt.Println(calcultae(2, 5))
}
