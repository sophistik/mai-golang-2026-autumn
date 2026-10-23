package main

import (
	"cmp"
	"fmt"
	"slices"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{{"Вера", 22}, {"Аня", 20}, {"Боря", 18}}
	slices.SortFunc(users, func(a, b User) int { return cmp.Compare(a.Age, b.Age) })
	fmt.Println(users) // [{Боря 18} {Аня 20} {Вера 22}]
}
