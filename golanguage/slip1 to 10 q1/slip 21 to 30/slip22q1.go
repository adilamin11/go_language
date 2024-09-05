package main

import "fmt"

func isplind(num int) bool {
	orginal_num := num
	reversed := 0
	for num > 0 {
		rem := num % 10
		reversed = reversed*10 + rem
		num = num / 10

	}
	return orginal_num == reversed

}

func main() {
	var num int
	fmt.Print("enter the number ")
	fmt.Scan(&num)
	result := isplind(num)
	if result {
		fmt.Printf("%d is a plindrome \n", num)

	} else {
		fmt.Printf("%d is not a plindrome \n", num)
	}
}
