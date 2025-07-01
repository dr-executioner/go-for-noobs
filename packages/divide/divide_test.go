package divide

import "testing"

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Error("Unexpected error for valid division")
	}
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}

	_, err = Divide(5, 0)
	if err != ErrDivideByZero {
		t.Error("Expected ErrDivideByZero error")
	}
}
