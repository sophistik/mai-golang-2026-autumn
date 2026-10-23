package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() { c1 <- "результат" }()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("таймаут")
	}
}
