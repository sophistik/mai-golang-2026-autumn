package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func loadUser(id int64) error {
	return fmt.Errorf("load user %d: %w", id, ErrNotFound)
}
func main() {
	err := loadUser(0)
	fmt.Println(err)                         // load user 0: not found
	fmt.Println(errors.Is(err, ErrNotFound)) // true — %w сохранил цепочку
}
