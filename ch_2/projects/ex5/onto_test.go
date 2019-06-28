package ex4

import (
	"testing"
)

func TestOnto(t *testing.T) {
	if _, err := IsOnto([][2]int{
		[2]int{0, 0},
		[2]int{0, 1},
	}); err == nil {
		t.Fatal("want error, got nil")
	}

	type tc struct {
		f          [][2]int
		isOnto bool
	}

	testCases := []tc{
		tc{
			f: [][2]int{
				[2]int{1, 1},
				[2]int{2, 3},
				[2]int{3, 2},
			},
			isOnto: true,
		},
		tc{
			f: [][2]int{
				[2]int{1, 1},
				[2]int{2, 1},
				[2]int{3, 2},
			},
			isOnto: false,
		},
	}

	for _, testCase := range testCases {
		got, err := IsOnto(testCase.f)
		if err != nil {
			t.Error(err)
		}

		want := testCase.isOnto

		if got != want {
			t.Errorf("%v: want %t, but got %t", testCase.f, want, got)
		}
	}
}
