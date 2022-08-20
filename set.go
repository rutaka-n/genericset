// genericset provides a simple map-based implementation of generic set.
package genericset

import (
	"fmt"
	"sync"
)

type Set[K comparable] struct {
	m    sync.RWMutex
	data map[K]struct{}
}

// New returns new set of type K.
func New[K comparable]() Set[K] {
	return Set[K]{
		m:    sync.RWMutex{},
		data: make(map[K]struct{}),
	}
}

// Add — add elements to a set.
func (s *Set[T]) Add(elements ...T) {
	s.m.Lock()
	defer s.m.Unlock()
	for _, elem := range elements {
		s.data[elem] = struct{}{}
	}
}

// Del – remove element from set.
func (s *Set[T]) Del(elem T) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.data, elem)
}

// Size returns size of the set.
func (s *Set[T]) Size() int {
	s.m.RLock()
	defer s.m.RUnlock()
	return len(s.data)
}

// String returns string representation of the set.
func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}

// ToSlice returns slice with elements of the set.
func (s *Set[T]) ToSlice() []T {
	s.m.RLock()
	defer s.m.RUnlock()
	result := make([]T, 0, len(s.data))
	for elem, _ := range s.data {
		result = append(result, elem)
	}
	return result
}

func (s *Set[T]) IsElement(elem T) bool {
	s.m.RLock()
	defer s.m.RUnlock()
	_, exists := s.data[elem]
	return exists
}

// IsEmpty returns true if there are no elements in set.
func (s *Set[T]) IsEmpty() bool {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.Size() == 0
}

// IsSubset returns true when every element of set `s` is also a member of set `another`, otherwise false.
func (s *Set[T]) IsSubset(another *Set[T]) bool {
	s.m.RLock()
	defer s.m.RUnlock()
	another.m.RLock()
	defer another.m.RUnlock()
	for elem, _ := range s.data {
		if _, exists := another.data[elem]; !exists {
			return false
		}
	}
	return true
}

// Intersection returns the intersection of set `s` and set `another`.
func (s *Set[T]) Intersection(another *Set[T]) *Set[T] {
	result := New[T]()
	s.m.RLock()
	defer s.m.RUnlock()
	another.m.RLock()
	defer another.m.RUnlock()
	for elem, _ := range s.data {
		if _, exists := another.data[elem]; exists {
			result.Add(elem)
		}
	}
	return &result
}

// IsDisjoint returns true if set `s` and set `another` are disjoint (have no elements in common), otherwise false.
func (s *Set[T]) IsDisjoint(another *Set[T]) bool {
	s.m.RLock()
	defer s.m.RUnlock()
	another.m.RLock()
	defer another.m.RUnlock()
	for elem, _ := range s.data {
		if _, exists := another.data[elem]; exists {
			return false
		}
	}
	return true
}

// Union returns the merged (union) set of `s` and `another`.
func (s *Set[T]) Union(another *Set[T]) *Set[T] {
	res := New[T]()
	res.Add(s.ToSlice()...)
	res.Add(another.ToSlice()...)
	return &res
}

// Difference returns the difference between two sets. The new set consists of all elements that are in `s` but not in `another`.
func (s *Set[T]) Difference(another *Set[T]) *Set[T] {
	result := New[T]()
	s.m.RLock()
	defer s.m.RUnlock()
	another.m.RLock()
	defer another.m.RUnlock()
	for elem, _ := range s.data {
		if _, exists := another.data[elem]; !exists {
			result.Add(elem)
		}
	}
	return &result
}
