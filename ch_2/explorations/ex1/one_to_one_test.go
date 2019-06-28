package ex1

import (
	"testing"
)

func TestOneToOne(t *testing.T) {
	type tc struct {
		s int
		t int
		n int
	}

	testCases := []tc{
		tc{
			s: 10,
			t: 5,
			n: 0,
		},
		tc{
			s: 4,
			t: 4,
			n: 24,
		},
		tc{
			s: 3,
			t: 5,
			n: 60,
		},
	}

	for _, testCase := range testCases {
		got := OneToOne(testCase.s, testCase.t)
		want := testCase.n

		if got != want {
			t.Fatalf("want %d, but got %d", want, got)
		}
	}
}
