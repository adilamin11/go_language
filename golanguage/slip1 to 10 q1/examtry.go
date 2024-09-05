package main

import (
	"fmt"
)

// func main() {
// 	var num int
// 	fmt.Println("enter the number")
// 	fmt.Scan(&num)
// 	if num%2 == 0 {
// 		fmt.Println("number is even ")

//		} else {
//			fmt.Println("number is odd")
//		}
//	}
//
//	func main() {
//		var year int
//		fmt.Println("enter the year ")
//		fmt.Scan(&year)
//		if year%4 == 0 && (year%100 != 0 && year%400 == 0) {
//			fmt.Println("year is leap")
//		} else {
//			fmt.Println("year is not leap ")
//		}
//	}
// func main() {
// 	go func() {
// 		fmt.Println("anon go ")
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("done")
// 	}()
// 	time.Sleep(3 * time.Second)
// }

//	func main() {
//		numbers := []int{22, 12, 13, 2, 3, 4, 5, 1}
//		sort.Ints(numbers)
//		fmt.Println("sorted arrays is ", numbers)
//	}
// var num int
// var rem int
// var rev int = 0

// func plind(num int) {
// 	var temp int = num
// 	for num > 0 {
// 		rem = num % 10
// 		rev = rev*10 + rem
// 		num = num / 10

// 	}
// 	if temp == rev {
// 		fmt.Println("num is plindrome")
// 	} else {
// 		fmt.Println("num is not plindrome")
// 	}
// }

// func main() {
// 	var num int
// 	fmt.Println("enter number ")
// 	fmt.Scanf("%d", &num)
// 	plind(num)

// }
// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 		delay(100 * time.Millisecond)

//		}
//	}
//
//	func delay(d time.Duration) {
//		time.Sleep(d)
//	}
type student struct {
	name   string
	rollno int
	marks  int
}

func (s *student)show(smarks int )  {
	(*s).marks=smarks
}


func main() {
	res := student{
		name:  "adil",
		marks: 90,
	}
	fmt.Println("stude nt name ", res.name)
	fmt.Println("student marks ", res.marks)
	p := &res
	p.show(80)
	fmt.Println("student name ", res.name)
	fmt.Println("marks beforw", res.marks)
}
