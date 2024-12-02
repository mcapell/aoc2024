package solver

import (
	"testing"
)

func TestDay01(t *testing.T) {
	d := Day01{input: [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}}}

	result := d.First()
	if 11 != result {
		t.Errorf("invalid first result. Expected 11, got: %d", result)
	}

	result = d.Second()
	if 31 != result {
		t.Errorf("invalid second result. Expected 31, got: %d", result)
	}
}
