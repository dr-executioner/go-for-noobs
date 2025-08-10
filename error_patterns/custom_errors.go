package errorpatterns

import "fmt"

type ValidationError struct {
	Field string
	Msg   string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Invalid %s: %s", v.Field, v.Msg)
}

func ValidateAge(age int) error{
	if age < 0 || age > 120{
		return ValidationError{"Age", "must be between 0 and 120"}
	}
	return nil
}
