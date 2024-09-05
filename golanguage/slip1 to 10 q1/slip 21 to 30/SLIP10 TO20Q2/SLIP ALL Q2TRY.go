package main

// import "fmt"

// var average int
// var total int

// type student struct {
// 	rollno int
// 	name   string
// 	marks1 int
// 	marks2 int
// 	marks3 int
// }

//	func main() {
//		var n, i int
//		fmt.Println("how many student uhh want ")
//		fmt.Scan(&n)
//		obj := make([]student, n)
//		for i = 0; i < n; i++ {
//			fmt.Println("enter student roll no ")
//			fmt.Scan(&obj[i].rollno)
//			fmt.Println("enter stu name ")
//			fmt.Scan(&obj[i].name)
//			fmt.Println("enter marks1 ")
//			fmt.Scan(&obj[i].marks1)
//			fmt.Println("enter marks 2")
//			fmt.Scan(&obj[i].marks2)
//			fmt.Println("enter marks3 ")
//			fmt.Scan(&obj[i].marks3)
//			fmt.Println()
//		}
//		fmt.Println("---display stu details ")
//		for i = 0; i < n; i++ {
//			display(obj[i])
//		}
//	}
//
//	func display(obj student) {
//		total = obj.marks1 + obj.marks2 + obj.marks3
//		average = total / 3
//		fmt.Println("stud rollno ", obj.rollno)
//		fmt.Println("name ", obj.name)
//		fmt.Println("total marks ", total)
//		fmt.Println("average ", average)
//		fmt.Println()
//	}
//
// SLIP2Q2
// func main() {

// 	var rows, columns, i, j int
// 	var mat1 [10][10]int
// 	var mat2 [10][10]int
// 	var multimat [10][10]int
// 	fmt.Println("enter rows ")
// 	fmt.Scan(&rows)
// 	fmt.Println("enter colums ")
// 	fmt.Scan(&columns)
// 	fmt.Println("enter the first matrics ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Scan(&mat1[i][j])
// 		}
// 	}
// 	fmt.Println("enter second matrics ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Scan(&mat2[i][j])
// 		}
// 	}
// 	fmt.Println("the first matric is ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Print(mat1[i][j], "\t")
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println("the second met is ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Print(mat2[i][j], "\t")
// 		}
// 		fmt.Println()
// 	}
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			multimat[i][j] = mat1[i][j] * mat2[i][j]
// 		}
// 	}
// 	fmt.Println("the result of mat multipl is ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Print(multimat[i][j], "\t")
// 		}
// 		fmt.Println()
// 	}
// }

// SLIP3Q2
// func main() {
// 	for i := 0; i <= 10; i++ {
// 		fmt.Println(i)
// 		delay(100 * time.Millisecond)
// 	}
// }
// func delay(d time.Duration) {
// 	time.Sleep(d)
// }
//SLIP4Q2
// func main() {
// 	var rows, columns, i, j int
// 	var mat1 [10][10]int
// 	var mat2 [10][10]int
// 	var multimat [10][10]int
// 	fmt.Print("enter rows ")
// 	fmt.Scan(&rows)
// 	fmt.Println("enter colums ")
// 	fmt.Scan(&columns)
// 	fmt.Println("enter the first matrics ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Scan(&mat1[i][j])
// 		}
// 	}
// 	fmt.Println("enter the second mat")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Scan(&mat2[i][j])
// 		}
// 	}
// 	fmt.Println("the first mat is ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Print(mat1[i][j], "\t")
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println("the second matric is ")
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			fmt.Print(mat2[i][j], "\t")
// 		}
// 		fmt.Println()
// 	}
// 	for i = 0; i < rows; i++ {
// 		for j = 0; j < columns; j++ {
// 			multimat[i][j] = mat1[i][j] + mat2[i][j]

//			}
//			fmt.Println()
//		}
//		fmt.Println("the matrci addition is ")
//		for i = 0; i < rows; i++ {
//			for j = 0; j < columns; j++ {
//				fmt.Print(multimat[i][j], "\t")
//			}
//			fmt.Println()
//		}
//	}
//
// SLIP5Q2
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }
// type Circle struct {
// 	radius float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// type rectangle struct {
// 	length float64
// 	width  float64
// }

// func (r rectangle) area() float64 {
// 	return (r.length * r.width)
// }
// func (r rectangle) Perimeter() float64 {
// 	return 2 * (r.length + r.width)
// }
// func main() {
// 	c := Circle{
// 		radius: 9,
// 	}
// 	fmt.Println("area of circle is ", c.Area())
// 	fmt.Println("perimeter of circle is ", c.Perimeter())
// 	r := rectangle{
// 		length: 8,
// 		width:  6,
// 	}
// 	fmt.Println("area of rect is ", r.area())
// 	fmt.Println("per is ", r.Perimeter())
// }

// SLIP6Q2
// pending bro
// slip7 q2
// var max int = 0

// type emp struct {
// 	eno    int
// 	ename  string
// 	salary int
// }

// func main() {
// 	var n, i, j int
// 	fmt.Println("enter how many studen uh want ")
// 	fmt.Scan(&n)
// 	obj := make([]emp, n)
// 	for i = 0; i < n; i++ {
// 		fmt.Println("enter emp no ")
// 		fmt.Scan(&obj[i].eno)
// 		fmt.Println("enter emp name ")
// 		fmt.Scan(&obj[i].ename)
// 		fmt.Println("enter emp salary ")
// 		fmt.Scan(&obj[i].salary)
// 		fmt.Println()
// 	}
// 	fmt.Println("--display emp details whose salry is max ")
// 	max = int(obj[0].salary)
// 	for i = 0; i < n; i++ {
// 		if obj[i].salary > max {
// 			max = int(obj[i].salary)
// 		}
// 	}
// 	for j = 0; j < n; j++ {
// 		if max == int(obj[j].salary) {
// 			display(obj[j])
// 		}
// 	}

// }
//
//	func display(obj emp) {
//		fmt.Println("emp no is ", obj.eno)
//		fmt.Println("emp name is ", obj.ename)
//		fmt.Println("emp salary is ", obj.salary)
//	}
//
// SLIP8Q2
// tryed
// SLIP9Q2
// func main() {
// 	var i int
// 	var j int
// 	var c int
// 	fmt.Println("print 1 to 100 prime number ")
// 	for i = 1; i <= 100; i++ {
// 		c = 0
// 		for j = 1; j <= i; j++ {
// 			if i%j == 0 {
// 				c++
// 			}
// 		}
// 		if c == 2 {
// 			fmt.Println(i, "\n\n")
// 		}
// 	}

// }
// slip10q2
// func main() {
// 	var rows int
// 	var temp int = 1
// 	fmt.Println("enter the no of rows ")
// 	fmt.Scan(&rows)
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j <= rows-i; j++ {
// 			fmt.Println(" ")
// 		}
// 		for k := 0; k <= i; k++ {

// 			if k == 0 || i == 0 {
// 				temp = 1

// 			} else {
// 				temp = temp * (i - k + 1) / k
// 			}
// 			fmt.Printf("%d", temp)
// 		}
// 		fmt.Println()
// 	}

// }
// SLIP11Q2
// type student struct {
// 	name   string
// 	branch string
// 	rollno int
// }

// func (s *student) show(sbranch string) {
// 	(*s).branch = sbranch
// }
// func main() {
// 	res := student{
// 		name:   "adil",
// 		branch: "BCA",
// 	}
// 	fmt.Println("students name ", res.name)
// 	fmt.Println("stude branch ", res.branch)
// 	p := &res
// 	p.show("MCA")
// 	fmt.Println("stud name ", res.name)
// 	fmt.Println("stud  brach", res.branch)
// }
//SLIP12Q2
//pending hai bro
//SLIP13Q2
// type book struct {
// 	b_id int

// 	title string

// 	author string

// 	price int
// }

// func main() {
// 	var n, i int
// 	fmt.Println("how many records uhh want to store ")
// 	fmt.Scan(&n)
// 	obj := make([]book, n)
// 	for i = 0; i < n; i++ {
// 		fmt.Println("enter bid ")
// 		fmt.Scan(&obj[i].b_id)
// 		fmt.Println("enter book title ")
// 		fmt.Scan(&obj[i].title)
// 		fmt.Println("enter book author ")
// 		fmt.Scan(&obj[i].author)
// 		fmt.Println("enter book price ")
// 		fmt.Scan(&obj[i].price)
// 		fmt.Println()
// 	}
// 	fmt.Print("--display vbook details ")
// 	for i := 0; i < n; i++ {
// 		display(obj[i])
// 	}
// }
// func display(obj book) {
// 	fmt.Println("book id is ", obj.b_id)
// 	fmt.Println("book title is ", obj.title)
// 	fmt.Println("book author is ", obj.author)
// 	fmt.Println("book price is ", obj.price)
// }
//SLIP14Q2

// type student struct {
// 	name string

// 	branch string

// 	rollno int
// }

//	func (s *student) show(sbranch string) {
//		(*s).branch = sbranch
//	}
//
//	func main() {
//		res := student{
//			name:   "adil",
//			branch: "bca",
//		}
//		fmt.Println("enter name ", res.name)
//		fmt.Println("enter branch ", res.branch)
//		p := &res
//		p.show("mca")
//		fmt.Println("name is ", res.name)
//		fmt.Println("branch is ", res.branch)
//	}
//
// SLIP15Q2
//
//	func main() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(i)
//			delay()
//		}
//	}
//
//	func delay() {
//		time.Sleep(time.Duration(rand.Intn(251)) * time.Millisecond)
//	}
//
// SLIP16Q2
// pending of file
// slip1q2
// pending
// SLIP18Q2
//
//	func main() {
//		var str1 string
//		var str2 string
//		var p1 *string = &str1
//		var p2 *string = &str2
//		fmt.Println("enter the first struing")
//		fmt.Scan(&str1)
//		fmt.Println("enter the second string ")
//		fmt.Scan(&str2)
//		res := *p1 + *p2
//		fmt.Println("cont of string is ", res)
//	}
//
// SLIP20Q2
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
//
// SLIP21Q2
// type student struct {
// 	rollno int
// 	name   string
// 	marks  int
// }

// func (s *student) show(smarks int) {
// 	(*s).marks = smarks
// }
// func main() {
// 	res := student{
// 		name:  "adil",
// 		marks: 90,
// 	}
// 	fmt.Println("stud name ", res.name)
// 	fmt.Println("stud marks (before)", res.marks)
// 	p := &res
// 	p.show(80)
// 	fmt.Println("stud name ", res.name)
// 	fmt.Println("stud marks (after)", res.marks)
// }
//SLIP22Q2
//tryed
//SLIP23Q2
// var total int
// var average int

// type student struct {
// 	rollno int
// 	name   string
// 	marks1 int
// 	marks2 int
// 	marks3 int
// }

// func main() {
// 	var n, i int
// 	fmt.Println("enter how many st uhh want")
// 	fmt.Scan(&n)
// 	obj := make([]student, n)
// 	for i = 0; i < n; i++ {
// 		fmt.Println("enter st roll no ")
// 		fmt.Scan(&obj[i].rollno)
// 		fmt.Println("enter anem ")
// 		fmt.Scan(&obj[i].name)
// 		fmt.Println("enter mark 1 ")
// 		fmt.Scan(&obj[i].marks1)
// 		fmt.Println("enter marks 2")
// 		fmt.Scan(&obj[i].marks2)
// 		fmt.Println("enter marks 3 ")
// 		fmt.Scan(&obj[i].marks3)
// 		fmt.Println()
// 	}
// 	fmt.Println("--display st det ")
// 	for i := 0; i < n; i++ {
// 		display(obj[i])
// 	}
// }
// func display(obj student) {
// 	total = obj.marks1 + obj.marks2 + obj.marks3
// 	average = total / 3
// 	fmt.Println("stude name is ", obj.name)
// 	fmt.Println("stude name is ", obj.rollno)
// 	fmt.Println("stude total is ", total)
// 	fmt.Println("stude  average is ", average)

// 	fmt.Println()
// }
