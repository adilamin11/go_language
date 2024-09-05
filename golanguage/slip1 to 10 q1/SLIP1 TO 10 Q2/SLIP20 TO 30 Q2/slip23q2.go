package main

import "fmt"

var total int

var average int

type Student struct {
	rno int

	name string

	mar1 int

	mar2 int

	mar3 int
}

func main() {

	var n, i int

	fmt.Print("Enter how many student records you want to store:")

	fmt.Scan(&n)

	obj := make([]Student, n)

	for i = 0; i < n; i++ {

		fmt.Print("Enter Student rollno:")

		fmt.Scan(&obj[i].rno)

		fmt.Print("Enter Student name:")

		fmt.Scan(&obj[i].name)

		fmt.Print("Enter marks of subject1:")

		fmt.Scan(&obj[i].mar1)

		fmt.Print("Enter marks of subject2:")

		fmt.Scan(&obj[i].mar2)

		fmt.Print("Enter marks of subject3:")

		fmt.Scan(&obj[i].mar3)

		fmt.Println()

	}

	fmt.Println("<<<------Displaying Student Details------>>>")

	for i = 0; i < n; i++ {

		display(obj[i])

	}

}

func display(obj Student) {

	total = obj.mar1 + obj.mar2 + obj.mar3

	average = total / 3

	fmt.Println("STUDENT ROLLNO:", obj.rno)

	fmt.Println("STUDENT NAME:", obj.name)

	fmt.Println("TOTAL MARKS:", total)

	fmt.Println("AVERAGE OF MARKS:", average)

	fmt.Println()

}
