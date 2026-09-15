package main

import "fmt"

func main() {
	switch n := 2; n { // break встроен, проваливания нет
	case 1, 2:
		fmt.Println("один или два")
	default:
		fmt.Println("другое")
	}
	switch x := 5; { // безусловный — замена цепочке if/else if
	case x < 0:
		fmt.Println("negative")
	default:
		fmt.Println("positive")
	}
}
