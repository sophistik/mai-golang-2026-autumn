package main

import "fmt"

const (
	Read    = 1 << iota // 1
	Write               // 2
	Execute             // 4
)

const GiB = 1 << 30 // untyped: произвольная точность

func main() {
	fmt.Println(Read, Write, Execute, GiB)
}
