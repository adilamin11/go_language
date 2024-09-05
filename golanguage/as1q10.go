package main

import "fmt"

func main() {
	// Initialize the array
	arr := []int{4, 7, 9, 8, 6, 3, 2, 5, 4}

	// Make the array upper triangular
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			arr[j] = 0
		}
	}

	// Print the upper triangular array
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if j >= i {
				fmt.Print(arr[j], " ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
