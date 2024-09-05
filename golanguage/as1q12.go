package main

import (
	"fmt"
)

func main() {
	// Get the two numbers from the user.
	fmt.Println("Enter two numbers:")
	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	// Get the arithmetic operation choice from the user.
	fmt.Println("Enter an arithmetic operation (+, -, *, /):")
	var operation string
	fmt.Scanln(&operation)

	// Perform the arithmetic operation based on the user's choice.
	var result int
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Invalid operation.")
		return
	}

	// Print the answer.
	fmt.Println("The answer is:", result)
}
