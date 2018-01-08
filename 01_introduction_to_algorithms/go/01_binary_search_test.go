package binarySearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 3, 5, 7, 9}
	if pos := binarySearch(list, 3); pos != 1 {
		t.Errorf("Expected %d, but got %d", 1, pos)
	}

	if pos := binarySearch(list, 23); pos != -1 {
		t.Errorf("Expected %d, but got %d", -1, pos)
	}
}
