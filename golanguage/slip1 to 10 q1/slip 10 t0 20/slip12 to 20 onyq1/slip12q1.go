package main

import (
	"fmt"
)

func main() {
	number := []int{11, 21, 41, 51, 61, 71, 41}
	min, max := number[0], number[0]
	for _, num := range number {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	fmt.Println("smallest number is %d\n", min)
	fmt.Println("largest number is %d\n ", max)
}
