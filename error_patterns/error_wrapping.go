package errorpatterns

import (
	"errors"
	"fmt"
)

var ErrDatabase = errors.New("database connection failed")

func FetchData() error {
	// Wrap the original error with extra context
	return fmt.Errorf("fetchData failed: %w", ErrDatabase)
}