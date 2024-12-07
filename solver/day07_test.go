package solver

import "testing"

func TestDay07(t *testing.T) {
	var (
		input = [][]int{
			{190, 10, 19},
			{3267, 81, 40, 27},
			{83, 17, 5},
			{156, 15, 6},
			{7290, 6, 8, 6, 15},
			{161011, 16, 10, 13},
			{192, 17, 8, 14},
			{21037, 9, 7, 18, 13},
			{292, 11, 6, 16, 20},
		}

		d = &Day07{
			input: input,
		}
	)
	t.Run("first problem", func(t *testing.T) {
		var (
			expected uint64 = 3749
			result          = d.First()
		)
		if expected != result {
			t.Errorf("expected %d. got=%d", expected, result)
		}
	})

	t.Run("second problem", func(t *testing.T) {
		var (
			expected uint64 = 11387
			result          = d.Second()
		)
		if expected != result {
			t.Errorf("expected %d. got=%d", expected, result)
		}
	})
}
