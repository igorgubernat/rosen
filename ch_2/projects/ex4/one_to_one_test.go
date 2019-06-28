package ex4

import (
	"testing"
)

func TestOneToOne(t *testing.T) {
	if _, err := IsOneToOne([][2]int{
		[2]int{0, 0},
		[2]int{0, 1},
	}); err == nil {
		t.Fatal("want error, got nil")
	}

	type tc struct {
		f          [][2]int
		isOneToOne bool
	}

	testCases := []tc{
		tc{
			f: [][2]int{
				[2]int{1, 1},
				[2]int{2, 4},
				[2]int{3, 9},
			},
			isOneToOne: true,
		},
		tc{
			f: [][2]int{
				[2]int{1, 1},
				[2]int{2, 1},
				[2]int{3, 9},
			},
			isOneToOne: false,
		},
	}

	for _, testCase := range testCases {
		got, err := IsOneToOne(testCase.f)
		if err != nil {
			t.Error(err)
		}

		want := testCase.isOneToOne

		if got != want {
			t.Errorf("%v: want %t, but got %t", testCase.f, want, got)
		}
	}
}
