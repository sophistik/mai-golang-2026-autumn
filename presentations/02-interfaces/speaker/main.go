package main

import "fmt"

type Speaker interface{ Speak() string }

type Dog struct{}
type Cat struct{}

func (Dog) Speak() string { return "Гав" }
func (Cat) Speak() string { return "Мяу" }

func greet(s Speaker) { fmt.Println(s.Speak()) } // не знает ни о Dog, ни о Cat
func main() {
	greet(Dog{}) // компилятор сам видит: у Dog есть Speak() string
	greet(Cat{})
}
