package solver

import "testing"

func TestDay04(t *testing.T) {
	var (
		input = []string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
		}
		d = &Day04{input: input}
	)

	t.Run("first problem", func(t *testing.T) {
		var (
			expected uint64 = 18
			result          = d.First()
		)
		if result != expected {
			t.Errorf("expected %d. got: %d", expected, result)
		}
	})

	t.Run("second problem", func(t *testing.T) {
		var (
			expected uint64 = 9
			result          = d.Second()
		)
		if result != expected {
			t.Errorf("expected %d. got: %d", expected, result)
		}
	})
}
