package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(50 * time.Millisecond) // «какая-то работа»
		cancel()
	}()

	<-ctx.Done() // канал закрылся — отмена дошла
	fmt.Println(ctx.Err())
}
