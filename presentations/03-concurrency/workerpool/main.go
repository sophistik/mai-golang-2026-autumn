package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("воркер", id, "остановлен")
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			time.Sleep(30 * time.Millisecond) // имитация работы
			results <- j * j
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	go func() { // fan-in: собираем результаты, пока пул работает
		for r := range results {
			fmt.Println("результат:", r)
		}
	}()

produce:
	for j := 1; ; j++ {
		select {
		case jobs <- j:
		case <-ctx.Done():
			break produce // таймаут: новых заданий не будет
		}
	}
	close(jobs)    // задач больше нет — воркеры завершатся
	wg.Wait()      // дождались всех воркеров
	close(results) // теперь результаты закончились
	fmt.Println("пул остановлен по таймауту")
}
