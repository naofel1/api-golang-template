package pagination

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	t.Parallel()

	p := Front{}

	p.CurrentPage = 2
	p.ItemsPerPage = 10

	NumItems := 42

	p.Calculate(NumItems)

	// Output:
	// NumPages: 5
	require.Equal(t, 5, p.NumPages)
	// HasPrev: true
	require.Equal(t, true, p.HasPrev)
	// PrevPage: 1
	require.Equal(t, 1, p.PrevPage)
	// HasNext: true
	require.Equal(t, true, p.HasNext)
	// NextPage: 3
	require.Equal(t, 3, p.NextPage)
	// Offset: 10
	require.Equal(t, 10, p.Offset)
}
