package set_test

import (
	"testing"

	"github.com/nikovacevic/set"
)

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
				"Add(%v) should return %v with size %d. Returned %v with size %d",
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
				"Remove(%v) should return %v with size %d. Returned %v with size %d",
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
		t.Errorf("Size() should return %v. Returned %v", 0, actual)
	}

	s.Add(1)
	if actual := s.Size(); actual != 1 {
		t.Errorf("Size() should return %v. Returned %v", 1, actual)
	}

	s.Add(1)
	if actual := s.Size(); actual != 1 {
		t.Errorf("Size() should return %v. Returned %v", 1, actual)
	}

	s.Remove(1)
	if actual := s.Size(); actual != 0 {
		t.Errorf("Size() should return %v. Returned %v", 0, actual)
	}

	s.Remove(1)
	if actual := s.Size(); actual != 0 {
		t.Errorf("Size() should return %v. Returned %v", 0, actual)
	}
}
