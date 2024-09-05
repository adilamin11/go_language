package main

import "fmt"

func main() {
	// Declare two strings
	str1 := "Hello"
	str2 := "World"

	// Declare two pointers to the strings
	str1Ptr := &str1
	str2Ptr := &str2

	// Concatenate the two strings using pointers
	str3 := fmt.Sprintf("%s%s", *str1Ptr, *str2Ptr)

	// Print the concatenated string
	fmt.Println(str3)
}
