package main

import "fmt"

func main() {

	var str1 string

	var str2 string

	var p1 *string = &str1

	var p2 *string = &str2

	fmt.Print("Enter first string:")

	fmt.Scan(&str1)

	fmt.Print("Enter second string:")

	fmt.Scan(&str2)

	res := *p1 + *p2

	fmt.Println("The concatenation of both the string using pointer:", res)

}
