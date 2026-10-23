package main

import "fmt"

func nums() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; ; i++ {
			out <- i // бесконечный поток
		}
	}()
	return out
}
func main() {
	ch := nums()
	fmt.Println(<-ch, <-ch, <-ch) // 1 2 3
}
