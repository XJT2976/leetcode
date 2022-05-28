package leetcode216

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	type Input struct {
		n      int
		k      int
		expect [][]int
	}
	testInput := []Input{
		{
			7,
			3,
			[][]int{
				{1, 2, 4},
			},
		},
		{
			9,
			3,
			[][]int{
				{1, 2, 6},
				{1, 3, 5},
				{2, 3, 4},
			},
		},
	}

	for _, input := range testInput {
		result := CombinationSum3(input.k, input.n)
		assert.Equal(t, input.expect, result)
	}
}
