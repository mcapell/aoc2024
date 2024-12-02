package solver

import (
	"testing"
)

func TestDay02(t *testing.T) {
	d := Day02{[][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}}

	result := d.First()
	if 2 != result {
		t.Errorf("invalid first result. Expected 2, got: %d", result)
	}

	result = d.Second()
	if 4 != result {
		t.Errorf("invalid second result. Expected 4, got: %d", result)
	}
}

func TestDay02isSafe(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{62, 65, 67, 70, 73, 76, 75}, false},
		{[]int{68, 71, 73, 76, 78, 78}, false},
		{[]int{77, 80, 81, 82, 86}, false},
		{[]int{51, 53, 56, 57, 59, 56, 57, 60}, false},
		{[]int{16, 13, 14, 11, 16}, false},
		{[]int{16, 18, 19, 21, 22}, true},
		{[]int{16, 14, 13, 11, 8}, true},
	}

	for _, tc := range tests {
		d := &Day02{input: [][]int{tc.input}}

		result := d.First() == 1
		if result != tc.expected {
			t.Errorf("expected safeness of %v to be %v, got: %v", tc.input, tc.expected, result)
		}
	}
}
