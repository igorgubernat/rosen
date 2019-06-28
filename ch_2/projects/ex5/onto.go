package ex4

import (
	"fmt"
)

// IsOnto returns true if specified function is onto
func IsOnto(f [][2]int) (bool, error) {
	m := make(map[int]bool)
	for _, el := range f {
		m[el[0]] = true
	}

	if len(m) != len(f) {
		return false, fmt.Errorf("%v is not a function", f)
	}

	domain := make(map[int]bool)
	for _, el := range f {
		domain[el[1]] = true
	}

	var max int
	for el := range m {
		if el > max {
			max = el
		}
	}

	for i := 1; i <= max; i++ {
		if _, ok := domain[i]; !ok {
			return false, nil
		}
	}

	return true, nil
}
