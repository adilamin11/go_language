package main

import "fmt"

var num int
var rem int
var rev int = 0

func isplind(num int) {
	var temp int = num
	for num > 0 {
		rem = num % 10
		rev = rev*10 + rem
		num = num / 10
	}
	if temp == rev {
		fmt.Println("number is plindrome")
	} else {
		fmt.Println("number is not plindrome")
	}

}
func main() {
	fmt.Println("enter the number")
	fmt.Scanf("%d", &num)
	isplind(num)

}
