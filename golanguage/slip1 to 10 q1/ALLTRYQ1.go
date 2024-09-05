package main

import (
	"fmt"
	"time"
)

//func main() {
// SLIP1Q1
// var num int
// fmt.Println("enter the number ")
// fmt.Scan(&num)
//
//	if num%2 == 0 {
//		fmt.Println("number is even ")
//	} else {
//
//		fmt.Println("number is odd")
//	}
//
// //SLIP2q1
// var year int
// fmt.Println("enter the year ")
// fmt.Scan(&year)
//
//	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
//		fmt.Println("years is leap")
//	} else {
//
//		fmt.Println("year is not leap")
//	}
//
// SLIP3Q1
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
//SLIP4Q1
// var str1 string
// var str2 string
// fmt.Println("enter the string ")
// fmt.Scan(&str1)
// fmt.Println("enter the string2 ")
// fmt.Scan(&str2)
// if str1 == str2 {
// 	fmt.Println("string is equal ")
// } else {
// 	fmt.Println("not equal ")
// }
//SLIP5Q1
// var num int
// fmt.Println("enter the number ")
// fmt.Scan(&num)
// for i := 0; i <= 10; i++ {
// 	fmt.Println(num, " x", i, "=", num*i)
// }
//SLIP6Q1
// var char rune
// fmt.Println("enter the char ")
// fmt.Scanf("%c", &char)
// switch char {
// case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// 	fmt.Println("is a vowel")
// default:
// 	fmt.Println("is not const")

// }
//SLIP7Q1

// var a int
// var b int

//	func swap(a *int, b *int) {
//		var temp int
//		temp = *a
//		*a = *b
//		*b = temp
//		fmt.Println("after swap ", *a, *b)
//	}
//
//	func main() {
//		fmt.Println("enter num1")
//		fmt.Scan(&a)
//		fmt.Println("enter 2num")
//		fmt.Scan(&b)
//		fmt.Println("before swap ", a, b)
//		swap(&a, &b)
//	}
//
// SLIP8Q1
//
//	func main() {
//		arr := []int{1, 2, 3, 4, 5, 6, 7}
//		fmt.Println("original array ", arr)
//		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
//			arr[i], arr[j] = arr[j], arr[i]
//		}
//		fmt.Println("reversed array ", arr)
//	}
//
// SLIP9Q1
// var num1 int
// var num2 int
// var choice int

// func main() {
// 	fmt.Println("enter the num 1")
// 	fmt.Scan(&num1)
// 	fmt.Println("enter the num 2")
// 	fmt.Scan(&num2)
// 	fmt.Println("1.addition \n", "2.subract \n ", "3.mul \n ", "4.div", "enter the choice :")
// 	fmt.Scan(&choice)
// 	switch choice {
// 	case 1:
// 		fmt.Println("addition of number ", num1+num2)
// 	case 2:
// 		fmt.Println("sub is ", num1-num2)
// 	case 3:
// 		fmt.Println("mul is ", num1*num2)
// 	case 4:
// 		fmt.Println("div is ", num1/num2)
// 	default:
// 		fmt.Println("errror hao bahii")

// 	}
// }
//SLIP10Q1
// var sum int = 0

// func sumof(num int) int {
// 	if num > 0 {
// 		sum += (num % 10)
// 		sumof(num / 10)
// 	}
// 	return sum
// }
// func main() {
// 	var num int = 0
// 	var result int = 0
// 	fmt.Println("enter the number ")
// 	fmt.Scan(&num)
// 	result = sumof(num)
// 	fmt.Println("sum of digit is ", result)
// }
//SLIP11Q1
// var num int

//	func main() {
//		fmt.Println("enter the number of sreies")
//		fmt.Scan(&num)
//		a := 0
//		b := 1
//		c := b
//		fmt.Println("series is ", a)
//		for i := 0; i < num; i++ {
//			c = b
//			b = a + b
//			fmt.Println("\t")
//			if b >= num {
//				fmt.Println()
//			}
//			a = c
//			fmt.Println(b)
//		}
//	}
//
// SLIP12Q1
//
//	func main() {
//		number := []int{11, 22, 33, 44, 5, 55, 66}
//		min, max := number[0], number[0]
//		for _, num := range number {
//			if num < min {
//				min = num
//			}
//			if num > max {
//				max = num
//			}
//		}
//		fmt.Println("smallest no is %d\n ", min)
//		fmt.Println("largest num is %d \n ", max)
//	}
//
// SLIP13Q1
func main() {
	go func() {
		fmt.Println("anonymous go routine ")
		time.Sleep(2 * time.Second)
		fmt.Println("done")
	}()
	time.Sleep(3 * time.Second)
}
// SLIP14Q1
//
//	func main() {
//		numbers := []int{22, 12, 13, 2, 3, 4, 5, 1}
//		sort.Ints(numbers)
//		fmt.Println("sorted arrays is ", numbers)
//	}
//
// SLIP15Q1
// var n int

//	func main() {
//		fmt.Println("enter the size ")
//		fmt.Scan(&n)
//		arr, even, odd := make([]int, n), []int{}, []int{}
//		for i := 0; i < n; i++ {
//			fmt.Scan(&arr[i])
//			if arr[i]%2 == 0 {
//				even = append(even, arr[i])
//			} else {
//				odd = append(odd, arr[i])
//			}
//			//fmt.Println("plese4")
//		}
//		fmt.Println("even is ", even, "odd is ", odd)
//	}
//
// SLIP16Q1
// var n, evencnt, oddcnt int

//	func main() {
//		fmt.Println("enter the size of array ")
//		fmt.Scan(&n)
//		arr := make([]int, n)
//		fmt.Println("enter the element of array ")
//		for i := 0; i < n; i++ {
//			fmt.Scan(&arr[1])
//			if arr[i]%2 == 0 {
//				evencnt++
//			} else {
//				oddcnt++
//			}
//		}
//		fmt.Println("even counts is ", evencnt)
//		fmt.Println("odd count is ", oddcnt)
//	}
//
// SLIP17Q1
// PENDING
// try kiya SLIP18Q2
// SLIP19Q2
// func main() {
// 	var num int
// 	fmt.Println("enter the num ")
// 	fmt.Scan(&num)
// 	fmt.Println("factor of %d are ", num)
// 	finfact(num)

// }
// func finfact(n int) {

// 	for i := 1; i <= n; i++ {
// 		if n%i == 0 {
// 			fmt.Println("%d", i)

//			}
//		}
//		fmt.Println()
//	}
//
// SLIP20Q1
//
//	func main() {
//		var base, exp float64
//		fmt.Println("enter the base ")
//		fmt.Scan(&base)
//		fmt.Println("enter exp ")
//		fmt.Scan(&exp)
//		result := math.Pow(base, exp)
//		fmt.Printf("%.2f %.2f=%.2f\n", base, exp, result)
//	}
//
// SLIP21Q1
// plindrome pending
// SLIP22Q1
// var year int

// func main() {
// 	fmt.Println("enter the year ")
// 	fmt.Scan(&year)
// 	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
// 		fmt.Println("year is leap ")

//		} else {
//			fmt.Println("not leap year ")
//		}
//	}
//
// SLIP24Q1
// func main() {
// 	var num int

//		fmt.Println("enter the number")
//		fmt.Scanln(&num)
//		for i := 1; i <= 10; i++ {
//			fmt.Println(num, "x ", i, "=", num*i)
//		}
//	}
//
// SLIP25Q1
// func main() {
// 	var char rune
// 	fmt.Println("enter the character")
// 	fmt.Scanf("%c", &char)
// 	switch char {
// 	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// 		fmt.Println("its a vowel")
// 	default:
// 		fmt.Println("its a  consonant")

// 	}
// }
//SLIP26Q1
// package main

// import (
// 	"fmt"
// )

// var a int
// var b int

//	func swap(a *int, b *int) {
//		var temp int
//		temp = *a
//		*a = *b
//		*b = temp
//		fmt.Println("the number after swap is :", *a, *b)
//	}
//
//	func main() {
//		fmt.Print("enter first number ")
//		fmt.Scan(&a)
//		fmt.Print("enter second number ")
//		fmt.Scan(&b)
//		fmt.Println("the number befre swap is ", a, b)
//		swap(&a, &b)
//	}
//
// SLIP27Q1
//
//	func main() {
//		numbers := []int{12, 22, 33, 44, 55, 66}
//		min, max := numbers[0], numbers[0]
//		for _, num := range numbers {
//			if num < min {
//				min = num
//			}
//			if num > max {
//				max = num
//			}
//		}
//		fmt.Println("number is samll", min)
//		fmt.Println("number is big ", max)
//	}
//
// SLIP28Q1
//
//	func main() {
//		go func() {
//			fmt.Println("anonymous func ")
//			time.Sleep(2 * time.Second)
//			fmt.Println("its is done ")
//		}()
//		time.Sleep(3 * time.Second)
//	}
//
// SLIP29Q1
// var year int

//	func main() {
//		fmt.Println("enter the year ")
//		fmt.Scan(&year)
//		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
//			fmt.Println("year is leap ")
//		} else {
//			fmt.Println("not leap year ")
//		}
//	}
//
// SLIP30Q1
// func main() {
// 	var num int
// 	fmt.Println("enter the number ")
// 	fmt.Scan(&num)
// 	resilt := isplind(num)
// 	if resilt {
// 		fmt.Println("num is plindrome ", num)
// 	} else {
// 		fmt.Println("num is not plindroeme", num)
// 	}
// }
// func isplind(num int) bool {
// 	originalnum := 0
// 	reversed := 0
// 	for num > 0 {
// 		digit := num % 10
// 		reversed = reversed*10 + digit
// 		num = num / 10

// 	}
// 	return originalnum == reversed
// }
// var num int
// var rem int
// var rev int = 0

// func isplind(num int) {

// 	var temp int = num
// 	for num > 0 {
// 		rem = num % 10
// 		rev = rev*10 + rem
// 		num = num / 10
// 	}
// 	if temp == rev {
// 		fmt.Println("number is plindrome ")
// 	} else {
// 		fmt.Println("number is not plindrome")
// 	}
// }
// func main() {

//		fmt.Println("enter the num ")
//		fmt.Scanf("%d", &num)
//		isplind(num)
//	}
//
//	func main() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(i)
//			delay(1000 * time.Millisecond)
//		}
//	}
//
//	func delay(d time.Duration) {
//		time.Sleep(d)
//	}

// func main() {
// 	var i int
// 	var j int
// 	var c int

// 	fmt.Println("pritn number 1 to 100 ")
// 	for i = 0; i <= 100; i++ {
// 		c = 0

//			for j = 1; j <= i; j++ {
//				if i%j == 0 {
//					c++
//				}
//			}
//			if c == 2 {
//				fmt.Print(i, "\n")
//			}
//		}
//	}
var sum int = 0

func sumofdigit(num int) int {
	if num > 0 {
		sum += (num % 10)
		sumofdigit(num / 10)

	}
	return sum
}
func main() {
	var num int = 0
	var result int = 0
	fmt.Println("enter the number ")
	fmt.Scan(&num)
	result = sumofdigit(num)
	fmt.Println("sum of digit ", result)
}
