package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("из горутины")
	fmt.Println("из main")
	time.Sleep(10 * time.Millisecond) // только для демо!
}
