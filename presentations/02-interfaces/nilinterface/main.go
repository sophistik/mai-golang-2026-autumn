package main

import "fmt"

type Seller interface{ GetID() int64 }
type Store struct{}

func (st *Store) GetID() int64 { return 1 }

func main() {
	var seller Seller          // (nil, nil)
	fmt.Println(seller == nil) // true

	var store *Store           // nil, но тип уже есть
	seller = store             // (*Store, nil)
	fmt.Println(seller == nil) // false — типизированный nil!
}
