package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("anonymous go routine ")
		time.Sleep(2 * time.Second)
		fmt.Println("done")
	}()
	time.Sleep(3 * time.Second)
}
