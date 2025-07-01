package divide

import "errors"

var ErrDivideByZero = errors.New("cannot divide by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
