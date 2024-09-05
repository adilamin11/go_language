// 4.WAP in golang to print sum of digits of given number(use recursive function).
package main

import "fmt"

var sum int = 0

func sumofdigit(num int) int {
	if num > 0 {
		sum += (num % 10)
		sumofdigit(num / 10)
	}
	return sum
}
func main() {
	var num int = 0
	var result int = 0
	fmt.Println("enter the number ")
	fmt.Scan(&num)
	result = sumofdigit(num)
	fmt.Println("sum of digit is ", result)
}
