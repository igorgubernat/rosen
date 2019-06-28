package ex3

import (
	"fmt"
)

// Universe is universal set
type Universe struct {
	m map[string]float64
}

// NewUniverse is constructor for Universe
func NewUniverse(ss ...string) *Universe {
	m := make(map[string]float64)
	for _, s := range ss {
		m[s] = 0
	}

	return &Universe{
		m: m,
	}
}

// Set returns set in universe
func (u *Universe) Set(m map[string]float64) (*Set, error) {
	for key := range m {
		if _, ok := u.m[key]; !ok {
			return nil, fmt.Errorf("%s is not in universal set", key)
		}
	}

	return &Set{
		universe: u,
		m:        m,
	}, nil
}
