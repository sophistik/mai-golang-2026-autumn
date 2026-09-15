package main

import "fmt"

type Animal struct{ Name string }

func (a Animal) Speak() { fmt.Println(a.Name) }

type Dog struct {
	Animal // встраивание — анонимное поле
	Breed  string
}

// перекрывает «поднятый» метод
func (d Dog) Speak() { fmt.Println(d.Name, "говорит: гав") }

func main() {
	d := Dog{Animal: Animal{Name: "Рекс"}, Breed: "labrador"}

	d.Speak()        // метод Dog: «Рекс говорит: гав»
	d.Animal.Speak() // оригинальный метод доступен: «Рекс»

	fmt.Println(d.Name) // поле тоже «поднялось»
}
