package leetcode51

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	type Input struct {
		n      int
		expect [][]string
	}
	testInput := []Input{
		{
			1,
			[][]string{
				{"Q"},
			},
		},
		{
			4,
			[][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for _, input := range testInput {
		result := SolveNQueens(input.n)
		assert.Equal(t, input.expect, result)
	}
}
