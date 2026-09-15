package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var f float64 = float64(i) // явное преобразование
	var u uint = uint(f)
	fmt.Println(i, f, u)
	n, err := strconv.Atoi("42") // возвращает ошибку!
	fmt.Println(n, err)
	s := strconv.Itoa(42)
	fmt.Println(s)
}
