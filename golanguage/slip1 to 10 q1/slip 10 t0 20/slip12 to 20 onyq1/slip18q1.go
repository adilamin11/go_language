package main

import "fmt"

var a int
var b int

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	fmt.Println("afteer swap ", &a, &b)
}
func main() {
	var num1 int
	var num2 int
	fmt.Println("enter the num1 ")
	fmt.Scan(&num1)
	fmt.Println("enter the num2 ")
	fmt.Scan(&num2)
	fmt.Println("number befoer swap ", a, b)
	swap(&a, &b)
}
