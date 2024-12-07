package solver

import "testing"

func TestDay05(t *testing.T) {
	var (
		rules = map[int][]int{
			47: {53, 13, 61, 29},
			97: {13, 61, 47, 29, 53, 75},
			75: {29, 53, 47, 61, 13},
			61: {13, 53, 29},
			29: {13},
			53: {29, 13},
		}

		pages = [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}

		d = Day05{
			rules: rules,
			pages: pages,
		}
	)

	t.Run("first problem", func(t *testing.T) {
		var (
			expected uint64 = 143
			result          = d.First()
		)
		if expected != result {
			t.Errorf("expected %d. got=%d", expected, result)
		}
	})

	t.Run("second problem", func(t *testing.T) {
		var (
			expected uint64 = 123
			result          = d.Second()
		)
		if expected != result {
			t.Errorf("expected %d. got=%d", expected, result)
		}
	})
}
