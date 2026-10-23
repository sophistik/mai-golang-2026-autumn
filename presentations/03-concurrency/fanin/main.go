package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// fan-out: несколько обработчиков делят один входной канал
func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in { // канал сам отдаёт элемент тому, кто успел
			out <- n * n
		}
	}()
	return out
}

// fan-in: сливаем несколько выходных каналов в один
func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, ch := range chans {
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}
	go func() {
		wg.Wait()  // все источники закрыты
		close(out) // значит, закрыт и результат
	}()
	return out
}

func main() {
	in := gen(1, 2, 3, 4, 5, 6)

	out := merge(squarer(in), squarer(in), squarer(in))

	for v := range out {
		fmt.Println(v) // 1 4 9 16 ... порядок недетерминирован
	}
}
