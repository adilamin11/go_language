package main

import "fmt"

func main() {
	var year int
	fmt.Println("enter the year")
	fmt.Scanf("%d", &year)
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println(year, "year is a leap")

	} else {
		fmt.Println(year, "year not leap")
	}
}
