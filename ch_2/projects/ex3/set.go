package ex3

// Set is a fuzzy set
type Set struct {
	universe *Universe
	m        map[string]float64
}

// Complement returns complement of given set
func (s *Set) Complement() *Set {
	m := make(map[string]float64)

	for key := range s.universe.m {
		degree := 1 - s.m[key]
		if degree > 0 {
			m[key] = degree
		}
	}

	return &Set{
		universe: s.universe,
		m:        m,
	}
}

// Union returns union of two sets
func (s *Set) Union(other *Set) *Set {
	m := make(map[string]float64)

	for el, degree := range s.m {
		m[el] = degree
	}

	for el, degree := range other.m {
		if degree > m[el] {
			m[el] = degree
		}
	}

	return &Set{
		universe: s.universe,
		m:        m,
	}
}

// Intersection returns intersection of two sets
func (s *Set) Intersection(other *Set) *Set {
	m := make(map[string]float64)

	for el, degree := range s.m {
		if other.m[el] > 0 {
			if degree < other.m[el] {
				m[el] = degree
			} else {
				m[el] = other.m[el]
			}
		}
	}

	return &Set{
		universe: s.universe,
		m:        m,
	}
}
