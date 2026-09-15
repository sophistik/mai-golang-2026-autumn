package main

import (
	"fmt"

	"go-course/01-basics/packages/internal/greet"
)

func init() {
	fmt.Println("init: package main")
}

func main() {
	fmt.Println(greet.Hello("Go"))
	// fmt.Println(greet.hidden()) // compile error: greet.hidden is not exported
}
