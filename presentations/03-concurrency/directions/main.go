package main

import "fmt"

func producer(out chan<- int) { // только отправка
	out <- 1
	close(out)
}
func consumer(in <-chan int) { // только приём
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch) // перепутать не даст компилятор
}
