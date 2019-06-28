package ex2

// Multiset represents multiset (set with duplicates)
type Multiset struct {
	m map[string]int
}

// NewMultiset is constructor of Multiset
func NewMultiset(ss ...string) *Multiset {
	m := make(map[string]int)

	for _, s := range ss {
		m[s]++
	}

	return &Multiset{
		m: m,
	}
}

// Elements returns slice of elements of multiset
func (ms *Multiset) Elements() []string {
	var ss []string

	for el, mul := range ms.m {
		for i := 1; i <= mul; i++ {
			ss = append(ss, el)
		}
	}

	return ss
}

// Union returns union of two multisets
func (ms *Multiset) Union(other *Multiset) *Multiset {
	m := make(map[string]int)

	for el, mul := range ms.m {
		var max int
		if mul > other.m[el] {
			max = mul
		} else {
			max = other.m[el]
		}

		m[el] = max
	}

	for el, mul := range other.m {
		if _, ok := ms.m[el]; !ok {
			m[el] = mul
		}
	}

	return &Multiset{
		m: m,
	}
}

// Intersect returns intersection of two multisets
func (ms *Multiset) Intersect(other *Multiset) *Multiset {
	m := make(map[string]int)

	for el, mul := range ms.m {
		if omul, ok := other.m[el]; ok {
			var min int
			if mul < omul {
				min = mul
			} else {
				min = omul
			}

			m[el] = min
		}
	}

	return &Multiset{
		m: m,
	}
}

// Subtract subtracts argument multiset from receiver multiset
func (ms *Multiset) Subtract(other *Multiset) *Multiset {
	m := make(map[string]int)

	for el, mul := range ms.m {
		diff := mul - other.m[el]
		if diff > 0 {
			m[el] = diff
		}
	}

	return &Multiset{
		m: m,
	}
}

// Add adds two multisets
func (ms *Multiset) Add(other *Multiset) *Multiset {
	m := make(map[string]int)

	for el, mul := range ms.m {
		m[el] = mul
	}

	for el, mul := range other.m {
		m[el] += mul
	}

	return &Multiset{
		m: m,
	}
}
