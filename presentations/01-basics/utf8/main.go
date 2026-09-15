package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "мир"
	fmt.Println("bytes:", len(s))                    // 6
	fmt.Printf("hex:   % x\n", s)                    // d0 bc d0 b8 d1 80
	fmt.Println("runes:", utf8.RuneCountInString(s)) // 3
	for i, r := range s {                            // индексы 0, 2, 4 — не подряд!
		fmt.Printf("%d: %c (%U)\n", i, r, r)
	}
}
