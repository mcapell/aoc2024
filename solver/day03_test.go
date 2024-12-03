package solver

import "testing"

func TestDay03(t *testing.T) {
	t.Run("first problem", func(t *testing.T) {
		var (
			input = []string{
				"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			}
			expected = 161
		)

		d := &Day03{input: input}
		result := d.First()
		if result != expected {
			t.Errorf("expected %d, got: %d", expected, result)
		}
	})

	t.Run("second problem", func(t *testing.T) {
		var (
			input = []string{
				"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			}
			expected = 48
		)

		d := &Day03{input: input}
		result := d.Second()
		if result != expected {
			t.Errorf("expected %d, got: %d", expected, result)
		}
	})
}
