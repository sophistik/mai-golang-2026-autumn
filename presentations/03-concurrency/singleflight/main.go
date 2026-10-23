package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var calls int

func loadFromDB() (any, error) {
	calls++
	time.Sleep(100 * time.Millisecond) // имитация медленного запроса
	return "value(user:42)", nil
}

func main() {
	var g singleflight.Group
	var wg sync.WaitGroup

	// 10 горутин одновременно просят один и тот же ключ
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err, shared := g.Do("user:42", loadFromDB)
			fmt.Printf("got %v, err=%v, shared=%v\n", v, err, shared)
		}()
	}
	wg.Wait()
	fmt.Println("вызовов loadFromDB:", calls) // 1, а не 10
}
