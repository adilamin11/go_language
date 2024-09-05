package main

import "fmt"

func main() {
	var char rune
	fmt.Println("enter the character")
	fmt.Scanf("%c", &char)
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		fmt.Println("its a vowel")
	default:
		fmt.Println("its a  consonant")

	}
}
