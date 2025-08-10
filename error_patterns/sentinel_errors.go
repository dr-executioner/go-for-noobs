package errorpatterns

import "errors"

var ErrNotFound = errors.New("item not found")

func GetItem(index int) (string, error) {
	items := []string{"apple", "orange", "cherry"}

	if index < 0 || index >= len(items) {
		return "", ErrNotFound
	}
	return items[index], nil
}
