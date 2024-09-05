package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main routine:this is main routine ")
	//creating anonymous go routine
	go func() {
		fmt.Println("anonymous goroutine:")
	}()
	//sleep
	time.Sleep(2 * time.Second)
	fmt.Println("main routine ")
}
