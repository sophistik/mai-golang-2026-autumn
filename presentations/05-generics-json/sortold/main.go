package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}
type ByAge []User

func (u ByAge) Len() int           { return len(u) }
func (u ByAge) Less(i, j int) bool { return u[i].Age < u[j].Age }
func (u ByAge) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func main() {
	users := []User{{"Вера", 22}, {"Аня", 20}, {"Боря", 18}}
	sort.Sort(ByAge(users)) // тип-обёртка + три метода на каждый ключ
	fmt.Println(users)      // [{Боря 18} {Аня 20} {Вера 22}]
}
