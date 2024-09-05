package main

import "fmt"

var a int = 100
var b int = 200

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	fmt.Println("the number after swap is", *a, *b)

}
func main() {
	// fmt.Println("enter the first number ")
	// fmt.Scan(&a)
	// fmt.Println("enter the second number")
	// fmt.Scan(&b)

	fmt.Println("the number before swap is ", a, b)
	swap(&a, &b)
}
