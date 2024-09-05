package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{11, 31, 21, 8, 77, 51, 61}
	sort.Ints(numbers)
	fmt.Println("sorted array in ascending order ", numbers)
}
