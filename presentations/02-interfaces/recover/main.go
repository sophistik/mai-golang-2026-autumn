package main

import "fmt"

func safeCall() (result int, err error) {
	defer func() { // именованные возвраты: defer дописывает err
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	panic("something impossible")
}
func main() {
	res, err := safeCall()
	fmt.Println(res, err) // 0 panic: something impossible
}
