package main

import "fmt"

func main() {
	// несколько defer — LIFO
	defer fmt.Println("first deferred")
	defer fmt.Println("second deferred")

	// аргумент вычисляется в момент объявления defer
	i := 1
	defer fmt.Println("deferred i =", i)
	i = 10
	fmt.Println("current i =", i)
}
