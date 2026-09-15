package main

import "fmt"

type Money int64 // копейки

func (m Money) String() string {
	return fmt.Sprintf("%d.%02d ₽", m/100, m%100)
}

func main() {
	fmt.Println(Money(1500)) // 15.00 ₽ — fmt сам вызвал String()
	fmt.Println(Money(5))    // 0.05 ₽
}
