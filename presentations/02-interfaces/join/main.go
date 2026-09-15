package main

import "errors"
import "fmt"

func validateForm(email string, age int) error {
	var errs []error
	if email == "" {
		errs = append(errs, errors.New("email: required"))
	}
	if age < 18 {
		errs = append(errs, errors.New("age: must be 18+"))
	}
	return errors.Join(errs...) // nil, если errs пуст
}
func main() {
	fmt.Println(validateForm("", 15))      // email: required age: must be 18+
	fmt.Println(validateForm("a@b.c", 20)) // <nil>
}
