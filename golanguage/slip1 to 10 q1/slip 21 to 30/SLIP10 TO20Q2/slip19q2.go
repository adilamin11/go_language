package main

import (
	"fmt"
)

func main() {
	var str string
	var alphabet string
	var count int
	fmt.Println("enter the string ")
	fmt.Scanln(&str)
	fmt.Println("enter the alphabet ")
	fmt.Scanln(&alphabet)
	for i := 0; i < len(str); i++ {
		if str[i] == 1 {
			count++
		}
	}
	fmt.Println("count of string is ", alphabet, "in", str, "is ", count)
}
