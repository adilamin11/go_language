package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		delay(100 * time.Millisecond)
	}
}

func delay(d time.Duration) {
	time.Sleep(d)
}
