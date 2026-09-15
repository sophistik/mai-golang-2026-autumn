package custom

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("field %q: %s", e.Field, e.Message)
}
func validateEmail(email string) error {
	if email == "" {
		return &ValidationError{Field: "email", Message: "required"}
	}
	return nil
}
