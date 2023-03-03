package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToSet(t *testing.T) {
	t.Parallel()

	var (
		input    = []int{1, 1, 2, 3}
		expected = map[int]bool{
			1: true,
			2: true,
			3: true,
		}
	)

	require.Equal(t, expected, ToSet(input))
}

func TestFilter(t *testing.T) {
	t.Parallel()

	var (
		input    = []int{1, 1, 2, 3}
		filterFn = func(i int) bool { return i == 1 }
		expected = []int{1, 1}
	)

	require.Equal(t, expected, Filter(input, filterFn))
}

func TestDifference(t *testing.T) {
	t.Parallel()

	var (
		a        = []int{1, 2, 3}
		b        = []int{1, 2, 3, 4}
		expected = []int{4}
	)

	require.Equal(t, expected, Difference(a, b))
}

func TestIntersection(t *testing.T) {
	t.Parallel()

	var (
		a        = []int{1, 2, 3}
		b        = []int{1, 2, 3, 4}
		expected = []int{1, 2, 3}
	)

	require.Equal(t, expected, Intersection(a, b))
}

func TestEach(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3}

	var count int
	eachFn := func(i int) { count++ }

	Each(input, eachFn)

	require.Equal(t, len(input), count)
}

func TestIncludes(t *testing.T) {
	t.Parallel()

	t.Run("slice does not include target", func(t *testing.T) {
		t.Parallel()

		input := []int{1, 2, 3}

		require.False(t, Includes(input, 4))
	})

	t.Run("slice includes target", func(t *testing.T) {
		t.Parallel()

		input := []int{1, 2, 3}

		require.True(t, Includes(input, 3))
	})
}
