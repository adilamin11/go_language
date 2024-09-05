package main

import "fmt"

func main() {
	// Read a number and store it in a variable.
	var num int
	fmt.Println("Enter a number to print table:")
	fmt.Scanln(&num)

	// Print the multiplication table of the given number.
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
}
