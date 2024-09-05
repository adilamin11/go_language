package main

import "fmt"

// func main() {
// 	var num int
// 	fmt.Print("enter the number ")
// 	fmt.Scan(&num)
// 	if isplind(num) {
// 		fmt.Printf("%d is a plindrome \n", num)

// 	} else {
// 		fmt.Printf("%d is not a plindrome \n", num)
// 	}
// }
// func isplind(n int) bool {
// 	reversed, original := 0, n
// 	for ; n > 0; n /= 10 {
// 		reversed = reversed*10 + n%10
// 	}
// 	return reversed == original
// }
func main() {
	var num int
	fmt.Println("enter the number ")
	fmt.Scan(&num)
	resilt := isplind(num)
	if resilt {
		fmt.Println("num is plindrome ", num)
	} else {
		fmt.Println("num is not plindroeme", num)
	}
}
func isplind(num int) bool {
	originalnum := 0
	reversed := 0
	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num = num / 10

	}
	return originalnum == reversed
}
