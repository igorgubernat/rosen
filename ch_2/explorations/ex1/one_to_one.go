package ex1

func OneToOne(s, t int) int {
	if s > t {
		return 0
	}

	return fact(t) / fact(t-s)
}

func fact(a int) int {
	if a <= 0 {
		return 1
	}

	res := 1
	for i := 1; i <= a; i++ {
		res *= i
	}

	return res
}
