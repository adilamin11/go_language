package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter the number ")
	fmt.Scan(&num)
	fmt.Printf("factor of %d are: ", num)
	findfactor(num)
}
func findfactor(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
