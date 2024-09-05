package main

import (
	"fmt"
)

var num int

func main() {
	fmt.Println("enter the number of fib series ")
	fmt.Scan(&num)
	a := 0
	b := 1
	c := b
	fmt.Println("series is ", a)
	for i := 0; i < num; i++ {
		c = b
		b = a + b
		fmt.Println("\t")
		if b >= num {
			fmt.Println()
		}
		a = c
		fmt.Println(b)

	}
}
