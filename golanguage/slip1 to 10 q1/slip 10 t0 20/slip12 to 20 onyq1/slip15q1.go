package main

import "fmt"

func main() {
	var n int
	fmt.Print("enter the size ")
	fmt.Scan(&n)
	arr, even, odd := make([]int, n), []int{}, []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i]%2 == 0 {
			even = append(even, arr[i])
		} else {
			odd = append(odd, arr[i])
		}
	}
	fmt.Println("even", even, "odd", odd)
}
