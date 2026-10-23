package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func fetch(ctx context.Context, url string) error {
	select {
	case <-time.After(100 * time.Millisecond): // имитация сетевого запроса
		fmt.Println("ok:", url)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	urls := []string{
		"https://example.com/1",
		"https://example.com/broken",
		"https://example.com/3",
	}

	g, ctx := errgroup.WithContext(context.Background())
	for _, u := range urls {
		g.Go(func() error {
			if u == "https://example.com/broken" {
				return errors.New("503 " + u)
			}
			return fetch(ctx, u)
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("ошибка:", err)
	}
}
