package greet

import "fmt"

func init() {
	fmt.Println("init: greet") // выполняется ДО init пакета main
}

func Hello(name string) string {
	return "Hello, " + name + "!"
}
