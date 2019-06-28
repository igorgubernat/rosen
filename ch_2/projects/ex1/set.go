package ex1

// Set represents mathematical set
type Set struct {
	universe *UniversalSet
	bits     []uint64
}

// Elements return the list of the elements of the set
func (s *Set) Elements() []string {
	var res []string

	for i := range s.bits {
		for j := 0; j < 64; j++ {
			if (s.bits[i]>>uint(j))&uint64(1) == 1 {
				res = append(res, s.universe.els[i*64+j])
			}
		}
	}

	return res
}

// Complement returns complement of given set
func (s *Set) Complement() *Set {
	bits := make([]uint64, 0, len(s.bits))
	for _, el := range s.bits {
		bits = append(bits, ^el)
	}

	s.mask(bits)

	return &Set{
		universe: s.universe,
		bits:     bits,
	}
}

func (s *Set) mask(bits []uint64) {
	mask := (uint64(1) << uint(len(s.universe.els)%64)) - 1
	bits[len(bits)-1] = bits[len(bits)-1] & mask
}

// Union returns union of two sets
func (s *Set) Union(other *Set) *Set {
	bits := make([]uint64, len(s.bits))

	for i := range bits {
		bits[i] = s.bits[i] | other.bits[i]
	}

	return &Set{
		universe: s.universe,
		bits:     bits,
	}
}

// Intersect returns intersection of two sets
func (s *Set) Intersect(other *Set) *Set {
	bits := make([]uint64, len(s.bits))

	for i := range bits {
		bits[i] = s.bits[i] & other.bits[i]
	}

	return &Set{
		universe: s.universe,
		bits:     bits,
	}
}

// Subtract subtracts argument set from receiver set
func (s *Set) Subtract(other *Set) *Set {
	bits := make([]uint64, len(s.bits))

	for i := range bits {
		bits[i] = s.bits[i] &^ other.bits[i]
	}

	return &Set{
		universe: s.universe,
		bits:     bits,
	}
}

// Difference returns symmetric difference of two sets
func (s *Set) Difference(other *Set) *Set {
	bits := make([]uint64, len(s.bits))

	for i := range bits {
		bits[i] = s.bits[i] ^ other.bits[i]
	}

	return &Set{
		universe: s.universe,
		bits:     bits,
	}
}
