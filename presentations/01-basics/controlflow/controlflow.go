package main

import "fmt"

func main() {
	if v := len("Go"); v > 0 { // инициализатор: v видна только в if/else
		fmt.Println("positive:", v)
	}
	for i := 0; i < 3; i++ { // классический
		fmt.Print(i, " ")
	}
	x := 1
	for x < 10 { // как while
		x *= 2
	}
	fmt.Println(x)
}
