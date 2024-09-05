package main

import (
	"fmt"
)

var a int
var b int

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	fmt.Println("the number after swap is :", *a, *b)
}
func main() {
	fmt.Print("enter first number ")
	fmt.Scan(&a)
	fmt.Print("enter second number ")
	fmt.Scan(&b)
	fmt.Println("the number befre swap is ", a, b)
	swap(&a, &b)
}
