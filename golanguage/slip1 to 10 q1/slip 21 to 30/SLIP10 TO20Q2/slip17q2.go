package main

import (
	"fmt"
)

func main() {
	num, squares, cubes := 123, 0, 0

	go func() {
		for ; num > 0; num /= 10 {
			digit := num % 10
			squares += digit * digit
		}
	}()

	num = 123 // Reset num to its original value

	go func() {
		for ; num > 0; num /= 10 {
			digit := num % 10
			cubes += digit * digit * digit
		}
	}()

	fmt.Printf("Sum of squares: %d\n", squares)
	fmt.Printf("Sum of cubes: %d\n", cubes)
}
