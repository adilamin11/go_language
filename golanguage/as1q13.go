package main

import (
	"fmt"
)

func main() {
	// Get the number from the user
	fmt.Print("Enter a number in seconds: ")
	var num int
	fmt.Scanln(&num)

	// Convert the number to HH:MM:SS format
	hours := num / 3600
	minutes := (num % 3600) / 60
	seconds := num % 60

	// Display the number in HH:MM:SS format
	fmt.Printf("%d seconds is equivalent to %02d:%02d:%02d\n", num, hours, minutes, seconds)

	// Display the number in seconds
	fmt.Println("The number in seconds is:", num)
}
