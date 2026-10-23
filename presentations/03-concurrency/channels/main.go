package main

import "fmt"

func main() {
	ch := make(chan string)        // небуферизованный
	go func() { ch <- "привет" }() // блокируется, пока никто не читает
	msg := <-ch                    // блокируется, пока никто не отправит
	fmt.Println(msg)
}
