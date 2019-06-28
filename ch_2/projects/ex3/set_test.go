package ex3

import (
	"reflect"
	"testing"
)

func TestComplement(t *testing.T) {
	u := NewUniverse("one", "two", "three", "four", "five")

	if _, err := u.Set(map[string]float64{
		"blah": 0.6,
	}); err == nil {
		t.Fatal("want error, got nil")
	}

	s, err := u.Set(map[string]float64{
		"two":  0.3,
		"four": 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	got := s.Complement().m
	want := map[string]float64{
		"one":   1,
		"two":   0.7,
		"three": 1,
		"five":  1,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestUnion(t *testing.T) {
	u := NewUniverse("one", "two", "three", "four", "five")

	s1, err := u.Set(map[string]float64{
		"two":  0.3,
		"four": 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	s2, err := u.Set(map[string]float64{
		"one":   0.8,
		"two":   0.7,
		"three": 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	got := s1.Union(s2).m
	want := map[string]float64{
		"one":   0.8,
		"two":   0.7,
		"three": 1,
		"four":  1,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestIntersection(t *testing.T) {
	u := NewUniverse("one", "two", "three", "four", "five")

	s1, err := u.Set(map[string]float64{
		"two":  0.3,
		"four": 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	s2, err := u.Set(map[string]float64{
		"one":   0.8,
		"two":   0.7,
		"three": 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	got := s1.Intersection(s2).m
	want := map[string]float64{
		"two": 0.3,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}
