package set

import (
	"fmt"
	"strings"
)

// This implementation leverages the empty interface and empty struct. If you
// are curious about why, please reference Dave Cheney's excellent blog post
// (http://dave.cheney.net/2014/03/25/the-empty-struct) on the matter.
// Furthermore, struct{}{} creates a struct literal of type struct{}

// MapSet implements Set. Elements of the set are represented as keys of the
// map. Hence, the key is of type interface{} to allow any type to be added to
// the set. The value of the map is type struct{} because it has zero width.
type MapSet map[interface{}]struct{}

// NewMapSet creates and returns a pointer to a new MapSet instance
func NewMapSet() *MapSet {
	return &MapSet{}
}

// Add adds e elements to the Set
// returns true if the element was added to the set
// returns false if the set already contained the element e
func (ms *MapSet) Add(e interface{}) bool {
	if _, exists := (*ms)[e]; exists {
		return false
	}
	(*ms)[e] = struct{}{}
	return true
}

// Remove removes element e from the Set
// returns true if the element was removed from the set
// returns false if the set did not contain the element e
func (ms *MapSet) Remove(e interface{}) bool {
	if _, exists := (*ms)[e]; !exists {
		return false
	}
	delete((*ms), e)
	return true
}

// Size returns the number of elements in the set
func (ms *MapSet) Size() int {
	return len(*ms)
}

// Clear removes all elements from the set
// returns true if all elements were removed
// returns false if the set was empty
func (ms *MapSet) Clear() bool {
	if ms.Size() == 0 {
		return false
	}
	ms = NewMapSet()
	return true
}

// Contains returns true if the set contains the element e
func (ms *MapSet) Contains(e interface{}) bool {
	if _, exists := (*ms)[e]; exists {
		return true
	}
	return false
}

// Returns true if the set contains all elements in set s and the two
// sets are of equal size
func (ms *MapSet) Equals(s Set) bool {
	if s.Size() != ms.Size() {
		return false
	}
	for e, _ := range *ms {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// Returns true if the set s contains all elements in the set
func (ms *MapSet) IsSubsetOf(s Set) bool {
	return true
}

// Returns the set of elements in either or both of the set and set s
func (ms *MapSet) Union(s Set) Set {
	return ms
}

// Returns the set of elements in both the set and the set s
func (ms *MapSet) Intersection(s Set) Set {
	return ms
}

// Returns the set of elements in either the set or the set s, but not in both
func (ms *MapSet) Difference(s Set) Set {
	return ms
}

func (ms *MapSet) String() string {
	str := "{ "
	elements := []string{}
	for e, _ := range *ms {
		elements = append(elements, fmt.Sprintf("%v", e))
	}
	str += strings.Join(elements, ", ")
	str += " }"
	return str
}
