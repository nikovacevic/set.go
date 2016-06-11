package set

// Set describes a collection of distinct elements
type Set interface {
	// Adds element e to the Set
	// returns true if the element was added to the set
	// returns false if the set already contained the element e
	Add(e interface{}) bool
	// Removes element e from the Set
	// returns true if the element was removed from the set
	// returns false if the set did not contain the element e
	Remove(e interface{}) bool
	// Returns the number of elements in the set
	Size() int
	// Removes all elements from the set
	// returns true if all elements were removed
	// returns false if the set was empty
	Clear() bool
	// Returns true if the set contains the element e
	Contains(e interface{}) bool
	// Returns true if the set contains all elements in set s and the two
	// sets are of equal size
	Equals(s Set) bool
	// Returns true if the set s contains all elements in the set
	IsSubsetOf(s Set) bool
	// Returns the set of elements in either or both of the set and set s
	Union(s Set) Set
	// Returns the set of elements in both the set and the set s
	Intersection(s Set) Set
	// Returns the set of elements in either the set or the set s, but not in both
	Difference(s Set) Set
}

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
