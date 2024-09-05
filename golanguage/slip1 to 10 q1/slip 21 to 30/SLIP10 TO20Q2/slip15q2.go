package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		delay()
	}
}

func delay() {
	time.Sleep(time.Duration(rand.Intn(251)) * time.Millisecond)
}
