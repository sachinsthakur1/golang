package main

import "fmt"

func addition(b ...int) int {
	sum := 0
	for _, val := range b {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("Sum of 1, 2, 3:", addition(1, 2, 3))       // Output: 6
	fmt.Println("Sum of 4, 5, 6, 7:", addition(4, 5, 6, 7)) // Output: 22
	fmt.Println("Sum of 8, 9:", addition(8, 9))             // Output: 17
}
