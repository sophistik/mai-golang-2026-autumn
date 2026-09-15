package main

import "fmt"

func main() {
	s := make([]int, 3, 6) // len=3, cap=6
	copy(s, []int{1, 2, 3})

	sub := s[:2] // окно: len=2, cap=6 — делит массив с s

	// пишем за len(sub), но внутри cap — в ОБЩИЙ массив
	_ = append(sub, 99)

	fmt.Println(s) // [1 2 99] — s[2] затёрт!
}
