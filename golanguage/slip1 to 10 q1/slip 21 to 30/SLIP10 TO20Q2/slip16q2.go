package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("hello.txt")
	defer f.Close()
	f.WriteString("hello world\n")
	fmt.Println("file craeted ")
}
