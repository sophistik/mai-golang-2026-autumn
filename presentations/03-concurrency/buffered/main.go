package main

import "fmt"

func main() {
	ch := make(chan int, 2) // буфер на 2
	ch <- 1                 // не блокирует: есть место
	ch <- 2                 // ещё влезает
	// ch <- 3 // заблокировало бы: буфер полон, читателя нет
	fmt.Println(<-ch, <-ch) // 1 2
}
