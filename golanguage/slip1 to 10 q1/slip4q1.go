package main

import "fmt"

func main() {
	var str1 string
	var str2 string
	fmt.Println("enter first string")
	fmt.Scan(&str1)
	fmt.Println("enter second string")
	fmt.Scan(&str2)
	if str1 == str2 {
		fmt.Println("string are same or equal")

	} else {
		fmt.Println("string are not equal")
	}
}
