package main

import "fmt"

func main() {

	var i int

	var j int

	var c int

	fmt.Println("Printing prime numbers from 1 to 100:")

	for i = 1; i <= 100; i++ {

		c = 0

		for j = 1; j <= i; j++ {

			if i%j == 0 {

				c++

			}

		}

		if c == 2 {

			fmt.Print(i, "\n\n")

		}

	}

}
