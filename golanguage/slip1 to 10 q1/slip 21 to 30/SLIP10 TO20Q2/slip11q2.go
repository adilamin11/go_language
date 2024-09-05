package main

import "fmt"

type student struct {
	name string

	branch string

	rollno int
}

func (s *student) show(sbranch string) {

	(*s).branch = sbranch

}

func main() {

	res := student{

		name: "adil",

		branch: "BCA",
	}

	fmt.Println("Student's name: ", res.name)

	fmt.Println("Branch Name(Before): ", res.branch)

	// Creating a pointer

	p := &res

	// Calling the show method

	p.show("MCA")

	fmt.Println("Student's name: ", res.name)

	fmt.Println("Branch Name(After): ", res.branch)

}
