package main

import (
	"fmt"

	"strings"
	"unicode"
)

func main() {
	var input string
	fmt.Print("enter a string ")
	fmt.Scan(&input)
	v, c := count(input)
	fmt.Printf("vowels :%d,cons : %d\n ", v, c)

}
func count(s string) (vowels, constants int) {
	for _, char := range s {
		if unicode.IsLetter(char) {
			if strings.ContainsRune("aeiouAEIOU", char) {
				vowels++
			} else {
				constants++
			}
		}
	}
	return
}

//short kro isko
