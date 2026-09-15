package main

import "fmt"

func main() {
	var s []int // nil-слайс: len=0, cap=0 — рабочее значение
	s1 := []int{1, 2, 3}
	s2 := make([]int, 5)     // len=5, cap=5, заполнен нулями
	s3 := make([]int, 0, 10) // len=0, cap=10

	fmt.Println(s == nil, len(s1), len(s2), cap(s3))
}
