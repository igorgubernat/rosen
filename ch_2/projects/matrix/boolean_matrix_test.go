package matrix

import (
	"reflect"
	"testing"
)

func TestNewBooleanMatrix(t *testing.T) {
	if _, err := NewBooleanMatrix([][]int{
		[]int{1, 0, 1},
		[]int{1, 1},
	}); err == nil {
		t.Fatal("want error but got nil")
	}

	if _, err := NewBooleanMatrix([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}); err == nil {
		t.Fatal("want error but got nil")
	}

	m, err := NewBooleanMatrix([][]int{
		[]int{1, 0, 1},
		[]int{0, 1, 1},
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

func TestBooleanX(t *testing.T) {
	m1, err := NewBooleanMatrix([][]int{
		[]int{0, 1, 1},
		[]int{1, 0, 1},
	})
	if err != nil {
		t.Fatal(err)
	}

	m2, err := NewBooleanMatrix([][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := m1.X(m2); err == nil {
		t.Fatal("want error, got nil")
	}

	m3, err := NewBooleanMatrix([][]int{
		[]int{0, 1},
		[]int{0, 0},
		[]int{1, 0},
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := m1.X(m3)
	if err != nil {
		t.Fatal(err)
	}

	want, _ := NewBooleanMatrix([][]int{
		[]int{1, 0},
		[]int{1, 1},
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestJoin(t *testing.T) {
	m1, err := NewBooleanMatrix([][]int{
		[]int{0, 1, 1},
		[]int{1, 0, 1},
	})
	if err != nil {
		t.Fatal(err)
	}

	m2, err := NewBooleanMatrix([][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	})
	if err != nil {
		t.Fatal(err)
	}

	m3, err := NewBooleanMatrix([][]int{
		[]int{0, 1},
		[]int{0, 0},
		[]int{1, 0},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err = m1.Join(m3); err == nil {
		t.Fatal("want err, got nil")
	}

	got, err := m1.Join(m2)
	if err != nil {
		t.Fatal(err)
	}

	want, _ := NewBooleanMatrix([][]int{
		[]int{0, 1, 1},
		[]int{1, 0, 1},
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestMeet(t *testing.T) {
	m1, err := NewBooleanMatrix([][]int{
		[]int{0, 1, 1},
		[]int{1, 0, 1},
	})
	if err != nil {
		t.Fatal(err)
	}

	m2, err := NewBooleanMatrix([][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	})
	if err != nil {
		t.Fatal(err)
	}

	m3, err := NewBooleanMatrix([][]int{
		[]int{0, 1},
		[]int{0, 0},
		[]int{1, 0},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err = m1.Meet(m3); err == nil {
		t.Fatal("want err, got nil")
	}

	got, err := m1.Meet(m2)
	if err != nil {
		t.Fatal(err)
	}

	want, _ := NewBooleanMatrix([][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestBooleanPower(t *testing.T) {
	m, err := NewBooleanMatrix([][]int{
		[]int{1, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := m.Power(10)
	if err != nil {
		t.Fatal(err)
	}

	want, _ := NewBooleanMatrix([][]int{
		[]int{1, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}
