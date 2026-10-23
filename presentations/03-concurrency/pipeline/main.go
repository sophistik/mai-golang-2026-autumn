package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in { // вход закрыт и пуст → закрываем выход
			out <- n * n
		}
	}()
	return out
}

func main() {
	// стадии композируются: gen → sq → sq
	for n := range sq(sq(gen(1, 2, 3))) {
		fmt.Println(n) // 1 16 81
	}
}
