package set_test

import (
	"fmt"
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
		t.Errorf("s1.Equals(s2) should return true. Returns false.")
	}
	if !s2.Equals(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.Equals(s1) should return true. Returns false.")
	}

	s1.Add(1)
	s1.Add(4)
	s1.Add(2)
	s1.Add(3)
	if s1.Equals(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.Equals(s2) should return false. Returns true.")
	}

	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)
	if !s1.Equals(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.Equals(s2) should return true. Returns false.")
	}

	s2.Remove(3)
	if s1.Equals(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.Equals(s2) should return false. Returns true.")
	}
}

func TestIsSubsetOf(t *testing.T) {
	s1 := set.NewMapSet()
	s2 := set.NewMapSet()

	if !s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return true. Returns false.")
	}
	if !s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return true. Returns false.")
	}

	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	if !s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return true. Returns false.")
	}
	if s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return false. Returns true.")
	}

	s1.Add(1)
	if !s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return true. Returns false.")
	}
	if s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return false. Returns true.")
	}

	s1.Add(2)
	s1.Add(3)
	if !s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return true. Returns false.")
	}
	if !s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return true. Returns falseg.")
	}

	s1.Add(4)
	s2.Add(5)
	if s1.IsSubsetOf(s2) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s1.IsSubsetOf(s2) should return fa;se. Returns true.")
	}
	if s2.IsSubsetOf(s1) {
		fmt.Printf("s1: %s\ns2: %s\n", s1, s2)
		t.Errorf("s2.IsSubsetOf(s1) should return false. Returns true.")
	}
}
