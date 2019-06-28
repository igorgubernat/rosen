package ex4

import (
	"fmt"
)

// IsOneToOne returns true if specified function is one to one
func IsOneToOne(f [][2]int) (bool, error) {
	m := make(map[int]bool)
	for _, el := range f {
		m[el[0]] = true
	}

	if len(m) != len(f) {
		return false, fmt.Errorf("%v is not a function", f)
	}

	m = make(map[int]bool)
	for _, el := range f {
		m[el[1]] = true
	}

	if len(m) == len(f) {
		return true, nil
	} else {
		return false, nil
	}
}
