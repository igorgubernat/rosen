package ex2

import (
	"reflect"
	"sort"
	"testing"
)

func TestElements(t *testing.T) {
	m := NewMultiset("one", "two", "two", "three", "three", "three")

	got := m.Elements()
	sort.Strings(got)
	want := []string{"one", "two", "two", "three", "three", "three"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestUnion(t *testing.T) {
	m1 := NewMultiset("one", "two", "two", "three", "three", "three")
	m2 := NewMultiset("one", "one", "one", "three", "three")

	got := m1.Union(m2).Elements()
	sort.Strings(got)
	want := []string{"one", "one", "one", "two", "two", "three", "three", "three"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestIntersect(t *testing.T) {
	m1 := NewMultiset("one", "two", "two", "three", "three", "three")
	m2 := NewMultiset("two", "three", "three", "three", "three", "three")

	got := m1.Intersect(m2).Elements()
	sort.Strings(got)
	want := []string{"two", "three", "three", "three"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSubtract(t *testing.T) {
	m1 := NewMultiset("one", "two", "two", "three", "three", "three")
	m2 := NewMultiset("two", "three", "three", "three", "three", "three")

	got := m1.Subtract(m2).Elements()
	sort.Strings(got)
	want := []string{"one", "two"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestAdd(t *testing.T) {
	m1 := NewMultiset("one", "two", "two", "three", "three", "three")
	m2 := NewMultiset("two", "three", "three", "three", "three", "three", "four")

	got := m1.Add(m2).Elements()
	sort.Strings(got)
	want := []string{"one", "two", "two", "two", "three", "three", "three", "three", "three", "three", "three", "three", "four"}
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}
