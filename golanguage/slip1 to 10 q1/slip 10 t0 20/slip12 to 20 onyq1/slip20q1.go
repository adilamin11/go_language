package main

import (
	"fmt"
	"math"
)

func main() {
	var base, exp float64
	fmt.Print("enter the base ")
	fmt.Scan(&base)
	fmt.Print("enter the exponent ")
	fmt.Scan(&exp)
	result := math.Pow(base, exp)
	fmt.Printf("%.2f % .2f = %.2f\n", base, exp, result)

}
