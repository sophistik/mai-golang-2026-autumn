package main

import "fmt"

func divmod(a, b int) (int, int) { return a / b, a % b }
func sum(nums ...int) int { // nums внутри — обычный []int
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
func main() {
	q, r := divmod(17, 5)
	fmt.Println(q, r)
	fmt.Println(sum(1, 2, 3), sum([]int{1, 2, 3}...)) // развёртывание слайса
}
