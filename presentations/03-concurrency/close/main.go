package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch) // «данных больше не будет»
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	for v := range ch { // читает до закрытия
		fmt.Println(v)
	}
	v, ok := <-ch // 0 false: закрыт и пуст
	fmt.Println(v, ok)
}
