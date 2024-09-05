package main

import "fmt"

func main() {
	// Create an array to store the input numbers
	var inputArray [10]int

	// Get the input numbers from the user
	fmt.Println("Enter 10 numbers:")
	for i := 0; i < 10; i++ {
		fmt.Scanln(&inputArray[i])
	}

	// Create two arrays to store the even and odd numbers
	var evenArray [10]int
	var oddArray [10]int

	for i := 0; i < 10; i++ {
		if inputArray[i]%2 == 0 {
			evenArray[i] = inputArray[i]
		} else {
			oddArray[i] = inputArray[i]
		}
	}

	fmt.Println("Even numbers:")
	for i := 0; i < 10; i++ {
		fmt.Println(evenArray[i])
	}

	fmt.Println("Odd numbers:")
	for i := 0; i < 10; i++ {
		fmt.Println(oddArray[i])
	}
}
