package main

import "fmt"
import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("горутина", i) // Go 1.22+: своя i на итерацию
		}()
	}
	wg.Wait() // блокирует, пока все Done
}
