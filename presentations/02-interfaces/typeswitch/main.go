package main

import "fmt"

func describe(i any) {
	switch v := i.(type) {
	case string:
		fmt.Println("string:", v) // v здесь — string
	case nil:
		fmt.Println("nil")
	default:
		fmt.Printf("unknown: %T\n", v)
	}
}
func main() {
	describe("hello")
	describe(42) // default: unknown: int
	describe(nil)
}
