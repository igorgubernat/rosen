package ex1

import (
	"reflect"
	"sort"
	"testing"
)

func TestElements(t *testing.T) {
	u := NewUniversalSet("one", "two", "three")

	oneTwo, err := u.Set("one", "two")
	if err != nil {
		t.Fatal(err)
	}

	got := oneTwo.Elements()
	want := []string{"one", "two"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got elements %v, but want %v", got, want)
	}
}

func TestComplement(t *testing.T) {
	u := NewUniversalSet("one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten")

	odd, err := u.Set("one", "three", "five", "seven", "nine")
	if err != nil {
		t.Fatal(err)
	}

	got := odd.Complement().Elements()
	want := []string{"two", "four", "six", "eight", "ten"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("complement of %v: got %v, but want %v", odd.Elements(), got, want)
	}
}

func TestUnion(t *testing.T) {
	u := NewUniversalSet("one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten")

	s12, err := u.Set("one", "two")
	if err != nil {
		t.Fatal(err)
	}

	s156, err := u.Set("one", "five", "six")
	if err != nil {
		t.Fatal(err)
	}

	got := s12.Union(s156).Elements()
	want := []string{"one", "two", "five", "six"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("join of %v and %v: got %v, but want %v", s12.Elements(), s156.Elements(), got, want)
	}
}

func TestIntersect(t *testing.T) {
	u := NewUniversalSet("one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten")

	s12345, err := u.Set("one", "two", "three", "four", "five")
	if err != nil {
		t.Fatal(err)
	}

	s156, err := u.Set("one", "five", "six")
	if err != nil {
		t.Fatal(err)
	}

	got := s12345.Intersect(s156).Elements()
	want := []string{"one", "five"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("intersect of %v and %v: got %v, but want %v", s12345.Elements(), s156.Elements(), got, want)
	}
}

func TestSubtract(t *testing.T) {
	u := NewUniversalSet("one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten")

	s12345, err := u.Set("one", "two", "three", "four", "five")
	if err != nil {
		t.Fatal(err)
	}

	s156, err := u.Set("one", "five", "six")
	if err != nil {
		t.Fatal(err)
	}

	got := s12345.Subtract(s156).Elements()
	want := []string{"two", "three", "four"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("subtract of %v and %v: got %v, but want %v", s12345.Elements(), s156.Elements(), got, want)
	}
}

func TestDifference(t *testing.T) {
	u := NewUniversalSet("one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten")

	s12345, err := u.Set("one", "two", "three", "four", "five")
	if err != nil {
		t.Fatal(err)
	}

	s156, err := u.Set("one", "five", "six")
	if err != nil {
		t.Fatal(err)
	}

	got := s12345.Difference(s156).Elements()
	want := []string{"two", "three", "four", "six"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("difference of %v and %v: got %v, but want %v", s12345.Elements(), s156.Elements(), got, want)
	}
}
