package ex1

import (
	"fmt"
	"sort"
)

// UniversalSet is Universe in which sets are defined
type UniversalSet struct {
	els []string
}

// NewUniversalSet is constructor of UniversalSet
func NewUniversalSet(els ...string) *UniversalSet {
	els = dedup(els)
	sort.Strings(els)

	return &UniversalSet{
		els: els,
	}
}

// isIn returns true if all elements in els are from universal set
func (u *UniversalSet) isIn(els []string) bool {
	m := make(map[string]bool, len(u.els))

	for _, el := range u.els {
		m[el] = true
	}

	for _, el := range els {
		if _, ok := m[el]; !ok {
			return false
		}
	}

	return true
}

// Set returns set in universal set that is defined by els
func (u *UniversalSet) Set(els ...string) (*Set, error) {
	if !u.isIn(els) {
		return nil, fmt.Errorf("%v is not in universe %v", els, u.els)
	}

	els = dedup(els)
	sort.Strings(els)

	bitLen := len(u.els) / 64
	if len(u.els)%64 != 0 {
		bitLen++
	}
	bits := make([]uint64, bitLen)

	for _, el := range els {
		for i := range u.els {
			if el == u.els[i] {
				index := i / 64
				offset := i % 64
				bits[index] = bits[index] | (uint64(1) << uint(offset))
			}
		}
	}

	return &Set{
		universe: u,
		bits:     bits,
	}, nil
}

// dedup removes duplicates from els
func dedup(els []string) []string {
	m := make(map[string]bool)

	for _, el := range els {
		m[el] = true
	}

	res := make([]string, 0, len(m))

	for el := range m {
		res = append(res, el)
	}

	return res
}
