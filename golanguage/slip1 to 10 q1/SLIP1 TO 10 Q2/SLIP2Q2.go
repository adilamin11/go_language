package main

import "fmt"

func main() {

	var rows, columns, i, j int

	var mat1 [10][10]int

	var mat2 [10][10]int

	var multimat [10][10]int

	fmt.Print("Enter no.of.Rows=")

	fmt.Scan(&rows)

	fmt.Print("Enter no.of.Columns=")

	fmt.Scan(&columns)

	fmt.Println("Enter the First Matrix elements=")

	for i = 0; i < rows; i++ {

		for j = 0; j < columns; j++ {

			fmt.Scan(&mat1[i][j])

		}

	}

	fmt.Print("Enter the Second Matrix elements=")

	for i = 0; i < rows; i++ {

		for j = 0; j < columns; j++ {

			fmt.Scan(&mat2[i][j])

		}

	}

	fmt.Println("the first matrix is :")

	for i = 0; i < rows; i++ {

		for j = 0; j < columns; j++ {

			fmt.Print(mat1[i][j], "\t")

		}

		fmt.Println()

	}

	fmt.Println("the second matrix is :")

	for i = 0; i < rows; i++ {

		for j = 0; j < columns; j++ {

			fmt.Print(mat2[i][j], "\t")

		}

		fmt.Println()

	}

	for i = 0; i < rows; i++ {

		for j = 0; j < columns; j++ {

			multimat[i][j] = mat1[i][j] * mat2[i][j]

		}

	}

	fmt.Println("The Result of Matrix Multiplication = ")

	for i = 0; i < rows; i++ {

		for j = 0; j < columns; j++ {

			fmt.Print(multimat[i][j], "\t")

		}

		fmt.Println()

	}
}
