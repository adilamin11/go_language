package main

import (
	"fmt"

	"math"
)

type Shape interface {
	Area() float64

	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {

	return math.Pi * c.Radius * c.Radius

}

func (c Circle) Perimeter() float64 {

	return 2 * math.Pi * c.Radius

}

type Rectangle struct {
	Length float64

	Width float64
}

func (r Rectangle) Area() float64 {

	return r.Length * r.Width

}

func (r Rectangle) Perimeter() float64 {

	return 2 * (r.Length + r.Width)

}

func main() {

	c := Circle{

		Radius: 8,
	}

	fmt.Println("Area of the Circle:", c.Area())

	fmt.Println("Perimeter of the Circle:", c.Perimeter())

	r := Rectangle{

		Length: 12,

		Width: 8,
	}

	fmt.Println("Area of the rectangle:", r.Area())

	fmt.Println("Perimeter of the rectangle:", r.Perimeter())

}
