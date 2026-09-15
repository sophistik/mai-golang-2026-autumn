package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	m["c"] = 3
	v, ok := m["x"] // 0, false — comma ok: «нет ключа» ≠ «нулевое значение»
	fmt.Println(v, ok, m["x"])
	delete(m, "a")
	for k, v := range m { // порядок итерации случаен
		fmt.Println(k, v)
	}
}
