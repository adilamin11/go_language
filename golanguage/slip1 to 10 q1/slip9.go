// 12.WAP to accept two numbers and perform mathematical operation choice from user(use switch case).
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var choice int
	fmt.Println("enter the ist number ")
	fmt.Scan(&num1)
	fmt.Println("enter the 2nd number ")
	fmt.Scan(&num2)
	fmt.Println("1.addition \n", "2.for subtraction \n", "3.for division \n", "4.for multi", "enter the choice \n :")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("additon of two number ", num1+num2)
	case 2:
		fmt.Println("subtraction of two ", num1-num2)
	case 3:
		fmt.Println("division of two num ", num1/num2)
	case 4:
		fmt.Println("multiplication of number ", num1*num2)
	default:
		fmt.Println("its a error")
	}

}
