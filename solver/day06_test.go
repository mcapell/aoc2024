package solver

import "testing"

func TestDay06(t *testing.T) {
	var (
		input = [][]byte{
			{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
			{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		}
		d = &Day06{input: input}
	)

	t.Run("first problem", func(t *testing.T) {
		var (
			expected uint64 = 41
			result          = d.First()
		)
		if expected != result {
			t.Errorf("expected %d. got=%d", expected, result)
		}
	})

	t.Run("second problem", func(t *testing.T) {
		var (
			expected uint64 = 6
			result          = d.Second()
		)
		if expected != result {
			t.Errorf("expected %d. got=%d", expected, result)
		}
	})
}
