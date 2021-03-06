// Package set provides an interface and implementations for the set data structure
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
