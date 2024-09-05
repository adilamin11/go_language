package main

import "fmt"

// import "fmt"

// func main() {

// 	arr := []int{1, 2, 3, 4, 5}

// 	var pos, val int

// 	fmt.Println("Displaying or iterating into Slice:")

// 	for i, v := range arr {

// 		fmt.Println("index:", i, "value:", v)

// 	}

// 	fmt.Println()

// 	fmt.Println("1)Copying into the slice:")

// 	s2 := make([]int, len(arr))

// 	copy(s2, arr)

// 	fmt.Println("Original slice:", arr)

// 	fmt.Println("Copied array:", s2)

// 	fmt.Println()

// 	fmt.Println("2)Appending into slice")

// 	fmt.Print("Enter the value to append into slice:")

// 	fmt.Scan(&val)

// 	fmt.Println("After appending into slice")

// 	arr = append(arr, val)

// 	fmt.Println(arr)

// 	fmt.Println()

// 	fmt.Println("3)Deleting from the slice")

// 	fmt.Println("Before deleting from slice:", arr)

// 	fmt.Print("Enter position to delete::")

// 	fmt.Scan(&pos)

// 	new_arr := make([]int, (len(arr)))

// 	k := 0

// 	for i := 0; i < (len(arr)); i++ {

// 		if i != pos {

// 			new_arr[i] = arr[k]

// 			k++

// 		} else {

// 			k++

// 		}

// 	}

// 	fmt.Println("After deleting from slice:", new_arr)

// }
func main() {
	//create
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("orginal slice is ", numbers)
	//insert
	numbers = append(numbers[:5], 6)
	fmt.Println("slice after inserting ", numbers)
	//delete
	numbers = append(numbers[:2])
	fmt.Println("slice after delete ", numbers)
	//append
	numbers = append(numbers, 7)
	fmt.Println("slice after append ", numbers)
	//copy
	newnumber := make([]int, len(numbers))
	copy(newnumber, numbers)
	fmt.Println("copied slice ", newnumber)
}
