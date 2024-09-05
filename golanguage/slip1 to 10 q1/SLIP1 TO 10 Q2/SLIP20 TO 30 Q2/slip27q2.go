package main

import "fmt"

type Books struct {
	b_id int

	title string

	author string

	price int
}

func main() {

	var n, i int

	fmt.Print("Enter how many records you want to store:")

	fmt.Scan(&n)

	obj := make([]Books, n)

	for i = 0; i < n; i++ {

		fmt.Print("Enter Book_id:")

		fmt.Scan(&obj[i].b_id)

		fmt.Print("Enter Book title:")

		fmt.Scan(&obj[i].title)

		fmt.Print("Enter Book author:")

		fmt.Scan(&obj[i].author)

		fmt.Print("Enter Book price:")

		fmt.Scan(&obj[i].price)

		fmt.Println()

	}

	fmt.Println("<<<------Displaying Book Details------>>>")

	for i = 0; i < n; i++ {

		printbook(obj[i])

	}

}

func printbook(obj Books) {

	fmt.Println("BOOK ID is:", obj.b_id)

	fmt.Println("BOOK TITLE is:", obj.title)

	fmt.Println("BOOK AUTHOR is:", obj.author)

	fmt.Println("BOOK PRICE is:", obj.price)

	fmt.Println()

}
