package set_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nikovacevic/set"
)

func TestInterface(t *testing.T) {
	s := set.NewMapSet()

	if _, ok := interface{}(s).(set.Set); !ok {
		t.Errorf("MapSet should implement Set interface. It does not.")
	}
}

var addTests = []struct {
	e            interface{}
	expected     bool
	expectedSize int
}{
	{1, true, 1},
	{2, true, 2},
	{1, false, 2},
	{3, true, 3},
}

func TestAdd(t *testing.T) {
	s := set.NewMapSet()

	for _, test := range addTests {
		if actual := s.Add(test.e); actual != test.expected || s.Size() != test.expectedSize {
			t.Errorf(
				"Add(%v) should return %v with size %d; returned %v with size %d",
				test.e,
				test.expected,
				test.expectedSize,
				actual,
				s.Size(),
			)
		}
	}
}

var removeTests = []struct {
	e            interface{}
	expected     bool
	expectedSize int
}{
	{1, true, 2},
	{4, false, 2},
	{1, false, 2},
	{2, true, 1},
}

func TestRemove(t *testing.T) {
	s := set.NewMapSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	for _, test := range removeTests {
		if actual := s.Remove(test.e); actual != test.expected || s.Size() != test.expectedSize {
			t.Errorf(
				"Remove(%v) should return %v with size %d; returned %v with size %d",
				test.e,
				test.expected,
				test.expectedSize,
				actual,
				s.Size(),
			)
		}
	}
}

func TestSize(t *testing.T) {
	s := set.NewMapSet()

	if actual := s.Size(); actual != 0 {
		t.Errorf("Size() should return %v; returned %v", 0, actual)
	}

	s.Add(1)
	if actual := s.Size(); actual != 1 {
		t.Errorf("Size() should return %v; returned %v", 1, actual)
	}

	s.Add(1)
	if actual := s.Size(); actual != 1 {
		t.Errorf("Size() should return %v; returned %v", 1, actual)
	}

	s.Remove(1)
	if actual := s.Size(); actual != 0 {
		t.Errorf("Size() should return %v; returned %v", 0, actual)
	}

	s.Remove(1)
	if actual := s.Size(); actual != 0 {
		t.Errorf("Size() should return %v; returned %v", 0, actual)
	}
}

func TestClear(t *testing.T) {
	s := set.NewMapSet()

	if actual := s.Clear(); actual != false {
		t.Errorf("Clear() should return %v; returned %v", false, actual)
	}

	s.Add(1)
	if actual := s.Clear(); actual != true {
		t.Errorf("Clear() should return %v; returned %v", true, actual)
	}
}

var containsTests = []struct {
	e        interface{}
	expected bool
}{
	{1, true},
	{2, false},
	{3, false},
	{4, true},
}

func TestContains(t *testing.T) {
	s := set.NewMapSet()
	s.Add(1)
	s.Add(3)
	s.Remove(3)
	s.Add(4)
	s.Remove(4)
	s.Add(4)

	for _, test := range containsTests {
		if actual := s.Contains(test.e); actual != test.expected {
			t.Errorf("Contains(%v) should return %v; returned %v", test.e, test.expected, actual)
		}
	}
}

func TestEquals(t *testing.T) {
	s1 := set.NewMapSet()
	s2 := set.NewMapSet()

	if !s1.Equals(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.Equals(s2) should return true; returns false.")
	}
	if !s2.Equals(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.Equals(s1) should return true; returns false.")
	}

	s1.Add(1)
	s1.Add(4)
	s1.Add(2)
	s1.Add(3)
	if s1.Equals(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.Equals(s2) should return false; returns true.")
	}

	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)
	if !s1.Equals(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.Equals(s2) should return true; returns false.")
	}

	s2.Remove(3)
	if s1.Equals(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.Equals(s2) should return false; returns true.")
	}
}

func TestIsSubsetOf(t *testing.T) {
	s1 := set.NewMapSet()
	s2 := set.NewMapSet()

	if !s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return true; returns false.")
	}
	if !s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return true; returns false.")
	}

	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	if !s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return true; returns false.")
	}
	if s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return false; returns true.")
	}

	s1.Add(1)
	if !s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return true; returns false.")
	}
	if s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return false; returns true.")
	}

	s1.Add(2)
	s1.Add(3)
	if !s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return true; returns false.")
	}
	if !s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return true; returns false.")
	}

	s1.Add(4)
	s2.Add(5)
	if s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return false; returns true.")
	}
	if s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return false; returns true.")
	}
}

func TestUnion(t *testing.T) {
	s1 := set.NewMapSet()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.NewMapSet()
	s2.Add(2)
	s2.Add(4)
	s2.Add(6)

	u := s1.Union(s2)

	expected := set.NewMapSet()
	expected.Add(1)
	expected.Add(2)
	expected.Add(3)
	expected.Add(4)
	expected.Add(6)

	if !u.Equals(expected) {
		t.Errorf("Union() should create expected: %s; actual: %s", expected, u)
	}
}

func TestIntersection(t *testing.T) {
	s1 := set.NewMapSet()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.NewMapSet()
	s2.Add(2)
	s2.Add(4)
	s2.Add(6)

	u := s1.Intersection(s2)

	expected := set.NewMapSet()
	expected.Add(2)

	if !u.Equals(expected) {
		t.Errorf("Intersection() should create expected: %s; actual: %s", expected, u)
	}
}

func TestDifference(t *testing.T) {
	s1 := set.NewMapSet()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.NewMapSet()
	s2.Add(2)
	s2.Add(4)
	s2.Add(6)

	u := s1.Difference(s2)

	expected := set.NewMapSet()
	expected.Add(1)
	expected.Add(3)
	expected.Add(4)
	expected.Add(6)

	if !u.Equals(expected) {
		t.Errorf("Difference() should create expected: %s; actual: %s", expected, u)
	}
}

func TestChannel(t *testing.T) {
	s1 := set.NewMapSet()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.NewMapSet()

	for e := range s1.Channel() {
		s2.Add(e)
		// Simulate RPC
		// time.Sleep(time.Second)
	}

	if !s1.Equals(s2) {
		t.Errorf("Channel() completed; s1.Equals(s2) should return true; returns false.")
	}
}

func TestSlice(t *testing.T) {
	s1 := set.NewMapSet()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	expected := []interface{}{1, 2, 3}

	// TODO determine if Slice() will deterministically create the same
	// slice each time in the same order that elements were added.
	if actual := s1.Slice(); !reflect.DeepEqual(actual, expected) {
		t.Errorf("Slice() should return %v; returns %v.", expected, actual)
	}
}
