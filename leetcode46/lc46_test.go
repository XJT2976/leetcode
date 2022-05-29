package leetcode46

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermute(t *testing.T) {
	type Input struct {
		nums   []int
		expect [][]int
	}
	testInput := []Input{
		{
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			[]int{0, 1},
			[][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			[]int{1},
			[][]int{
				{1},
			},
		},
	}

	for _, input := range testInput {
		result := Permute(input.nums)
		assert.Equal(t, input.expect, result)
	}
}
