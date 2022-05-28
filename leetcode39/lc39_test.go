package leetcode39

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	type Input struct {
		candidates []int
		target     int
		expect     [][]int
	}
	testInput := []Input{
		{
			[]int{2, 3, 5},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			[]int{2},
			1,
			[][]int{},
		},
	}

	for _, input := range testInput {
		result := CombinationSum(input.candidates, input.target)
		assert.Equal(t, input.expect, result)
	}
}
