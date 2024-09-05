package main

func main() {
	var a int = 100
	var b int = 200
	swap(&a, &b)
}
func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}
