package main

import "fmt"

func main() {
	var n, evencnt, oddcnt int
	fmt.Print("enter the size of array ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Print("enter the element of array ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i]%2 == 0 {
			evencnt++
		} else {
			oddcnt++
		}
	}
	fmt.Println("even number count ", evencnt)
	fmt.Println("odd number count ", oddcnt)
}
