package genericset

import (
	"testing"
)

func TestIsElement(t *testing.T) {
	cases := []struct {
		desc     string
		items    []int
		elem     int
		expected bool
	}{
		{"element should be in the set", []int{1, 2, 3}, 1, true},
		{"element should be in the set", []int{1, 2, 3, 4, 5}, 5, true},
		{"element in not in the set", []int{1, 2, 3, 4, 5}, 10, false},
		{"element in not in the set", []int{}, 10, false},
	}
	for _, tc := range cases {
		s := New[int]()
		s.Add(tc.items...)
		actual := s.IsElement(tc.elem)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %v got: %v for set: %v and element %d",
				tc.desc, tc.expected, actual, s, tc.elem)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	s := New[string]()
	if !s.IsEmpty() {
		t.Error("set should be empty")
	}
	s.Add("elem")
	if s.IsEmpty() {
		t.Error("set is not empty")
	}
}

func TestIsSubset(t *testing.T) {
	cases := []struct {
		desc     string
		items1   []int
		items2   []int
		expected bool
	}{
		{"equal sets", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"s1 is a subset of s2", []int{1, 2, 3}, []int{1, 2, 3, 4, 5}, true},
		{"sets has no common elements", []int{1, 2}, []int{3, 4}, false},
		{"s1 cannot be a subset of an empty set", []int{1, 2, 3}, []int{}, false},
		{"both sets are empty", []int{}, []int{}, true},
		{"empty set are subset of any set", []int{}, []int{1}, true},
	}
	for _, tc := range cases {
		s1 := New[int]()
		s1.Add(tc.items1...)
		s2 := New[int]()
		s2.Add(tc.items2...)
		actual := s1.IsSubset(&s2)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %v got: %v for set: %v and set %v",
				tc.desc, tc.expected, actual, s1, s2)
		}
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		desc     string
		items1   []int
		items2   []int
		expected []int
	}{
		{"equal sets", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"sets has no common elements", []int{1, 2}, []int{3, 4}, []int{}},
		{"one set is empty", []int{1, 2, 3}, []int{}, []int{}},
		{"both sets are empty", []int{}, []int{}, []int{}},
	}
	for _, tc := range cases {
		s1 := New[int]()
		s1.Add(tc.items1...)
		s2 := New[int]()
		s2.Add(tc.items2...)
		expected := New[int]()
		expected.Add(tc.expected...)

		actual := s1.Intersection(&s2)
		if actual.Size() != expected.Size() || !expected.IsSubset(actual) {
			t.Fatalf("%s: expected: %v got: %v for set: %v and set %v",
				tc.desc, tc.expected, actual, s1, s2)
		}
	}
}

func TestIsDisjoint(t *testing.T) {
	cases := []struct {
		desc     string
		items1   []int
		items2   []int
		expected bool
	}{
		{"equal sets", []int{1, 2, 3}, []int{1, 2, 3}, false},
		{"s1 is a subset of s2", []int{1, 2, 3}, []int{1, 2, 3, 4, 5}, false},
		{"sets has no common elements", []int{1, 2}, []int{3, 4}, true},
		{"both sets are empty", []int{}, []int{}, true},
		{"one set is empty", []int{}, []int{1}, true},
	}
	for _, tc := range cases {
		s1 := New[int]()
		s1.Add(tc.items1...)
		s2 := New[int]()
		s2.Add(tc.items2...)
		actual := s1.IsDisjoint(&s2)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %v got: %v for set: %v and set %v",
				tc.desc, tc.expected, actual, s1, s2)
		}
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		desc     string
		items1   []int
		items2   []int
		expected []int
	}{
		{"equal sets", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"sets has no common elements", []int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{"one set is empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"both sets are empty", []int{}, []int{}, []int{}},
	}
	for _, tc := range cases {
		s1 := New[int]()
		s1.Add(tc.items1...)
		s2 := New[int]()
		s2.Add(tc.items2...)
		expected := New[int]()
		expected.Add(tc.expected...)

		actual := s1.Union(&s2)
		if actual.Size() != expected.Size() || !expected.IsSubset(actual) {
			t.Fatalf("%s: expected: %v got: %v for set: %v and set %v",
				tc.desc, tc.expected, actual, s1, s2)
		}
	}
}
