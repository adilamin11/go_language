package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter the number")
	fmt.Scanf("%d", &num)
	if num%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
}
