package main

import "fmt"

func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	c := counter()
	fmt.Println(c(), c(), c()) // 1 2 3
	fmt.Println(counter()())   // 1 — у нового замыкания своя n
}
