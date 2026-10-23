package main

import "fmt"
import "sync"

func main() {
	var wg sync.WaitGroup
	n := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n++ // ГОНКА: прочитать, сложить, записать
		}()
	}
	wg.Wait()
	fmt.Println(n) // каждый запуск — новое число
}
