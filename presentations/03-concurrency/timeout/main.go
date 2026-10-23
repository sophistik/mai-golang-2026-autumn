package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel() // освободить таймер, даже если успели

	select {
	case <-time.After(time.Second): // «долгая работа»
		fmt.Println("готово")
	case <-ctx.Done():
		fmt.Println("таймаут:", ctx.Err())
	}
}
