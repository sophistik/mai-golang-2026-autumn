package main

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(Max(3, 7))        // T = int — выведен сам
	fmt.Println(Max("go", "c++")) // T = string
	fmt.Println(Max[int](3, 7))   // тип можно указать явно
}
