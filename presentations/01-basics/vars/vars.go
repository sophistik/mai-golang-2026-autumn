package main

import "fmt"

func main() {
	var x int    // 0
	var s string // ""
	var ok bool  // false
	var p *int   // nil
	var sl []int // nil
	fmt.Println(x, s == "", ok, p == nil, sl == nil)

	y := 10 // краткая форма — только внутри функций
	fmt.Println(y)
}
