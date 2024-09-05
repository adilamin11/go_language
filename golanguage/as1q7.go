package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare variables
	var num1, num2, num3 float64

	// Read input from user
	fmt.Println("Enter three numbers:")
	fmt.Scanln(&num1, &num2, &num3)

	// Find the maximum number
	max := math.Max(math.Max(num1, num2), num3)

	// Print the result
	fmt.Println("The maximum number is:", max)
}
