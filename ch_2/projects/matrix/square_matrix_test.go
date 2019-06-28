package matrix

import (
	"reflect"
	"testing"
)

func TestNewSquareMatrix(t *testing.T) {
	if _, err := NewSquareMatrix([][]float64{
		[]float64{1, 2},
		[]float64{3, 4},
		[]float64{5, 6},
	}); err == nil {
		t.Fatal("want err, got nil")
	}

	if _, err := NewSquareMatrix([][]float64{
		[]float64{1, 2, 3},
		[]float64{3, 4, 5},
		[]float64{5, 6, 7},
	}); err != nil {
		t.Fatal(err)
	}
}

func TestPower(t *testing.T) {
	m, err := NewSquareMatrix([][]float64{
		[]float64{2, 0, 0},
		[]float64{0, 2, 0},
		[]float64{0, 0, 2},
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := m.Power(3)
	if err != nil {
		t.Fatal(err)
	}

	want, _ := NewSquareMatrix([][]float64{
		[]float64{8, 0, 0},
		[]float64{0, 8, 0},
		[]float64{0, 0, 8},
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestIsSymmetric(t *testing.T) {
	m1, err := NewSquareMatrix([][]float64{
		[]float64{2, 0, 0},
		[]float64{0, 2, 0},
		[]float64{0, 0, 2},
	})
	if err != nil {
		t.Fatal(err)
	}

	if !m1.IsSymmetric() {
		t.Fatalf("matrix %v is symmetric, but got false", m1)
	}

	m2, err := NewSquareMatrix([][]float64{
		[]float64{2, 0, 1},
		[]float64{0, 2, 0},
		[]float64{0, 0, 2},
	})
	if err != nil {
		t.Fatal(err)
	}

	if m2.IsSymmetric() {
		t.Fatalf("matrix %v is not symmetric, but got true", m2)
	}
}
