package main

import "fmt"

func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

func Filter[T any](s []T, keep func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	evens := Filter(nums, func(v int) bool { return v%2 == 0 })
	fmt.Println(evens) // [2 4 6]

	labels := Map(evens, func(v int) string { return fmt.Sprint(v, "!") })
	fmt.Println(labels) // [2! 4! 6!]
}
