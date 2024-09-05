package main

import "fmt"

var max int = 0

type Employee struct {
	eno int

	name string

	salary int
}

func main() {

	var n, i, j int

	fmt.Print("Enter how many Employee records you want to store:")

	fmt.Scan(&n)

	obj := make([]Employee, n)

	for i = 0; i < n; i++ {

		fmt.Print("Enter Employee no:")

		fmt.Scan(&obj[i].eno)

		fmt.Print("Enter Employee name:")

		fmt.Scan(&obj[i].name)

		fmt.Print("Enter Employee salary:")

		fmt.Scan(&obj[i].salary)

		fmt.Println()

	}

	fmt.Println("<<<-Displaying Employee Details Whose Salary Is Maximun->>>")

	max = obj[0].salary

	for i = 0; i < n; i++ {

		if obj[i].salary > max {

			max = obj[i].salary

		}

	}

	for j = 0; j < n; j++ {

		if max == obj[j].salary {

			display(obj[j])

		}

	}

}

func display(obj Employee) {

	fmt.Println("EMPLOYEE NO  is:", obj.eno)

	fmt.Println("EMPLOYEE NAME is:", obj.name)

	fmt.Println("EMPLOYEE SALARY is:", obj.salary)

}
