package main

import "fmt"

func main() {
	var name string
	var rollno int
	var division string
	var college string
	fmt.Print("enter student name ")
	fmt.Scan(&name)

	fmt.Print("enter roll no= ")
	fmt.Scan(&rollno)

	fmt.Print("enter dividion = ")
	fmt.Scan(&division)

	fmt.Print("enter colloge name =")
	fmt.Scan(&college)
	//output
	fmt.Println("\n students information")
	fmt.Printf(" name is : %s\n ", name)
	fmt.Printf(" rollno is : %d\n ", rollno)
	fmt.Printf(" division is: %s\n ", division)
	fmt.Printf(" colloge is %s\n ", college)
}
