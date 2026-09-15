package main

import "fmt"

// BinOp — именованный тип функции
type BinOp func(a, b int) int

func main() {
	double := func(x int) int { return x * 2 } // функция — значение
	fmt.Println(double(21))
	ops := map[string]BinOp{ // таблица диспетчеризации
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
	}
	fmt.Println(ops["+"](2, 3), ops["-"](5, 2))
}
