package ex4

import (
	"fmt"
	"sort"
	"reflect"
)

// Invert returns inverse function
func Invert(f map[int]int) (map[int]int, error) {
	if !isBijection1n(f) {
		return nil, fmt.Errorf("%v is not bijection from 1, 2, ... , n to itself", f)
	}

	res := make(map[int]int)

	for x, y := range f {
		res[y] = x
	}

	return res, nil
}

func isBijection1n(f map[int]int) bool {
	var domain, rang []int

	for key, val := range f {
		domain = append(domain, key)
		rang = append(rang, val)
	}

	sort.Ints(domain)
	sort.Ints(rang)

	var ref []int
	for i := 1; i <= domain[len(domain)-1]; i++ {
		ref = append(ref, i)
	}

	if reflect.DeepEqual(domain, ref) && reflect.DeepEqual(rang, ref) {
		return true
	} else {
		return false
	}
}