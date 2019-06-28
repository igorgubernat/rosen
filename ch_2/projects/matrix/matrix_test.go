package matrix

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	if _, err := NewMatrix([][]float64{
		[]float64{1, 2, 3},
		[]float64{4, 5},
	}); err == nil {
		t.Fatal("want error but got nil")
	}

	m, err := NewMatrix([][]float64{
		[]float64{1, 2, 3},
		[]float64{4, 5, 6},
	})
	if err != nil {
		t.Fatal(err)
	}

	if m.numRows != 2 {
		t.Fatalf("want 2 rows, but got %d", m.numRows)
	}

	if m.numCols != 3 {
		t.Fatalf("want 3 columns, but got %d", m.numCols)
	}
}

func TestX(t *testing.T) {
	m1, err := NewMatrix([][]float64{
		[]float64{0, 1, 2},
		[]float64{2, 3, 1},
	})

	m2, err := NewMatrix([][]float64{
		[]float64{0, 1, 2},
		[]float64{2, 3, 1},
	})

	if _, err := m1.X(m2); err == nil {
		t.Fatal("want error, got nil")
	}

	m3, err := NewMatrix([][]float64{
		[]float64{0, 1},
		[]float64{2, 3},
		[]float64{1, 1},
	})

	got, err := m1.X(m3)
	if err != nil {
		t.Fatal(err)
	}

	want, _ := NewMatrix([][]float64{
		[]float64{4, 5},
		[]float64{7, 12},
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}
