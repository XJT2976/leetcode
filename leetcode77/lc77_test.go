package leetcode77

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombine(t *testing.T) {
	type Input struct {
		n      int
		k      int
		expect [][]int
	}
	testInput := []Input{
		{
			4,
			2,
			[][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
		{
			1,
			1,
			[][]int{
				{1}},
		},
		{
			6,
			3,
			[][]int{
				{1, 2, 3},
				{1, 2, 4},
				{1, 2, 5},
				{1, 2, 6},
				{1, 3, 4},
				{1, 3, 5},
				{1, 3, 6},
				{1, 4, 5},
				{1, 4, 6},
				{1, 5, 6},
				{2, 3, 4},
				{2, 3, 5},
				{2, 3, 6},
				{2, 4, 5},
				{2, 4, 6},
				{2, 5, 6},
				{3, 4, 5},
				{3, 4, 6},
				{3, 5, 6},
				{4, 5, 6},
			},
		},
	}

	for _, input := range testInput {
		result := Combine(input.n, input.k)
		assert.Equal(t, input.expect, result)
	}
}
