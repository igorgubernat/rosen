package ex4

import (
	"testing"
	"reflect"
)

func TestInvert(t *testing.T) {
	type tc struct {
		f          map[int]int
		invert map[int]int
		err bool
	}

	testCases := []tc{
		tc{
			f: map[int]int{
				1: 1,
				2: 1,
				3: 9,
			},
			err: true,
		},
		tc{
			f: map[int]int{
				1: 2,
				2: 3,
				3: 1,
			},
			invert: map[int]int{
				1: 3,
				2: 1,
				3: 2,
			},
		},
	}

	for _, testCase := range testCases {
		got, err := Invert(testCase.f)
		if testCase.err && err == nil || !testCase.err && err != nil {
			t.Fatalf("error want is %t, but got %v", testCase.err, err)
		}

		want := testCase.invert

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v: want %t, but got %t", testCase.f, want, got)
		}
	}
}
