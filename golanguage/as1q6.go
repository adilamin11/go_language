package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of terms: ")
	fmt.Scanln(&n)

	a := 0
	b := 1
	fmt.Println("The Fibonacci series is:")
	for i := 0; i < n; i++ {
		fmt.Println(a)
		temp := a
		a = b
		b = temp + b
	}
}
