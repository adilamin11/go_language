package main

import (
	"fmt"
)

func main() {
	var num, squares, cubes, digit int
	fmt.Scan(&num)

	for ; num > 0; num /= 10 {
		digit = num % 10
		squares += digit * digit
		cubes += digit * digit * digit
	}

	fmt.Printf("Sum of squares = %d\n", squares)
	fmt.Printf("Sum of cubes = %d\n", cubes)
	fmt.Printf("Final sum of squares and cubes = %d\n", squares+cubes)
}
