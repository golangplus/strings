package stringsp

import (
	"fmt"
)

// StrSet is a set of strings
type StrSet map[string]struct{}

// NewStrSet creates a string set with specified elements.
// You need to use this function only when you want to create a StrSet with
// intial elements. Different from a normal map, a nil value of StrSet is ok to add
// elemnents by calling the Add method.
func NewStrSet(els ...string) (s StrSet) {
	return s.Add(els...)
}

// Add adds elements to the set. The set can be nil
func (s *StrSet) Add(els ...string) StrSet {
	if *s == nil {
		*s = make(StrSet)
	}
	for _, el := range els {
		(*s)[el] = struct{}{}
	}

	return *s
}

// Delete removes elements from the set
func (s StrSet) Delete(els ...string) StrSet {
	if s == nil {
		return s
	}

	for _, el := range els {
		delete(s, el)
	}

	return s
}

// Contain returns true if the specified element is in the set, false otherwise
func (s StrSet) Contain(el string) bool {
	_, in := s[el]
	return in
}

// Elements returns all elements in the set as a string slice.
// NOTE the order of elements may change even if the set does not change.
func (s StrSet) Elements() (els []string) {
	els = make([]string, 0, len(s))
	for el := range s {
		els = append(els, el)
	}

	return els
}

// Equal checks whether the set has same elements with another set.
func (s StrSet) Equal(t StrSet) bool {
	if len(s) != len(t) {
		return false
	}

	for el := range s {
		if !t.Contain(el) {
			return false
		}
	}

	return true
}

// String returns the string presentation of the set
func (s StrSet) String() string {
	return fmt.Sprint(s.Elements())
}
