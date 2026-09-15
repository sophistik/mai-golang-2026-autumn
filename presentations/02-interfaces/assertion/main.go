package main

import "fmt"

func main() {
	var i any = "hello"

	s, ok := i.(string) // безопасная форма
	fmt.Println(s, ok)  // hello true

	n, ok := i.(int)
	fmt.Println(n, ok) // 0 false

	// n = i.(int) // panic: interface conversion
}
