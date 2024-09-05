package main

import "fmt"

func main() {

	var rows int

	var temp int = 1

	fmt.Print("Enter the no_of_rows:")

	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < rows-i; j++ {
			fmt.Println(" ")
		}
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
		}
		fmt.Printf(" %d", temp)
	}

	fmt.Println()

}
