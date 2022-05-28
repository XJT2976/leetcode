package leetcode40

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	type Input struct {
		candidates []int
		target     int
		expect     [][]int
	}
	testInput := []Input{
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for _, input := range testInput {
		result := CombinationSum2(input.candidates, input.target)
		assert.Equal(t, input.expect, result)
	}
}
