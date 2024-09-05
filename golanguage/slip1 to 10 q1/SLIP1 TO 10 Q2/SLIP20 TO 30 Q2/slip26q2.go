package main

import (
	"fmt"

	"sort"
)

type student struct {
	rollNo int

	name string

	perc float64
}

type byPercentage []student

func (s byPercentage) Len() int {

	return len(s)

}

func (s byPercentage) Swap(i, j int) {

	s[i], s[j] = s[j], s[i]

}

func (s byPercentage) Less(i, j int) bool {

	return s[i].perc > s[j].perc

}

func main() {

	var n int

	fmt.Print("Enter number of students: ")

	fmt.Scanf("%d", &n)

	students := make([]student, n)

	for i := 0; i < n; i++ {

		fmt.Printf("Enter details of student %d\n", i+1)

		fmt.Print("Roll number: ")

		fmt.Scan(&students[i].rollNo)

		fmt.Print("Name: ")

		fmt.Scan(&students[i].name)

		fmt.Print("Percentage: ")

		fmt.Scan(&students[i].perc)

	}

	sort.Sort(byPercentage(students))

	fmt.Println("Student information in descending order of percentage:")

	for _, student := range students {

		fmt.Printf("Roll number: %d, Name: %s, Percentage: %.2f\n", student.rollNo, student.name, student.perc)

	}

}
